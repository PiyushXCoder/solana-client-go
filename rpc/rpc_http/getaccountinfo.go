package rpc_http

import (
	"encoding/json"

	"github.com/PiyushXCoder/solana-client-go/rpc_client"
)

type GetAccountInfoConfig struct {
	Commitment string `json:"commitment"`
	Encoding   string `json:"encoding"`
	DataSlice  struct {
		Length uint `json:"length"`
		Offset uint `json:"offset"`
	} `json:"dataSlice"`
	MinContextSlot int `json:"minContextSlot"`
}

type GetAccountInfoResult struct {
	Context struct {
		ApiVersion string `json:"apiVersion"`
		Slot       uint64 `json:"slot"`
	} `json:"context"`
	Value struct {
		Lamports   uint64   `json:"lamports"`
		Owner      string   `json:"owner"`
		Data       []string `json:"data"`
		Executable bool     `json:"executable"`
		RentEpoch  uint64   `json:"rentEpoch"`
		Space      uint64   `json:"space"`
	} `json:"value"`
}

func GetAccountInfo(client *rpc_client.Client, request_id uint, query_account_base58 string, config *GetAccountInfoConfig) (*GetAccountInfoResult, *rpc_client.Response, error) {
	req := rpc_client.NewRequestBuilder().Id(request_id).Method("getAccountInfo").Params([]any{
		query_account_base58,
		config,
	}).Build()

	resp, err := client.Send(req)
	if err != nil {
		return nil, nil, err
	}

	result := new(GetAccountInfoResult)
	err = json.Unmarshal(resp.RawResult, result)

	return result, resp, err
}
