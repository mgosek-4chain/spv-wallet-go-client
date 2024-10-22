package main

import (
	"context"
	"fmt"
	"log"

	client "github.com/bitcoin-sv/spv-wallet-go-client"
	"github.com/bitcoin-sv/spv-wallet-go-client/example/exampleutil"
)

func main() {
	xPriv := "xprv9s21ZrQH143K3xqwhES2ujnTYsbCq79ZD2WX9WzZpJB3ACnYAaRQSxmPf6KxYFFVPqg9XKg3b3tx4gCkhPmmMrskZhHUfKZL3kUDXyfkSx4"
	opt := client.WithXPub(xPriv)
	spv, err := client.New(opt)
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()

	ID := "e2a500ca70d0195ec09ae9faf5b1e17528ef9ad530f8c56972b18a7634bb6c5a"
	t, err := spv.Transaction(ctx, ID)
	if err != nil {
		log.Fatal(err)
	}
	exampleutil.Print(fmt.Sprintf("HTTP api/v1/transaction - ID: %s", ID), t)
}
