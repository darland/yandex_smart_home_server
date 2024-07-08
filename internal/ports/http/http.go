package http

import (
	"go.uber.org/fx"

	yashOperating "github.com/darland/yandex_smart_home_server/internal/ports/http/yash_operating"
	"github.com/webdevelop-pro/go-common/server"
)

// NewServer returns a new instance of Server
func NewServer() fx.Option {
	return fx.Options(
		server.NewHandlerGroups(
			yashOperating.NewHandlerGroup,
		),
		server.InitAndRun(),
	)
}
