package main

import (
	"go.uber.org/fx"

	"github.com/darland/yandex_smart_home_server/internal/app"
	"github.com/darland/yandex_smart_home_server/internal/ports/http"
	"github.com/darland/yandex_smart_home_server/pkg/common/fxevent"
)

func main() {
	fx.New(
		fx.WithLogger(
			fxevent.NewZeroLogger(),
		),
		fx.Provide(
			// Application
			app.New,
		),
		http.NewServer(),
	).Run()
}
