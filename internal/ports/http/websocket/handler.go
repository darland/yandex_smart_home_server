package websocket

import (
	"github.com/darland/yandex_smart_home_server/internal/service"
	"github.com/webdevelop-pro/go-common/server"
)

type HandlerGroup struct {
	app service.App
}

func (hg *HandlerGroup) GetRoutes() []server.Route {
	return []server.Route{}
}
