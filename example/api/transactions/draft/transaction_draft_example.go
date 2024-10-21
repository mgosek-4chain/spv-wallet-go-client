package main

import (
	"context"
	"fmt"
	"log"

	client "github.com/bitcoin-sv/spv-wallet-go-client"
	"github.com/bitcoin-sv/spv-wallet-go-client/example/exampleutil"
	"github.com/bitcoin-sv/spv-wallet-go-client/internal/api/transactions"
	"github.com/bitcoin-sv/spv-wallet/models/response"
)

func main() {
	xPriv := "xprv9s21ZrQH143K3xqwhES2ujnTYsbCq79ZD2WX9WzZpJB3ACnYAaRQSxmPf6KxYFFVPqg9XKg3b3tx4gCkhPmmMrskZhHUfKZL3kUDXyfkSx4"
	opt := client.WithXPub(xPriv)
	spv, err := client.New(opt)
	if err != nil {
		log.Fatal(err)
	}

	t, err := spv.DraftTransaction(context.Background(), transactions.DraftTransactionRequest{
		Config: response.TransactionConfig{
			ChangeStrategy:             "",
			ChangeMinimumSatoshis:      0,
			ChangeNumberOfDestinations: 0,
			ChangeSatoshis:             0,
			ExpiresIn:                  0,
			Fee:                        0,
			Outputs: []*response.TransactionOutput{
				{
					Satoshis:     1,
					Script:       "",
					To:           "michal.gosek.test4@mgosek.test.4chain.space",
					UseForChange: false,
				},
			},
		},
		Metadata: transactions.Metadata{
			"receiver": "michal.gosek.test4@mgosek.test.4chain.space",
			"sender":   "michal.gosek.test4@mgosek.test.4chain.space",
		},
	})
	if err != nil {
		log.Fatal(err)
	}
	exampleutil.Print(fmt.Sprintf("HTTP api/v1/transactions - draft - ID: %s", t.ID), t)
}
