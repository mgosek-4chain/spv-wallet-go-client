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
	xPriv := "xprv9s21ZrQH143K3xqwhES2ujnTYsbCq79ZD2WX9WzZpJB3ACnYAaRQSxmPf6KxYFFVPqg9XKg3b3tx4gCkhPmmMrskZhHUfKZL3kUDXyfkSx4"
	opt := client.WithXPub(xPriv)
	spv, err := client.New(opt)
	if err != nil {
		log.Fatal(err)
	}

	ID := "765fe2a795e80c2a6a6c6e27bb1044e8d2bac2562f13f358dc577e7e072bea6b"
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
