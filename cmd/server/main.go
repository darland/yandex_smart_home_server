package main

import (
	"context"

	"github.com/darland/yandex_smart_home_server/internal/app"
	yashOperating "github.com/darland/yandex_smart_home_server/internal/ports/http/yash_operating"
	"github.com/webdevelop-pro/go-common/configurator"
	"github.com/webdevelop-pro/go-common/server"
	"github.com/webdevelop-pro/go-common/server/route"

	"go.uber.org/fx"
)

func main() {
	// log := logger.NewComponentLogger(context.Background(), "fx")
	fx.New(
		// fx.Logger(log),
		fx.Provide(
			// Application
			context.Background,
			configurator.NewConfigurator,
			app.New,
			server.New,
			yashOperating.NewHandlerGroup,
			// converts
			func(rte *yashOperating.HandlerGroup) route.ConfiguratorIn {
				return route.ConfiguratorIn{
					Configurators: []route.Configurator{
						rte,
					},
				}
			},
			// middleware.NewAuthNoneMW(),
		),
		fx.Invoke(
			// Registration routes and handlers for http server
			server.InitAllRoutes,
			// Run HTTP server
			server.StartServer,
		),
	).Run()
}
