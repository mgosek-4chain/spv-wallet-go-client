package main

import (
	"context"
	"fmt"
	"log"

	client "github.com/bitcoin-sv/spv-wallet-go-client"
	"github.com/bitcoin-sv/spv-wallet-go-client/example/exampleutil"
	"github.com/bitcoin-sv/spv-wallet-go-client/internal/api/v1/user/transactions"
	"github.com/bitcoin-sv/spv-wallet/models/filter"
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

	bh := uint64(865365)
	t, err := spv.Transactions(ctx, transactions.WithTransactionFilter(filter.TransactionFilter{
		BlockHeight: &bh,
	}))
	if err != nil {
		log.Fatal(err)
	}
	exampleutil.Print(fmt.Sprintf("HTTP api/v1/transactions filtered - all transactions with block height: %d", bh), t)
}
