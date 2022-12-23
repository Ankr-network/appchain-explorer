package staking

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/Ankr-network/ankr-protocol/shared"
	common2 "github.com/Ankr-network/ankr-protocol/shared/common"
	"github.com/Ankr-network/ankr-protocol/shared/database"
	"github.com/Ankr-network/ankr-protocol/shared/staking/abigen"
	"github.com/Ankr-network/ankr-protocol/shared/types"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	types2 "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/pkg/errors"
	"github.com/samber/lo"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"math/big"
	"net/http"
	"strings"
	"time"
)

type Service struct {
	state *database.StateDb
	// start
	config *Config
	eth    *ethclient.Client
	// abi
	staking     *abigen.Staking
	chainConfig *abigen.ChainConfig
	// state
	cachedChainConfig *types.ChainConfig
	cachedChains      []*types.Chain
}

func NewService(state *database.StateDb) *Service {
	return &Service{state: state}
}

func (s *Service) Start(cp shared.IConfigProvider) (err error) {
	config := &Config{}
	if err := cp.Parse(config); err != nil {
		return err
	}
	s.config = config
	s.eth, err = ethclient.Dial(config.Eth1Url)
	if err != nil {
		return errors.Wrapf(err, "failed to connect to Web3 node (%s)", config.Eth1Url)
	}
	s.staking, _ = abigen.NewStaking(config.StakingContract, s.eth)
	s.chainConfig, _ = abigen.NewChainConfig(config.ChainConfigContract, s.eth)
	newChainConfig, err := s.refreshChainConfig()
	if err != nil {
		return err
	}
	s.cachedChainConfig = newChainConfig
	go func() {
		if err := s.backgroundWorker(); err != nil {
			log.Fatalf("failed to start background worker: %+v", err)
		}
	}()
	return nil
}

func (s *Service) backgroundWorker() (err error) {
	// do recovery
	hasMore := true
	for hasMore {
		hasMore, err = s.processEventLogs(context.TODO())
		if err != nil {
			return err
		}
	}
	// run worker
	refreshChainConfig := time.Tick(10 * time.Minute)
	updateChainsTick := time.Tick(15 * time.Minute)
	processEventLogsTick := time.Tick(10 * time.Second)
	for {
		var err error
		select {
		case <-refreshChainConfig:
			var newChainConfig *types.ChainConfig
			newChainConfig, err = s.refreshChainConfig()
			if err != nil {
				return err
			}
			s.cachedChainConfig = newChainConfig
		case <-updateChainsTick:
			s.cachedChains, err = s.GetChains(context.TODO())
			if err != nil {
				return err
			}
		case <-processEventLogsTick:
			_, err = s.processEventLogs(context.TODO())
			if err != nil {
				return err
			}
		}
	}
}

func (s *Service) refreshChainConfig() (chainConfig *types.ChainConfig, err error) {
	chainConfig = &types.ChainConfig{
		AverageBlockTime: s.config.BlockTime,
	}
	opts := &bind.CallOpts{}
	chainConfig.ActiveValidatorsLength, err = s.chainConfig.GetActiveValidatorsLength(opts)
	if err != nil {
		return nil, err
	}
	chainConfig.EpochBlockInterval, err = s.chainConfig.GetEpochBlockInterval(opts)
	if err != nil {
		return nil, err
	}
	chainConfig.MisdemeanorThreshold, err = s.chainConfig.GetMisdemeanorThreshold(opts)
	if err != nil {
		return nil, err
	}
	chainConfig.FelonyThreshold, err = s.chainConfig.GetFelonyThreshold(opts)
	if err != nil {
		return nil, err
	}
	chainConfig.ValidatorJailEpochLength, err = s.chainConfig.GetValidatorJailEpochLength(opts)
	if err != nil {
		return nil, err
	}
	chainConfig.UndelegatePeriod, err = s.chainConfig.GetUndelegatePeriod(opts)
	if err != nil {
		return nil, err
	}
	var bigValue *big.Int
	bigValue, err = s.chainConfig.GetMinValidatorStakeAmount(opts)
	if err != nil {
		return nil, err
	}
	chainConfig.MinValidatorStakeAmount = common2.ToEther(bigValue)
	bigValue, err = s.chainConfig.GetMinStakingAmount(opts)
	if err != nil {
		return nil, err
	}
	chainConfig.MinStakingAmount = common2.ToEther(bigValue)
	return
}

func (s *Service) processEventLogs(ctx context.Context) (hasMore bool, err error) {
	lastAffectedBlock := s.state.GetLastAffectedBlock(ctx)
	const confirmationBlocks = 10
	latestKnownBlock, err := s.eth.BlockNumber(context.TODO())
	if err != nil {
		return false, err
	}
	const processBatch = 100_000
	fromBlock := lastAffectedBlock
	toBlock := fromBlock + processBatch
	if toBlock > latestKnownBlock-confirmationBlocks {
		toBlock = latestKnownBlock - confirmationBlocks
	} else {
		hasMore = true
	}
	logs, err := s.eth.FilterLogs(context.TODO(), ethereum.FilterQuery{
		FromBlock: big.NewInt(int64(fromBlock)),
		ToBlock:   big.NewInt(int64(toBlock)),
		Addresses: []common.Address{
			s.config.StakingContract,
		},
	})
	pip := s.state.NewPipeline()
	defer func() {
		if err == nil {
			err = pip.Commit(ctx, toBlock)
		} else {
			pip.Cancel()
		}
	}()
	for _, l := range logs {
		if err = s.processEventLog(ctx, pip, l); err != nil {
			return false, err
		}
	}
	return
}

func (s *Service) processEventLog(ctx context.Context, pip *database.Pipeline, l types2.Log) error {
	if event, _ := s.staking.ParseValidatorAdded(l); event != nil {
		val, _, err := pip.AddValidator(ctx, event)
		if err != nil {
			return err
		}
		log.WithFields(logValidatorFields(val)).Infof("validator added")
		return nil
	}
	if event, _ := s.staking.ParseValidatorModified(l); event != nil {
		val, _, err := pip.ModifyValidator(ctx, event)
		if err != nil {
			return err
		}
		log.WithFields(logValidatorFields(val)).Infof("validator modified")
		return nil
	}
	if event, _ := s.staking.ParseValidatorRemoved(l); event != nil {
		val, _, err := pip.RemoveValidator(ctx, event)
		if err != nil {
			return err
		}
		log.WithFields(logValidatorFields(val)).Infof("validator removed")
		return nil
	}
	if event, _ := s.staking.ParseValidatorSlashed(l); event != nil {
		val, err := pip.SlashValidator(ctx, event)
		if err != nil {
			return err
		}
		log.WithFields(logSlashingFields(val)).Infof("validator slashed")
		return nil
	}
	if event, _ := s.staking.ParseValidatorJailed(l); event != nil {
		val, _, err := pip.JailValidator(ctx, event)
		if err != nil {
			return err
		}
		log.WithFields(logValidatorFields(val)).Infof("validator slashed")
		return nil
	}
	if event, _ := s.staking.ParseValidatorReleased(l); event != nil {
		val, _, err := pip.ReleaseValidator(ctx, event)
		if err != nil {
			return err
		}
		log.WithFields(logValidatorFields(val)).Infof("validator released")
		return nil
	}
	if event, _ := s.staking.ParseValidatorDeposited(l); event != nil {
		val, err := pip.DepositValidator(ctx, event)
		if err != nil {
			return err
		}
		log.WithFields(logDepositFields(val)).Infof("validator deposited")
		return nil
	}
	if event, _ := s.staking.ParseDelegated(l); event != nil {
		ev, err := pip.AddDelegated(ctx, event)
		if err != nil {
			return err
		}
		log.WithFields(logValidatorEventFields(ev)).Infof("validator delegated")
		return nil
	}
	if event, _ := s.staking.ParseUndelegated(l); event != nil {
		ev, err := pip.AddUndelegated(ctx, event)
		if err != nil {
			return err
		}
		log.WithFields(logValidatorEventFields(ev)).Infof("validator undelegated")
		return nil
	}
	if event, _ := s.staking.ParseClaimed(l); event != nil {
		ev, err := pip.AddClaimed(ctx, event)
		if err != nil {
			return err
		}
		log.WithFields(logValidatorEventFields(ev)).Infof("validator claimed")
		return nil
	}
	if event, _ := s.staking.ParseRedelegated(l); event != nil {
		ev, err := pip.AddRedelegated(ctx, event)
		if err != nil {
			return err
		}
		log.WithFields(logValidatorEventFields(ev)).Infof("validator redelegated")
		return nil
	}
	log.WithField("topic", l.Topics[0]).Warnf("not supported event type")
	return nil
}

type chainListResponse struct {
	Name           string   `json:"name"`
	Chain          string   `json:"chain"`
	Rpc            []string `json:"rpc"`
	Faucets        []string `json:"faucets"`
	NativeCurrency struct {
		Name     string `json:"name"`
		Symbol   string `json:"symbol"`
		Decimals uint32 `json:"decimals"`
	}
	InfoUrl   string `json:"infoUrl"`
	ChainId   uint64 `json:"chainId"`
	Slip44    uint64 `json:"slip44"`
	Explorers []struct {
		Name     string `json:"name"`
		Url      string `json:"url"`
		Standard string `json:"standard"`
	}
}

func (s *Service) getChains(ctx context.Context) (result []*types.Chain, err error) {
	req, err := http.NewRequestWithContext(ctx, "GET", s.config.ChainListUrl, nil)
	if err != nil {
		return nil, err
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	requestBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var response []chainListResponse
	err = json.Unmarshal(requestBody, &response)
	if err != nil {
		return nil, err
	}
	chains := lo.Map(response, func(t chainListResponse, i int) *types.Chain {
		return &types.Chain{
			ChainId:    t.ChainId,
			ProjectUrl: t.InfoUrl,
			Name:       t.Name,
			ShortName:  t.Chain,
			Rpc:        t.Rpc,
			Faucet:     t.Faucets,
			NativeCurrency: &types.ChainNativeCurrency{
				Name:     t.NativeCurrency.Name,
				Symbol:   t.NativeCurrency.Symbol,
				Decimals: t.NativeCurrency.Decimals,
			},
			Explorer: lo.Map(t.Explorers, func(t struct {
				Name     string `json:"name"`
				Url      string `json:"url"`
				Standard string `json:"standard"`
			}, i int) *types.ChainExplorer {
				return &types.ChainExplorer{
					Name:     t.Name,
					Url:      t.Url,
					Standard: t.Standard,
				}
			}),
		}
	})
	return lo.Filter(chains, func(t *types.Chain, i int) bool {
		if len(s.config.VisibleNetworks) > 0 {
			return s.config.VisibleNetworks[t.ChainId]
		} else {
			return !s.config.HiddenNetworks[t.ChainId]
		}
	}), nil
}

func (s *Service) GetChains(ctx context.Context) (result []*types.Chain, err error) {
	if s.cachedChains != nil {
		return s.cachedChains, nil
	}
	s.cachedChains, err = s.getChains(ctx)
	return s.cachedChains, err
}

func (s *Service) GetChain(ctx context.Context, chain string) (*types.Chain, error) {
	chains, err := s.GetChains(ctx)
	if err != nil {
		return nil, err
	}
	for _, c := range chains {
		if strings.ToLower(c.ShortName) == strings.ToLower(chain) {
			return c, nil
		}
	}
	return nil, fmt.Errorf("unknown chain (%s)", chain)
}

func (s *Service) GetTotalStaked(ctx context.Context) (string, error) {
	res := big.NewInt(0)
	vals, err := s.staking.GetValidators(&bind.CallOpts{Context: ctx})
	if err != nil {
		return "", err
	}
	for _, val := range vals {
		status, err := s.staking.GetValidatorStatus(&bind.CallOpts{Context: ctx}, val)
		if err != nil {
			return "", err
		}
		res.Add(res, status.TotalDelegated)
	}
	return common2.ToEther(res), nil
}

func (s *Service) GetLatestBlock(ctx context.Context) (uint64, uint64, uint64, error) {
	latestKnownBlockNumber, err := s.eth.BlockByNumber(ctx, nil)
	if err != nil {
		return 0, 0, 0, err
	}
	latestAffectedBlockNumber := s.state.GetLastAffectedBlock(ctx)
	return latestKnownBlockNumber.NumberU64(), latestAffectedBlockNumber, latestKnownBlockNumber.Time(), nil
}

func (s *Service) GetChainConfig() *types.ChainConfig {
	return s.cachedChainConfig
}

func (s *Service) GetValidators(ctx context.Context) (result []*types.Validator, err error) {
	return s.state.GetValidators(ctx)
}

func (s *Service) GetValidatorSlashings(ctx context.Context, validator common.Address, offset, size uint32) (result []*types.ValidatorSlashing, err error) {
	return s.state.GetValidatorSlashings(ctx, validator, int64(offset), int64(size))
}

func (s *Service) GetValidatorHistories(ctx context.Context, validator common.Address, offset, size uint32) (result []*types.ValidatorHistory, err error) {
	return s.state.GetValidatorHistories(ctx, validator, int64(offset), int64(size))
}

func (s *Service) GetValidatorDeposits(ctx context.Context, validator common.Address, offset, size uint32) (result []*types.ValidatorDeposit, err error) {
	return s.state.GetValidatorDeposits(ctx, validator, int64(offset), int64(size))
}

func (s *Service) GetValidatorEvents(ctx context.Context, validator common.Address, offset, size uint32) (result []*types.ValidatorEvent, err error) {
	return s.state.GetValidatorEvents(ctx, validator, int64(offset), int64(size))
}

func (s *Service) GetDelegatorsByValidator(ctx context.Context, validator common.Address, offset, size uint32) (result []*types.Delegator, err error) {
	return s.state.GetDelegators(ctx, int64(offset), int64(size), validator)
}

func (s *Service) GetDelegators(ctx context.Context, offset, size uint32) (result []*types.Delegator, err error) {
	return s.state.GetDelegators(ctx, int64(offset), int64(size), common.Address{})
}
