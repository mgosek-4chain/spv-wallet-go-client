package main

import (
	"context"
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
	tt, err := spv.SharedConfig(ctx)
	if err != nil {
		log.Fatal(err)
	}
	exampleutil.Print("HTTP api/v1/config/shared - shared config", tt)
}
