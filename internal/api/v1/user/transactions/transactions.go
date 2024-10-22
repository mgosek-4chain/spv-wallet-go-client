package transactions

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/bitcoin-sv/spv-wallet-go-client/internal/httpx"
	"github.com/bitcoin-sv/spv-wallet/models/response"
)

type Metadata map[string]any

type RecordTransactionRequest struct {
	Metadata    Metadata `json:"metadata"`
	Hex         string   `json:"hex"`
	ReferenceID string   `json:"referenceId"`
}

type DraftTransactionRequest struct {
	Config   response.TransactionConfig `json:"config"`
	Metadata Metadata                   `json:"metadata"`
}

type UpdateTransactionMetadataRequest struct {
	ID       string   `json:"-"`
	Metadata Metadata `json:"metadata"`
}

type HTTP interface {
	Get(ctx context.Context, URL string) (*http.Response, error)
	Patch(ctx context.Context, URL string, body any) (*http.Response, error)
	Post(ctx context.Context, URL string, body any) (*http.Response, error)
}

type API struct {
	Addr string
	HTTP HTTP
}

func (a *API) RecordTransaction(ctx context.Context, r RecordTransactionRequest) (*response.Transaction, error) {
	res, err := a.HTTP.Post(ctx, a.Addr, r)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	if res.StatusCode >= http.StatusBadRequest {
		return nil, httpx.SpvError(res)
	}

	var val *response.Transaction
	dec := json.NewDecoder(res.Body)
	if err := dec.Decode(&val); err != nil {
		return nil, fmt.Errorf("json - new decoder - decode op failure: %w", err)
	}
	return val, nil
}

func (a *API) DraftTransaction(ctx context.Context, r DraftTransactionRequest) (*response.DraftTransaction, error) {
	URL := fmt.Sprintf("%s/drafts", a.Addr)
	res, err := a.HTTP.Post(ctx, URL, r)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	if res.StatusCode >= http.StatusBadRequest {
		return nil, httpx.SpvError(res)
	}

	var val *response.DraftTransaction
	dec := json.NewDecoder(res.Body)
	if err := dec.Decode(&val); err != nil {
		return nil, fmt.Errorf("json - new decoder - decode op failure: %w", err)
	}
	return val, nil
}

func (a *API) UpdateTransactionMetadata(ctx context.Context, r UpdateTransactionMetadataRequest) (*response.Transaction, error) {
	URL := fmt.Sprintf("%s/%s", a.Addr, r.ID)
	res, err := a.HTTP.Patch(ctx, URL, r)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	if res.StatusCode >= http.StatusBadRequest {
		return nil, httpx.SpvError(res)
	}

	var val *response.Transaction
	dec := json.NewDecoder(res.Body)
	if err := dec.Decode(&val); err != nil {
		return nil, fmt.Errorf("json - new decoder - decode op failure: %w", err)
	}
	return val, nil
}

func (a *API) Transaction(ctx context.Context, ID string) (*response.Transaction, error) {
	URL := fmt.Sprintf("%s/%s", a.Addr, ID)
	res, err := a.HTTP.Get(ctx, URL)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	if res.StatusCode >= http.StatusBadRequest {
		return nil, httpx.SpvError(res)
	}

	var val *response.Transaction
	dec := json.NewDecoder(res.Body)
	if err := dec.Decode(&val); err != nil {
		return nil, fmt.Errorf("json - new decoder - decode op failure: %w", err)
	}
	return val, nil
}

func (a *API) Transactions(ctx context.Context, opts ...QueryBuilderOption) ([]*response.Transaction, error) {
	params, err := NewQueryBuilder(opts...).Build()
	if err != nil {
		return nil, err
	}
	URL := fmt.Sprintf("%s/?%s", a.Addr, params.Encode())
	res, err := a.HTTP.Get(ctx, URL)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	if res.StatusCode >= http.StatusBadRequest {
		return nil, httpx.SpvError(res)
	}

	var val response.PageModel[response.Transaction]
	dec := json.NewDecoder(res.Body)
	if err := dec.Decode(&val); err != nil {
		return nil, fmt.Errorf("json - new decoder - decode op failure: %w", err)
	}
	return val.Content, nil
}
