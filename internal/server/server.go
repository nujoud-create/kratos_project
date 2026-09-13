package server

import (
	"net/http"

	"github.com/google/wire"
)

// ProviderSet is server providers.
var ProviderSet = wire.NewSet(
	newHTTPServer,
)

func newHTTPServer() *http.Server {
	return &http.Server{
		Addr: ":8000",
	}
}
