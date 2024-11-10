package rpc_client

import (
	"encoding/json"
)

type Request struct {
	jsonrpc string
	id      uint
	method  *string
	params  []any
}

func NewRequest(method string) *Request {
	req := new(Request)

	req.jsonrpc = "2.0"

	// TODO: Check if nil if fesible!?
	req.id = 1
	req.method = &method
	req.params = nil

	return req
}

func (self *Request) SetMethod(method string) *Request {
	self.method = &method
	return self
}

func (self *Request) SetParam(params []any) *Request {
	self.params = params
	return self
}

func (self *Request) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]any{
		"jsonrpc": self.jsonrpc,
		"id":      self.id,
		"method":  self.method,
		"params":  self.params,
	})
}
