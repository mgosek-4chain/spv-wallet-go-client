package httpx

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/bitcoin-sv/spv-wallet/models"
)

type AuthHeaderAppender interface {
	Append(req *http.Header) error
}

type HTTP struct {
	Client      *http.Client
	AuthHeaders AuthHeaderAppender
}

func (h *HTTP) Get(ctx context.Context, URL string) (*http.Response, error) {
	return h.Do(ctx, http.MethodGet, URL, nil)
}

func (h *HTTP) Patch(ctx context.Context, URL string, body any) (*http.Response, error) {
	return h.Do(ctx, http.MethodPatch, URL, body)
}

func (h *HTTP) Post(ctx context.Context, URL string, body any) (*http.Response, error) {
	return h.Do(ctx, http.MethodPost, URL, body)
}

func (h *HTTP) Do(ctx context.Context, method, URL string, body any) (*http.Response, error) {
	var rbody io.Reader
	if body != nil {
		bb, err := json.Marshal(body)
		if err != nil {
			return nil, err
		}
		rbody = bytes.NewBuffer(bb)
	}

	req, err := http.NewRequestWithContext(ctx, method, URL, rbody)
	if err != nil {
		return nil, fmt.Errorf("http - new request with context op failure: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")
	err = h.AuthHeaders.Append(&req.Header)
	if err != nil {
		return nil, fmt.Errorf("auth header appender - append op failure: %w", err)
	}
	return h.Client.Do(req)
}

func SpvError(res *http.Response) error {
	var dst models.SPVError
	if err := json.NewDecoder(res.Body).Decode(&dst); err != nil {
		return err
	}
	dst.StatusCode = res.StatusCode
	return dst
}
