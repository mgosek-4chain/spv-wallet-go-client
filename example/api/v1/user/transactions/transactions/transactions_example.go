package main

import (
	"context"
	"log"

	client "github.com/bitcoin-sv/spv-wallet-go-client"
	"github.com/bitcoin-sv/spv-wallet-go-client/example/exampleutil"
)

func main() {
	xPriv := ""
	opt := client.WithXPriv(xPriv)
	spv, err := client.New(opt)
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()
	tt, err := spv.Transactions(ctx)
	if err != nil {
		log.Fatal(err)
	}
	exampleutil.Print("HTTP api/v1/transactions - all transactions", tt)
}
