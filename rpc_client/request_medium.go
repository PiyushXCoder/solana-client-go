package rpc_client

type RequestMedium interface {
	Send(request_body *Request) (*Response, error)
}
