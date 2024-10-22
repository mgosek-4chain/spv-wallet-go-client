package client

import (
	"context"
	"net/http"
	"time"

	bip32 "github.com/bitcoin-sv/go-sdk/compat/bip32"
	"github.com/bitcoin-sv/spv-wallet-go-client/internal/api/v1/user/configurations"
	"github.com/bitcoin-sv/spv-wallet-go-client/internal/api/v1/user/transactions"
	"github.com/bitcoin-sv/spv-wallet-go-client/internal/httpx"
	"github.com/bitcoin-sv/spv-wallet/models/response"
)

type SPVWallet struct {
	transactionsAPI   *transactions.API
	configurationsAPI *configurations.API
	http              *httpx.HTTP
	addr              string
}

func (s *SPVWallet) Transactions(ctx context.Context, opts ...transactions.QueryBuilderOption) ([]*response.Transaction, error) {
	return s.transactionsAPI.Transactions(ctx, opts...)
}

func (s *SPVWallet) Transaction(ctx context.Context, ID string) (*response.Transaction, error) {
	return s.transactionsAPI.Transaction(ctx, ID)
}

func (s *SPVWallet) UpdateTransactionMetadata(ctx context.Context, r transactions.UpdateTransactionMetadataRequest) (*response.Transaction, error) {
	return s.transactionsAPI.UpdateTransactionMetadata(ctx, r)
}

func (s *SPVWallet) DraftTransaction(ctx context.Context, r transactions.DraftTransactionRequest) (*response.DraftTransaction, error) {
	return s.transactionsAPI.DraftTransaction(ctx, r)
}

func (s *SPVWallet) RecordTransaction(ctx context.Context, r transactions.RecordTransactionRequest) (*response.Transaction, error) {
	return s.transactionsAPI.RecordTransaction(ctx, r)
}

func (s *SPVWallet) SharedConfig(ctx context.Context) (*response.SharedConfig, error) {
	return s.configurationsAPI.SharedConfig(ctx)
}

type Option func(h *SPVWallet) error

func WithXPriv(xPriv string) Option {
	return func(s *SPVWallet) error {
		key, err := bip32.GenerateHDKeyFromString(xPriv)
		if err != nil {
			return err
		}
		s.http.AuthHeaders = &httpx.AuthHeaders{
			ExtendedKey: key,
			Sign:        true,
		}
		return nil
	}
}

func WithXPub(xPub string) Option {
	return func(s *SPVWallet) error {
		key, err := bip32.GetHDKeyFromExtendedPublicKey(xPub)
		if err != nil {
			return err
		}
		s.http.AuthHeaders = &httpx.AuthHeaders{
			ExtendedKey: key,
			Sign:        false,
		}
		return nil
	}
}

func WithAccessKey(key string) Option {
	return func(s *SPVWallet) error {
		pk, err := httpx.PrivateKeyFromHexOrWIF(key)
		if err != nil {
			return err
		}
		s.http.AuthHeaders = &httpx.AuthHeaders{AccessKey: pk}
		return nil
	}
}

func New(opts ...Option) (*SPVWallet, error) {
	spv := SPVWallet{
		addr: "http://localhost:3003/api/v1",
		http: &httpx.HTTP{
			Client: &http.Client{
				Timeout:   time.Minute,
				Transport: http.DefaultTransport,
			},
			AuthHeaders: &httpx.NoopAuthHeaders{},
		},
	}
	for _, o := range opts {
		err := o(&spv)
		if err != nil {
			return nil, err
		}
	}

	spv.transactionsAPI = &transactions.API{
		Addr: spv.addr + "/transactions",
		HTTP: spv.http,
	}
	spv.configurationsAPI = &configurations.API{
		Addr: spv.addr + "/configs",
		HTTP: spv.http,
	}
	return &spv, nil
}
