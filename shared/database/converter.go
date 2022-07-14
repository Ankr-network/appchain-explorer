package database

import (
	"fmt"
	"github.com/Ankr-network/ankr-protocol/shared/entity"
	"github.com/Ankr-network/ankr-protocol/shared/types"
	"strconv"
)

func blocksToProto(blocks []*entity.Block) (result []*types.Block) {
	for _, b := range blocks {
		result = append(result, blockToProto(b))
	}
	return
}

func blockToProto(block *entity.Block) *types.Block {
	if block == nil {
		return nil
	}
	return &types.Block{
		BlockHash:   fmt.Sprintf("0x%x", block.Hash),
		BlockNumber: uint32(block.Number),
		MinerHash:   fmt.Sprintf("0x%x", block.MinerHash),
		Timestamp:   uint64(block.Timestamp.Unix()),
		GasLimit:    uint64(block.GasLimit),
		GasUsed:     uint64(block.GasUsed),
		SizeBytes:   uint32(block.Size.Int64),
	}
}

func blockDetailsToProto(block *entity.Block) *types.BlockDetails {
	if block == nil {
		return nil
	}
	return &types.BlockDetails{
		BlockHash:       fmt.Sprintf("0x%x", block.Hash),
		BlockNumber:     uint32(block.Number),
		MinerHash:       fmt.Sprintf("0x%x", block.MinerHash),
		Timestamp:       uint64(block.Timestamp.Unix()),
		GasLimit:        uint64(block.GasLimit),
		GasUsed:         uint64(block.GasUsed),
		SizeBytes:       uint32(block.Size.Int64),
		ParentHash:      fmt.Sprintf("0x%x", block.ParentHash),
		Difficulty:      uint32(block.Difficulty.Float64),
		TotalDifficulty: uint32(block.TotalDifficulty.Float64),
		Nonce:           fmt.Sprintf("0x%x", block.Nonce),
	}
}

func txsToProto(txs []*entity.Transaction) (result []*types.Transaction) {
	for _, tx := range txs {
		result = append(result, txToProto(tx))
	}
	return
}

func txToProto(tx *entity.Transaction) *types.Transaction {
	if tx == nil {
		return nil
	}
	txType := types.TransactionType_UNKNOWN
	if tx.Input != nil {
		txType = types.TransactionType_CONTRACT_CALL
	}
	return &types.Transaction{
		TxHash:      fmt.Sprintf("0x%x", tx.Hash),
		Status:      types.TransactionStatus(tx.Status.Int64),
		Value:       strconv.FormatFloat(tx.Value/1e18, 'f', 18, 64),
		TxFee:       strconv.FormatFloat(tx.GasUsed.Float64*tx.GasPrice/1e18, 'f', 18, 64),
		BlockNumber: uint64(tx.BlockNumber.Int64),
		Timestamp:   uint32(tx.UpdatedAt.Unix()),
		Error:       tx.Error.String,
		Sender:      fmt.Sprintf("0x%x", tx.FromAddressHash),
		Recipient:   fmt.Sprintf("0x%x", tx.ToAddressHash),
		TxType:      txType,
	}
}

func txDetailsToProto(tx *entity.Transaction) *types.TransactionDetails {
	if tx == nil {
		return nil
	}
	return &types.TransactionDetails{
		TxHash: fmt.Sprintf("0x%x", tx.Hash),
		Status: types.TransactionStatus(tx.Status.Int64),
		//Gas:               uint64(tx.Gas),
		//CumulativeGasUsed: uint64(tx.CumulativeGasUsed.Float64),
		GasUsed:  uint64(tx.GasUsed.Float64),
		GasPrice: uint64(tx.GasPrice),
		//TxIndex:           uint64(tx.Index.Int64),
		//Input:             tx.Input,
		Nonce: uint64(tx.Nonce),
		Value: strconv.FormatFloat(tx.Value, 'f', 18, 64),
		Error: tx.Error.String,
		//BlockHash:         tx.BlockHash,
		//BlockNumber:       uint64(tx.BlockNumber.Int64),
		Sender:    fmt.Sprintf("0x%x", tx.FromAddressHash),
		Recipient: fmt.Sprintf("0x%x", tx.ToAddressHash),
		//Contract:          tx.CreatedContractAddressHash,
		//RevertReason:      tx.RevertReason.String,
		Type: uint32(tx.Type.Int64),
		//InternalFailed:    tx.HasErrorInInternalTxs.Bool,
		Timestamp: uint64(tx.UpdatedAt.Unix()),
	}
}

func tokenTransferToProto(transfer *entity.TokenTransfer) *types.TokenTransfer {
	if transfer == nil {
		return nil
	}
	return &types.TokenTransfer{
		TxHash:        transfer.TransactionHash,
		AddressFrom:   transfer.FromAddressHash,
		AddressTo:     transfer.ToAddressHash,
		TokenContract: transfer.TokenContractAddressHash,
		Amount:        uint64(transfer.Amount.Float64),
		BlockNumber:   uint64(transfer.BlockNumber.Int64),
		Timestamp:     uint64(transfer.UpdatedAt.Unix()),
	}
}

func tokenTransfersToProto(transfers []*entity.TokenTransfer) (result []*types.TokenTransfer) {
	for _, b := range transfers {
		result = append(result, tokenTransferToProto(b))
	}
	return
}

func addressToProto(address *entity.Address) (result *types.Address) {
	if address == nil {
		return nil
	}
	return &types.Address{
		Hash:         fmt.Sprintf("0x%x", address.Hash),
		Balance:      uint64(address.FetchedCoinBalance.Float64),
		Transactions: uint32(address.TransactionsCount.Int64),
		GasUsed:      uint64(address.GasUsed.Int64),
	}
}
