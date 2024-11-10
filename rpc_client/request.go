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

func (self *Request) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]any{
		"jsonrpc": self.jsonrpc,
		"id":      self.id,
		"method":  self.method,
		"params":  self.params,
	})
}

type RequestBuilder struct {
	request *Request
}

func NewRequestBuilder() *RequestBuilder {
	reqBuilder := new(RequestBuilder)
	reqBuilder.request = new(Request)

	return reqBuilder
}

func (self *RequestBuilder) Id(id uint) *RequestBuilder {
	self.request.id = id
	return self
}

func (self *RequestBuilder) Method(method string) *RequestBuilder {
	self.request.method = &method
	return self
}

func (self *RequestBuilder) Params(params []any) *RequestBuilder {
	self.request.params = params
	return self
}

func (self RequestBuilder) Build() *Request {
	self.request.jsonrpc = "2.0"

	return self.request
}
