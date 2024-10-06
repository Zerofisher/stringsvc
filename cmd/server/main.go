package main

import (
	"net/http"
	"os"

	"github.com/Zerofisher/stringsvc/internal/endpoints"
	"github.com/Zerofisher/stringsvc/internal/service"
	"github.com/Zerofisher/stringsvc/internal/transport"
	"github.com/go-kit/log"
)

func main() {
	logger := log.NewLogfmtLogger(os.Stderr)

	svc := service.NewStringService()
	eps := endpoints.MakeEndpoints(svc)
	handler := transport.NewHTTPHandler(eps)

	logger.Log("msg", "HTTP", "addr", ":8080")
	logger.Log("err", http.ListenAndServe(":8080", handler))
}
