package main

import (
	"net/http"

	"github.com/google/wire"

	"KratosNew/internal/biz"
	"KratosNew/internal/data"
	"KratosNew/internal/server"
)

type App struct {
	httpServer *http.Server
}

func newApp(httpServer *http.Server) *App {
	return &App{httpServer: httpServer}
}

func initApp() *App {
	wire.Build(
		data.NewData,
		biz.NewUserUsecase,
		server.NewHTTPServer,
		newApp,
	)

	return nil
}