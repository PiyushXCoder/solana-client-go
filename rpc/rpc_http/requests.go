// This file needs to be translated to new syntax

package rpc_http

import (
	"github.com/PiyushXCoder/solana-client-go/rpc_client"
)

func GetBalance(request_id uint, addount_id string) *rpc_client.Request {
	return rpc_client.NewRequestBuilder().Id(request_id).Method("getBalance").Params([]any{
		addount_id,
	}).Build()
}

func GetBlock(request_id uint, slot_number uint64) *rpc_client.Request {
	return rpc_client.NewRequestBuilder().Id(request_id).Method("getBlock").Params([]any{
		slot_number,
	}).Build()
}

func GetBlockCommitment(request_id uint, block_number uint64) *rpc_client.Request {
	return rpc_client.NewRequestBuilder().Id(request_id).Method("getBlockCommitment").Params([]any{
		block_number,
	}).Build()
}

func GetBlockHeight(request_id uint) *rpc_client.Request {
	return rpc_client.NewRequestBuilder().Id(request_id).Method("getBlockHeight").Build()
}

func GetBlockProduction(request_id uint) *rpc_client.Request {
	return rpc_client.NewRequestBuilder().Id(request_id).Method("getBlockProduction").Build()
}

func GetBlockTime(request_id uint, block_number uint64) *rpc_client.Request {
	return rpc_client.NewRequestBuilder().Id(request_id).Method("getBlockTime").Params([]any{
		block_number,
	}).Build()
}

func GetBlocks(request_id uint, start_slot_number uint64, end_slot_number uint64) *rpc_client.Request {
	return rpc_client.NewRequestBuilder().Id(request_id).Method("getBlocks").Params([]any{
		start_slot_number, end_slot_number,
	}).Build()
}

func GetBlocksWithLimit(request_id uint, start_slot_number uint64, end_slot_number uint64) *rpc_client.Request {
	return rpc_client.NewRequestBuilder().Id(request_id).Method("getBlocksWithLimit").Params([]any{
		start_slot_number, end_slot_number,
	}).Build()
}

func GetClusterNodes(request_id uint) *rpc_client.Request {
	return rpc_client.NewRequestBuilder().Id(request_id).Method("getClusterNodes").Build()
}

func GetEpochInfo(request_id uint) *rpc_client.Request {
	return rpc_client.NewRequestBuilder().Id(request_id).Method("getEpochInfo").Build()
}

func GetEpochSchedule(request_id uint) *rpc_client.Request {
	return rpc_client.NewRequestBuilder().Id(request_id).Method("getEpochSchedule").Build()
}

func GetFeeForMessage(request_id uint, message_base64 string) *rpc_client.Request {
	return rpc_client.NewRequestBuilder().Id(request_id).Method("getBlocksWithLimit").Params([]any{
		message_base64,
	}).Build()
}

func GetFirstAvailableBlock(request_id uint) *rpc_client.Request {
	return rpc_client.NewRequestBuilder().Id(request_id).Method("getFirstAvailableBlock").Build()
}

func GetGenesisHash(request_id uint) *rpc_client.Request {
	return rpc_client.NewRequestBuilder().Id(request_id).Method("getGenesisHash").Build()
}

func GetHealth(request_id uint) *rpc_client.Request {
	return rpc_client.NewRequestBuilder().Id(request_id).Method("getHealth").Build()
}

func GetHighestSnapshotSlot(request_id uint) *rpc_client.Request {
	return rpc_client.NewRequestBuilder().Id(request_id).Method("getHighestSnapshotSlot").Build()
}

func GetIdentity(request_id uint) *rpc_client.Request {
	return rpc_client.NewRequestBuilder().Id(request_id).Method("getIdentity").Build()
}

func GetInflationGovernor(request_id uint) *rpc_client.Request {
	return rpc_client.NewRequestBuilder().Id(request_id).Method("getInflationGovernor").Build()
}

func GetInflationRate(request_id uint) *rpc_client.Request {
	return rpc_client.NewRequestBuilder().Id(request_id).Method("getInflationRate").Build()
}

func GetInflationReward(request_id uint) *rpc_client.Request {
	return rpc_client.NewRequestBuilder().Id(request_id).Method("getInflationReward").Build()
}

func GetLargestAccounts(request_id uint) *rpc_client.Request {
	return rpc_client.NewRequestBuilder().Id(request_id).Method("getLargestAccounts").Build()
}

func GetLatestBlockhash(request_id uint) *rpc_client.Request {
	return rpc_client.NewRequestBuilder().Id(request_id).Method("getLatestBlockhash").Build()
}

func GetLeaderSchedule(request_id uint) *rpc_client.Request {
	return rpc_client.NewRequestBuilder().Id(request_id).Method("getLeaderSchedule").Build()
}

func GetMaxRetransmitSlot(request_id uint) *rpc_client.Request {
	return rpc_client.NewRequestBuilder().Id(request_id).Method("getMaxRetransmitSlot").Build()
}

func GetMaxShredInsertSlot(request_id uint) *rpc_client.Request {
	return rpc_client.NewRequestBuilder().Id(request_id).Method("getMaxShredInsertSlot").Build()
}

func GetMinimumBalanceForRentExemption(request_id uint) *rpc_client.Request {
	return rpc_client.NewRequestBuilder().Id(request_id).Method("getMinimumBalanceForRentExemption").Build()
}

func GetMultipleAccounts(request_id uint, accounts_public_key []string) *rpc_client.Request {
	return rpc_client.NewRequestBuilder().Id(request_id).Method("getMultipleAccounts").Params([]any{
		accounts_public_key,
	}).Build()
}

func GetProgramAccounts(request_id uint, program_account_public_key []string) *rpc_client.Request {
	return rpc_client.NewRequestBuilder().Id(request_id).Method("getProgramAccounts").Params([]any{
		program_account_public_key,
	}).Build()
}

func GetRecentPerformanceSamples(request_id uint) *rpc_client.Request {
	return rpc_client.NewRequestBuilder().Id(request_id).Method("getRecentPerformanceSamples").Build()
}

func GetRecentPrioritizationFees(request_id uint) *rpc_client.Request {
	return rpc_client.NewRequestBuilder().Id(request_id).Method("getRecentPrioritizationFees").Build()
}

func GetSignatureStatuses(request_id uint, transaction_signatures []string) *rpc_client.Request {
	return rpc_client.NewRequestBuilder().Id(request_id).Method("getSignatureStatuses").Params([]any{
		transaction_signatures,
	}).Build()
}

func GetSignaturesForAddress(request_id uint, account_base58 string) *rpc_client.Request {
	return rpc_client.NewRequestBuilder().Id(request_id).Method("getSignaturesForAddress").Params([]any{
		account_base58,
	}).Build()
}

func GetSlot(request_id uint) *rpc_client.Request {
	return rpc_client.NewRequestBuilder().Id(request_id).Method("getSlot").Build()
}

func GetSlotLeader(request_id uint) *rpc_client.Request {
	return rpc_client.NewRequestBuilder().Id(request_id).Method("getSlotLeader").Build()
}

func GetSlotLeaders(request_id uint, start_slot_number uint64, end_slot_number uint64) *rpc_client.Request {
	return rpc_client.NewRequestBuilder().Id(request_id).Method("getSlotLeaders").Params([]any{
		start_slot_number, end_slot_number,
	}).Build()
}

func GetStakeMinimumDelegation(request_id uint) *rpc_client.Request {
	return rpc_client.NewRequestBuilder().Id(request_id).Method("getStakeMinimumDelegation").Build()
}

func GetSupply(request_id uint) *rpc_client.Request {
	return rpc_client.NewRequestBuilder().Id(request_id).Method("getSupply").Build()
}

func GetTokenAccountBalance(request_id uint, account_public_key_base58 string) *rpc_client.Request {
	return rpc_client.NewRequestBuilder().Id(request_id).Method("getTokenAccountBalance").Params([]any{
		account_public_key_base58,
	}).Build()
}

func GetTokenAccountsByDelegate(request_id uint, account_public_key_base58 string) *rpc_client.Request {
	return rpc_client.NewRequestBuilder().Id(request_id).Method("getTokenAccountsByDelegate").Params([]any{
		account_public_key_base58,
	}).Build()
}

func GetTokenAccountsByOwner(request_id uint, account_public_key_base58 string) *rpc_client.Request {
	return rpc_client.NewRequestBuilder().Id(request_id).Method("getTokenAccountsByOwner").Params([]any{
		account_public_key_base58,
	}).Build()
}

func GetTokenLargestAccounts(request_id uint, account_public_key_base58 string) *rpc_client.Request {
	return rpc_client.NewRequestBuilder().Id(request_id).Method("getTokenLargestAccounts").Params([]any{
		account_public_key_base58,
	}).Build()
}

func GetTokenSupply(request_id uint, account_public_key_base58 string) *rpc_client.Request {
	return rpc_client.NewRequestBuilder().Id(request_id).Method("getTokenSupply").Params([]any{
		account_public_key_base58,
	}).Build()
}

func GetTransaction(request_id uint, transaction_signature_base58 string) *rpc_client.Request {
	return rpc_client.NewRequestBuilder().Id(request_id).Method("getTransaction").Params([]any{
		transaction_signature_base58,
	}).Build()
}

func GetTransactionCount(request_id uint) *rpc_client.Request {
	return rpc_client.NewRequestBuilder().Id(request_id).Method("getTransactionCount").Params([]any{}).Build()
}

func GetVersion(request_id uint) *rpc_client.Request {
	return rpc_client.NewRequestBuilder().Id(request_id).Method("getVersion").Params([]any{}).Build()
}

func GetVoteAccounts(request_id uint) *rpc_client.Request {
	return rpc_client.NewRequestBuilder().Id(request_id).Method("getVoteAccounts").Params([]any{}).Build()
}

func IsBlockhashValid(request_id uint, block_hash_base58 string) *rpc_client.Request {
	return rpc_client.NewRequestBuilder().Id(request_id).Method("isBlockhashValid").Params([]any{
		block_hash_base58,
	}).Build()
}

func MinimumLedgerSlot(request_id uint) *rpc_client.Request {
	return rpc_client.NewRequestBuilder().Id(request_id).Method("minimumLedgerSlot").Params([]any{}).Build()
}

func RequestAirdrop(request_id uint, account_public_key_base58 string) *rpc_client.Request {
	return rpc_client.NewRequestBuilder().Id(request_id).Method("requestAirdrop").Params([]any{
		account_public_key_base58,
	}).Build()
}

func SendTransaction(request_id uint, fully_signed_transaction_base58 string) *rpc_client.Request {
	return rpc_client.NewRequestBuilder().Id(request_id).Method("sendTransaction").Params([]any{
		fully_signed_transaction_base58,
	}).Build()
}

func SimulateTransaction(request_id uint, transaction_base58 string) *rpc_client.Request {
	return rpc_client.NewRequestBuilder().Id(request_id).Method("sendTransaction").Params([]any{
		transaction_base58,
	}).Build()
}
