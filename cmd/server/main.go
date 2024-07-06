package main

import (
	"go.uber.org/fx"

	"github.com/darland/yandex_smart_home_server/internal/app"
	"github.com/darland/yandex_smart_home_server/internal/ports/http"
	"github.com/darland/yandex_smart_home_server/internal/service"
	"github.com/darland/yandex_smart_home_server/pkg/common/fxevent"
	"github.com/darland/yandex_smart_home_server/pkg/common/logger"
)

func main() {
	fx.New(
		fx.WithLogger(
			fxevent.NewZeroLogger(),
		),
		fx.Provide(
			logger.New,
			// Application
			fx.Annotate(
				app.New,
				fx.As(new(service.App)),
			),
		),
		http.NewServer(),
	).Run()
}
