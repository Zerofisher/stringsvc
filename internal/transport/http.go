package transport

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/Zerofisher/stringsvc/internal/endpoints"
	"github.com/Zerofisher/stringsvc/pkg/stringsvc"
	httptransport "github.com/go-kit/kit/transport/http"
)

func NewHTTPHandler(endpoints endpoints.Endpoints) http.Handler {
	m := http.NewServeMux()

	m.Handle("/uppercase", httptransport.NewServer(
		endpoints.UppercaseEndpoint,
		decodeUppercaseRequest,
		encodeResponse,
	))

	m.Handle("/count", httptransport.NewServer(
		endpoints.CountEndpoint,
		decodeCountRequest,
		encodeResponse,
	))

	return m
}

func decodeUppercaseRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request stringsvc.UppercaseRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func decodeCountRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request stringsvc.CountRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func encodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}
