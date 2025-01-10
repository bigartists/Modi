//go:build wireinject
// +build wireinject

package inject

import (
	"github.com/bigartists/Modi/client"
	"github.com/bigartists/Modi/cmd/server"
	"github.com/bigartists/Modi/src/controllers"
	"github.com/bigartists/Modi/src/informer"
	"github.com/bigartists/Modi/src/middlewares"
	"github.com/bigartists/Modi/src/repo"
	"github.com/bigartists/Modi/src/routes"
	"github.com/bigartists/Modi/src/service"
	"github.com/google/wire"
)

func InitializeApp() (*server.App, error) {
	wire.Build(
		client.ProvideDB,
		repo.ProviderSetRepo,
		service.ProviderSetService,
		controllers.ProviderSetCtrl,
		middlewares.MiddlewareSets,
		informer.InformerSet,
		routes.ProvideRouter,
		server.ProvideApp,
	)
	return &server.App{}, nil
}
