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
	t, err := spv.RecordTransaction(context.Background(), transactions.RecordTransactionRequest{
		Metadata:    map[string]any{},
		Hex:         "b356f7fa00cd3f20cce6c21d704cd13e871d28d714a5ebd0532f5a0e0cde63f7",
		ReferenceID: ID,
	})
	if err != nil {
		log.Fatal(err)
	}
	exampleutil.Print(fmt.Sprintf("HTTP api/v1/transaction - ID -  %s - metadata update", ID), t)
}
