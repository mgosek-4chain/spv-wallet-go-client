package transactions_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/bitcoin-sv/spv-wallet-go-client/internal/api/v1/user/transactions"
	"github.com/bitcoin-sv/spv-wallet-go-client/internal/api/v1/user/transactions/transactionstest"
	"github.com/bitcoin-sv/spv-wallet-go-client/internal/httpx"
	"github.com/bitcoin-sv/spv-wallet-go-client/internal/httpx/httpxtest"
	"github.com/bitcoin-sv/spv-wallet/models"
	"github.com/bitcoin-sv/spv-wallet/models/response"
	"github.com/stretchr/testify/require"
)

func TestTransactionsAPI_Transaction(t *testing.T) {
	tests := map[string]struct {
		code        int
		response    *response.Transaction
		responseErr *models.ResponseError
		expectedErr *models.SPVError
	}{
		"HTTP GET /api/v1/transactions/id response: 200": {
			code:     http.StatusOK,
			response: transactionstest.NewTransaction(),
		},
		"HTTP GET /api/v1/transactions/id response: 401": {
			responseErr: transactionstest.NewResponseError(http.StatusBadRequest),
			expectedErr: transactionstest.NewSPVError(http.StatusBadRequest),
			code:        http.StatusBadRequest,
		},
		"HTTP GET /api/v1/transactions/id response: 400": {
			responseErr: transactionstest.NewResponseError(http.StatusInternalServerError),
			expectedErr: transactionstest.NewSPVError(http.StatusInternalServerError),
			code:        http.StatusInternalServerError,
		},
		"HTTP GET /api/v1/transactions/id response: 500": {
			responseErr: transactionstest.NewResponseError(http.StatusInternalServerError),
			expectedErr: transactionstest.NewSPVError(http.StatusInternalServerError),
			code:        http.StatusInternalServerError,
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			ts := httptest.NewServer(httpxtest.NewHTTPTestHandler(tc.response, tc.responseErr, tc.code))
			defer ts.Close()

			cli := transactions.API{
				Addr: ts.URL,
				HTTP: &httpx.HTTP{
					Client:      ts.Client(),
					AuthHeaders: &httpx.NoopAuthHeaders{},
				},
			}
			ID := "1"
			got, err := cli.Transaction(context.Background(), ID)
			if err != nil {
				require.ErrorIs(t, tc.expectedErr, err)
			}
			if tc.response != nil {
				require.EqualValues(t, tc.response, got)
			}
		})
	}
}

func TestTransactionsAPI_DraftTransction(t *testing.T) {
	tests := map[string]struct {
		code        int
		response    *response.DraftTransaction
		responseErr *models.ResponseError
		expectedErr *models.SPVError
		req         transactions.DraftTransactionRequest
	}{
		"HTTP POST /api/v1/transactions draft response: 201": {
			code:     http.StatusCreated,
			response: transactionstest.NewDraftTransaction(),
			req: transactions.DraftTransactionRequest{
				Config: response.TransactionConfig{
					Outputs: []*response.TransactionOutput{
						{
							Satoshis:     1,
							Script:       "",
							To:           "john.doe.test4@john.doe.test.4chain.space",
							UseForChange: false,
						},
					},
				},
				Metadata: transactions.Metadata{
					"receiver": "john.doe.test4@john.doe.test.4chain.space",
					"sender":   "john.doe.test4@john.doe.test.4chain.space",
				},
			},
		},
		"HTTP POST /api/v1/transactions draft response: 400": {
			responseErr: transactionstest.NewResponseError(http.StatusBadRequest),
			expectedErr: transactionstest.NewSPVError(http.StatusBadRequest),
			code:        http.StatusBadRequest,
		},
		"HTTP POST /api/v1/transactions draft response: 401": {
			responseErr: transactionstest.NewResponseError(http.StatusUnauthorized),
			expectedErr: transactionstest.NewSPVError(http.StatusUnauthorized),
			code:        http.StatusUnauthorized,
		},
		"HTTP POST /api/v1/transactions draft response: 500": {
			responseErr: transactionstest.NewResponseError(http.StatusInternalServerError),
			expectedErr: transactionstest.NewSPVError(http.StatusInternalServerError),
			code:        http.StatusInternalServerError,
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			ts := httptest.NewServer(httpxtest.NewHTTPTestHandler(tc.response, tc.responseErr, tc.code))
			defer ts.Close()

			cli := transactions.API{
				Addr: ts.URL,
				HTTP: &httpx.HTTP{
					Client:      ts.Client(),
					AuthHeaders: &httpx.NoopAuthHeaders{},
				},
			}
			got, err := cli.DraftTransaction(context.Background(), tc.req)
			if err != nil {
				require.ErrorIs(t, tc.expectedErr, err)
			}
			if tc.response != nil {
				require.EqualValues(t, tc.response, got)
			}
		})
	}
}

func TestTransactionsAPI_RecordTransction(t *testing.T) {
	tests := map[string]struct {
		code        int
		req         transactions.RecordTransactionRequest
		response    *response.Transaction
		responseErr *models.ResponseError
		expectedErr *models.SPVError
	}{
		"HTTP POST /api/v1/transactions record response: 200": {
			code:     http.StatusOK,
			response: transactionstest.NewTransaction(),
			req: transactions.RecordTransactionRequest{
				Metadata: transactions.Metadata{
					"key":  "value",
					"key2": "value2",
				},
				Hex:         "0100000002",
				ReferenceID: "b356f7fa00cd3f20cce6c21d704cd13e871d28d714a5ebd0532f5a0e0cde63f7",
			},
		},
		"HTTP POST /api/v1/transactions record response: 400": {
			responseErr: transactionstest.NewResponseError(http.StatusBadRequest),
			expectedErr: transactionstest.NewSPVError(http.StatusBadRequest),
			code:        http.StatusBadRequest,
		},
		"HTTP POST /api/v1/transactions record response: 401": {
			responseErr: transactionstest.NewResponseError(http.StatusUnauthorized),
			expectedErr: transactionstest.NewSPVError(http.StatusUnauthorized),
			code:        http.StatusUnauthorized,
		},
		"HTTP POST /api/v1/transactions record response: 500": {
			responseErr: transactionstest.NewResponseError(http.StatusInternalServerError),
			expectedErr: transactionstest.NewSPVError(http.StatusInternalServerError),
			code:        http.StatusInternalServerError,
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			ts := httptest.NewServer(httpxtest.NewHTTPTestHandler(tc.response, tc.responseErr, tc.code))
			defer ts.Close()

			cli := transactions.API{
				Addr: ts.URL,
				HTTP: &httpx.HTTP{
					Client:      ts.Client(),
					AuthHeaders: &httpx.NoopAuthHeaders{},
				},
			}
			got, err := cli.RecordTransaction(context.Background(), tc.req)
			if err != nil {
				require.ErrorIs(t, tc.expectedErr, err)
			}
			if tc.response != nil {
				require.EqualValues(t, tc.response, got)
			}
		})
	}
}

func TestTransactionsAPI_Transctions(t *testing.T) {
	tests := map[string]struct {
		response    *response.PageModel[response.Transaction]
		code        int
		responseErr *models.ResponseError
		expectedErr *models.SPVError
	}{
		"HTTP GET /api/v1/transactions response: 200": {
			response: transactionstest.NewTransactions(),
			code:     http.StatusOK,
		},
		"HTTP GET /api/v1/transactions response: 401": {
			responseErr: transactionstest.NewResponseError(http.StatusUnauthorized),
			expectedErr: transactionstest.NewSPVError(http.StatusUnauthorized),
			code:        http.StatusUnauthorized,
		},
		"HTTP GET /api/v1/transactions response: 400": {
			responseErr: transactionstest.NewResponseError(http.StatusBadRequest),
			expectedErr: transactionstest.NewSPVError(http.StatusBadRequest),
			code:        http.StatusBadRequest,
		},
		"HTTP GET /api/v1/transactions response: 500": {
			responseErr: transactionstest.NewResponseError(http.StatusInternalServerError),
			expectedErr: transactionstest.NewSPVError(http.StatusInternalServerError),
			code:        http.StatusInternalServerError,
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			ts := httptest.NewServer(httpxtest.NewHTTPTestHandler(tc.response, tc.responseErr, tc.code))
			defer ts.Close()

			cli := transactions.API{
				Addr: ts.URL,
				HTTP: &httpx.HTTP{
					Client:      ts.Client(),
					AuthHeaders: &httpx.NoopAuthHeaders{},
				},
			}
			got, err := cli.Transactions(context.Background())
			if err != nil {
				require.ErrorIs(t, tc.expectedErr, err)
			}
			if tc.response != nil {
				require.EqualValues(t, tc.response.Content, got)
			}
		})
	}
}

func TestTransactionsAPI_UpdateTransactionMetadata(t *testing.T) {
	tests := map[string]struct {
		code        int
		response    *response.Transaction
		responseErr *models.ResponseError
		expectedErr *models.SPVError
		req         transactions.UpdateTransactionMetadataRequest
	}{
		"HTTP PATCH /api/v1/transactions/id response: 200": {
			response: transactionstest.NewTransaction(),
			req: transactions.UpdateTransactionMetadataRequest{
				ID: "1",
				Metadata: transactions.Metadata{
					"example_key1": "example_key1_val",
					"example_key2": "example_key2_val",
				},
			},
			code: http.StatusOK,
		},
		"HTTP PATCH /api/v1/transactions/id response: 401": {
			responseErr: transactionstest.NewResponseError(http.StatusUnauthorized),
			expectedErr: transactionstest.NewSPVError(http.StatusUnauthorized),
			code:        http.StatusUnauthorized,
		},
		"HTTP PATCH /api/v1/transactions/id response: 400": {
			responseErr: transactionstest.NewResponseError(http.StatusBadRequest),
			expectedErr: transactionstest.NewSPVError(http.StatusBadRequest),
			code:        http.StatusBadRequest,
		},
		"HTTP PATCH /api/v1/transactions/id response: 500": {
			responseErr: transactionstest.NewResponseError(http.StatusInternalServerError),
			expectedErr: transactionstest.NewSPVError(http.StatusInternalServerError),
			code:        http.StatusInternalServerError,
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			ts := httptest.NewServer(httpxtest.NewHTTPTestHandler(tc.response, tc.responseErr, tc.code))
			defer ts.Close()

			cli := transactions.API{
				Addr: ts.URL,
				HTTP: &httpx.HTTP{
					Client:      ts.Client(),
					AuthHeaders: &httpx.NoopAuthHeaders{},
				},
			}
			got, err := cli.UpdateTransactionMetadata(context.Background(), tc.req)
			if err != nil {
				require.ErrorIs(t, tc.expectedErr, err)
			}
			if tc.response != nil {
				require.EqualValues(t, tc.response, got)
			}
		})
	}
}
