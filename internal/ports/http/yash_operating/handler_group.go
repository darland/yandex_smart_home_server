package yash_operating

import (
	"github.com/darland/yandex_smart_home_server/internal/service"
)

type HandlerGroup struct {
	app service.App
}

func NewHandlerGroup() *HandlerGroup {
	return &HandlerGroup{}
}
