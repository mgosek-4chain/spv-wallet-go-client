package httpx_test

import (
	"context"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/bitcoin-sv/spv-wallet-go-client/internal/httpx"
	"github.com/bitcoin-sv/spv-wallet-go-client/internal/httpx/httpxtest"
	"github.com/stretchr/testify/require"
)

func TestHTTP_Do(t *testing.T) {
	tests := map[string]struct {
		code        int
		request     *RequestBody
		response    *ResponseBody
		responseErr *ResponseBody
	}{
		"HTTP DO - POST - response: 200": {
			code: http.StatusOK,
			request: &RequestBody{
				Content: "hello world",
			},
			response: &ResponseBody{
				Content: "hello world",
			},
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			ts := httptest.NewServer(httpxtest.NewHTTPTestHandler(tc.response, tc.responseErr, tc.code))
			defer ts.Close()

			cli := &httpx.HTTP{
				Client:      ts.Client(),
				AuthHeaders: &httpx.NoopAuthHeaders{},
			}
			res, err := cli.Do(context.Background(), http.MethodPost, ts.URL, tc.request)
			defer res.Body.Close()

			require.EqualValues(t, tc.response, Decode(res.Body))
			require.Nil(t, err)
		})
	}
}

type RequestBody struct {
	Content string `json:"content"`
}

type ResponseBody struct {
	Content string `json:"content"`
}

func Decode(r io.Reader) *ResponseBody {
	var dst ResponseBody
	dec := json.NewDecoder(r)
	err := dec.Decode(&dst)
	if err != nil {
		log.Fatal(err)
	}
	return &dst
}
