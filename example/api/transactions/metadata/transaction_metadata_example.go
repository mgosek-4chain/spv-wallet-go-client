package metadata

import (
	"context"
	"fmt"
	"log"

	client "github.com/bitcoin-sv/spv-wallet-go-client"
	"github.com/bitcoin-sv/spv-wallet-go-client/example/exampleutil"
	"github.com/bitcoin-sv/spv-wallet-go-client/internal/api/transactions"
)

func main() {
	xPriv := "xprv9s21ZrQH143K3xqwhES2ujnTYsbCq79ZD2WX9WzZpJB3ACnYAaRQSxmPf6KxYFFVPqg9XKg3b3tx4gCkhPmmMrskZhHUfKZL3kUDXyfkSx4"
	opt := client.WithXPub(xPriv)
	spv, err := client.New(opt)
	if err != nil {
		log.Fatal(err)
	}

	ID := "e2a500ca70d0195ec09ae9faf5b1e17528ef9ad530f8c56972b18a7634bb6c5a"
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
