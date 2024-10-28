package rpc_client

import (
	"github.com/PiyushXCoder/solana-client-go/rpc_client/medium"
)

type Request[Param any] struct {
	request_medium medium.RequestMedium

	json_rpc string
	id       uint
	method   string
	param    []Param
}

func NewRequest() Request[any] {
	return Request[any]{}
}

func (self *Request[any]) SetRequestMedium(request_medium medium.RequestMedium) {
	self.request_medium = request_medium
}

func (self *Request[any]) Send() string {
	return self.request_medium.Send()
}
