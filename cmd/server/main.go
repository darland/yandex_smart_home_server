package main

import (
	"go.uber.org/fx"

	"github.com/darland/yandex_smart_home_server/internal/app"
	"github.com/darland/yandex_smart_home_server/internal/ports/http"
	"github.com/darland/yandex_smart_home_server/internal/service"
	"github.com/webdevelop-pro/go-common/logger"
	// "github.com/webdevelop-pro/go-common/logger/fxevent"
)

func main() {
	fx.New(
		/*
			fx.WithLogger(
				fxevent.NewZeroLogger(),
			),
		*/
		fx.Provide(
			logger.NewLogger,
			// Application
			fx.Annotate(
				app.New,
				fx.As(new(service.App)),
			),
		),
		http.NewServer(),
	).Run()
}
