package gateway

import (
	"context"
	"github.com/Ankr-network/ankr-protocol/shared/types"
	"github.com/ethereum/go-ethereum/common"
)

func (s *Server) GetValidators(ctx context.Context, req *types.GetValidatorsRequest) (*types.GetValidatorsReply, error) {
	result, err := s.stakingService.GetValidators(ctx)
	if err != nil {
		return nil, err
	}
	return &types.GetValidatorsReply{Validators: result}, nil
}

func (s *Server) GetValidatorHistory(ctx context.Context, req *types.GetValidatorHistoryRequest) (*types.GetValidatorHistoryReply, error) {
	if req.Size_ == 0 || req.Size_ > 1000 {
		req.Size_ = 1000
	}
	result, err := s.stakingService.GetValidatorHistories(ctx, common.HexToAddress(req.Validator), req.Offset, req.Size_)
	if err != nil {
		return nil, err
	}
	newSize, hasMore := checkHasMore(req.Size_, len(result))
	return &types.GetValidatorHistoryReply{ValidatorHistories: result[:newSize], HasMore: hasMore}, nil
}

func (s *Server) GetValidatorSlashings(ctx context.Context, req *types.GetValidatorSlashingsRequest) (*types.GetValidatorSlashingsReply, error) {
	if req.Size_ == 0 || req.Size_ > 1000 {
		req.Size_ = 1000
	}
	result, err := s.stakingService.GetValidatorSlashings(ctx, common.HexToAddress(req.Validator), req.Offset, req.Size_)
	if err != nil {
		return nil, err
	}
	newSize, hasMore := checkHasMore(req.Size_, len(result))
	return &types.GetValidatorSlashingsReply{ValidatorSlashings: result[:newSize], HasMore: hasMore}, nil
}

func (s *Server) GetValidatorDeposits(ctx context.Context, req *types.GetValidatorDepositsRequest) (*types.GetValidatorDepositsReply, error) {
	if req.Size_ == 0 || req.Size_ > 1000 {
		req.Size_ = 1000
	}
	result, err := s.stakingService.GetValidatorDeposits(ctx, common.HexToAddress(req.Validator), req.Offset, req.Size_)
	if err != nil {
		return nil, err
	}
	newSize, hasMore := checkHasMore(req.Size_, len(result))
	return &types.GetValidatorDepositsReply{ValidatorDeposits: result[:newSize], HasMore: hasMore}, nil
}

func (s *Server) GetDelegators(ctx context.Context, req *types.GetDelegatorsRequest) (*types.GetDelegatorsReply, error) {
	var result []*types.Delegator
	var err error
	if req.Validator != "" {
		result, err = s.stakingService.GetDelegatorsByValidator(ctx, common.HexToAddress(req.Validator))
	} else {
		result, err = s.stakingService.GetDelegators(ctx)
	}
	if err != nil {
		return nil, err
	}
	return &types.GetDelegatorsReply{Delegators: result}, nil
}

func (s *Server) GetChains(ctx context.Context, req *types.GetChainsRequest) (*types.GetChainsReply, error) {
	if req.Size_ == 0 || req.Size_ > 1000 {
		req.Size_ = 1000
	}
	if req.Chain != "" {
		chain, err := s.stakingService.GetChain(ctx, req.Chain)
		if err != nil {
			return nil, err
		}
		return &types.GetChainsReply{Chains: []*types.Chain{chain}, HasMore: false}, nil
	}
	result, err := s.stakingService.GetChains(ctx)
	if err != nil {
		return nil, err
	}
	newSize, hasMore := checkHasMore(req.Size_, len(result))
	return &types.GetChainsReply{Chains: result[:newSize], HasMore: hasMore}, nil
}
