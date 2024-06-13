package server

import (
	"go.uber.org/fx"
)

type RouteConfigurator interface {
	GetRoutes() []Route
}

type RouteConfiguratorOut struct {
	fx.Out

	RG RouteConfigurator `group:"route_configurator"`
}

type RouteConfiguratorIn struct {
	fx.In

	RouteConfigurators []RouteConfigurator `group:"route_configurator"`
}

var RouteConfiguratorTag = fx.ResultTags(`group:"route_configurator"`)

func InitRoutes(srv *HttpServer, params RouteConfiguratorIn) {
	for _, rg := range params.RouteConfigurators {
		srv.InitRoutes(rg)
	}
}

func NewRouteConfigurator(rg RouteConfigurator) RouteConfiguratorOut {
	return RouteConfiguratorOut{
		RG: rg,
	}
}

func NewHandlerGroups(groups ...any) fx.Option {
	var opts []fx.Option

	for group := range groups {
		opts = append(
			opts,
			fx.Provide(
				fx.Annotate(
					group,
					RouteConfiguratorTag,
				),
			),
		)
	}

	return fx.Options(opts...)
}
