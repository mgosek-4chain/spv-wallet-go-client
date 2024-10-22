package main

import (
	"context"
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
	tt, err := spv.SharedConfig(ctx)
	if err != nil {
		log.Fatal(err)
	}
	exampleutil.Print("HTTP api/v1/config/shared - shared config", tt)
}
