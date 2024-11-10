package rpc_client

type Response struct {
	Body string
}

// TODO: Make builder
func NewResponse(body string) *Response {
	resp := new(Response)
	resp.Body = body
	return resp
}
