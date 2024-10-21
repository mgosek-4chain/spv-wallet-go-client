package configurations

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/bitcoin-sv/spv-wallet-go-client/internal/httpx"
	"github.com/bitcoin-sv/spv-wallet/models/response"
)

type HTTP interface {
	Get(ctx context.Context, URL string) (*http.Response, error)
}

type API struct {
	Addr string
	HTTP HTTP
}

func (a *API) SharedConfig(ctx context.Context) (*response.SharedConfig, error) {
	URL := fmt.Sprintf("%s/shared", a.Addr)
	res, err := a.HTTP.Get(ctx, URL)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	if res.StatusCode >= http.StatusBadRequest {
		return nil, httpx.SpvError(res)
	}

	var val *response.SharedConfig
	dec := json.NewDecoder(res.Body)
	if err := dec.Decode(&val); err != nil {
		return nil, fmt.Errorf("json - new decoder - decode op failure: %w", err)
	}
	return val, nil
}
