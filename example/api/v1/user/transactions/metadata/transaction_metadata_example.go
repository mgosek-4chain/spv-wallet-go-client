package main

import (
	"context"
	"fmt"
	"log"

	client "github.com/bitcoin-sv/spv-wallet-go-client"
	"github.com/bitcoin-sv/spv-wallet-go-client/example/exampleutil"
	"github.com/bitcoin-sv/spv-wallet-go-client/internal/api/v1/user/transactions"
)

func main() {
	xPriv := ""
	opt := client.WithXPub(xPriv)
	spv, err := client.New(opt)
	if err != nil {
		log.Fatal(err)
	}

	ID := ""
	t, err := spv.UpdateTransactionMetadata(context.Background(), transactions.UpdateTransactionMetadataRequest{
		ID: ID,
		Metadata: transactions.Metadata{
			"example_key_101": "example_key_101_value",
		},
	})
	if err != nil {
		log.Fatal(err)
	}
	exampleutil.Print(fmt.Sprintf("HTTP api/v1/transaction - ID -  %s - metadata update", ID), t)
}
