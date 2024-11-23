package main

import (
	"log"

	"github.com/PiyushXCoder/solana-client-go/rpc/rpc_http"
	"github.com/PiyushXCoder/solana-client-go/rpc_client"
	"github.com/PiyushXCoder/solana-client-go/rpc_client/medium"
)

func main() {
	client := rpc_client.NewClient(medium.NewHttpRequestMedium("https://api.devnet.solana.com"))

	acc, _, _ := rpc_http.GetAccountInfo(client, 1, "vines1vzrYbzLMRdu58ou5XTby4qAqVRLmqo36NKPTg", nil)
	log.Println(acc.Value.Owner)
}
