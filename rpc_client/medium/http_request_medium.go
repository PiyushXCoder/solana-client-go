package medium

import (
	"bytes"
	"io"
	"log"
	"net/http"

	"github.com/PiyushXCoder/solana-client-go/rpc_client"
)

type HttpRequestMedium struct {
	url string
}

func NewHttpRequestMedium(url string) *HttpRequestMedium {
	medium := new(HttpRequestMedium)
	medium.url = url
	return medium
}

func (self HttpRequestMedium) Send(request *rpc_client.Request) (*rpc_client.Response, error) {
	request_json, err := request.MarshalJSON()
	if err != nil {
		return nil, err
	}
	log.Println("[HTTP Request]", string(request_json))

	http_resp, err := http.Post(self.url, "application/json", bytes.NewReader(request_json))
	if err != nil {
		return nil, err
	}

	response_body, err := io.ReadAll(http_resp.Body)
	if err != nil {
		return nil, err
	}

	// TODO: Use builder
	resp, err := rpc_client.NewResponse(response_body)

	return resp, err
}
