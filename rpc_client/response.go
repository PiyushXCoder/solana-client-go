package rpc_client

import (
	"encoding/json"
)

type Response struct {
	Jsonrpc   string          `json:"jsonrpc"`
	Id        uint            `json:"id"`
	RawResult json.RawMessage `json:"result"`
}

// TODO: Make builder
func NewResponse(response_body []byte) (*Response, error) {
	resp := new(Response)

	err := json.Unmarshal(response_body, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (self Response) Result() {

}
