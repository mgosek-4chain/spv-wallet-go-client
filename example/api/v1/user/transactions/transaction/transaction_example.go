package main

import (
	"context"
	"fmt"
	"log"

	client "github.com/bitcoin-sv/spv-wallet-go-client"
	"github.com/bitcoin-sv/spv-wallet-go-client/example/exampleutil"
)

func main() {
	xPriv := ""
	opt := client.WithXPub(xPriv)
	spv, err := client.New(opt)
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()

	ID := ""
	t, err := spv.Transaction(ctx, ID)
	if err != nil {
		log.Fatal(err)
	}
	exampleutil.Print(fmt.Sprintf("HTTP api/v1/transaction - ID: %s", ID), t)
}
