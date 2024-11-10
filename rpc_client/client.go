package rpc_client

type Client struct {
	request_medium RequestMedium
}

func NewClient(request_medium RequestMedium) *Client {
	client := new(Client)
	client.request_medium = request_medium
	return client
}

func (self Client) Send(req *Request) (*Response, error) {
	return self.request_medium.Send(req)
}
