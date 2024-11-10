package rpc_http

import (
	"github.com/PiyushXCoder/solana-client-go/rpc_client"
)

func GetAccountInfo(request_id uint, addount_id string) *rpc_client.Request {
	return rpc_client.NewRequestBuilder().Id(request_id).Method("getAccountInfo").Params([]any{
		addount_id,
		map[string]string{
			"encoding": "base58",
		},
	}).Build()
}
