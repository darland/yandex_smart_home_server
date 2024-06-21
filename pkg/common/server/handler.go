package server

import (
	"go.uber.org/fx"
)

type RouteConfigurator interface {
	GetRoutes() []Route
}

type RouteConfiguratorIn struct {
	fx.In

	RouteConfigurators []RouteConfigurator `group:"route_configurator"`
}

func InitHandlerGroups(srv *HttpServer, rg RouteConfiguratorIn) {
	for _, group := range rg.RouteConfigurators {
		srv.InitRoutes(group)
	}
}

func NewHandlerGroups(groups ...any) fx.Option {
	var annotates []any

	for _, group := range groups {
		annotates = append(
			annotates,
			fx.Annotate(
				group,
				fx.ResultTags(`group:"route_configurator"`),
				fx.As(new(RouteConfigurator)),
			),
		)
	}

	return fx.Provide(annotates...)
}
