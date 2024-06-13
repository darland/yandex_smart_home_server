package yash_operating

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/darland/yandex_smart_home_server/pkg/common/server"
)

func (hg *HandlerGroup) GetRoutes() []server.Route {
	tokenValidatorMiddleware := middleware.KeyAuth(func(key string, c echo.Context) (bool, error) {
		return key == "valid-key", nil
	})

	return []server.Route{
		{
			Method:  http.MethodHead,
			Path:    "/api/v1.0",
			Handler: hg.healthCheck,
		},
		{
			Method:      http.MethodPost,
			Path:        "/v1.0/user/unlink",
			Handler:     hg.unlink,
			Middlewares: []echo.MiddlewareFunc{tokenValidatorMiddleware},
		},
		{
			Method:      http.MethodGet,
			Path:        "/v1.0/user/devices",
			Handler:     hg.devices,
			Middlewares: []echo.MiddlewareFunc{tokenValidatorMiddleware},
		},
		{
			Method:      http.MethodPost,
			Path:        "/v1.0/user/devices/query",
			Handler:     hg.query,
			Middlewares: []echo.MiddlewareFunc{tokenValidatorMiddleware},
		},
		{
			Method:      http.MethodPost,
			Path:        "/v1.0/user/devices/action",
			Handler:     hg.action,
			Middlewares: []echo.MiddlewareFunc{tokenValidatorMiddleware},
		},
	}
}
