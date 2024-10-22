package httpxtest

import (
	"encoding/json"
	"log"
	"net/http"
)

func NewHTTPTestHandler(data any, respErr any, code int) http.HandlerFunc {
	return func(w http.ResponseWriter, _ *http.Request) {
		switch {
		case code >= http.StatusBadRequest:
			bb, err := json.Marshal(respErr)
			if err != nil {
				log.Fatal(err)
			}
			w.WriteHeader(code)
			w.Write(bb)

		default:
			bb, err := json.Marshal(data)
			if err != nil {
				log.Fatal(err)
			}
			w.WriteHeader(code)
			w.Write(bb)
		}
	}
}
