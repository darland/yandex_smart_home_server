package yash_operating

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/webdevelop-pro/go-common/server/route"
)

func (hg HandlerGroup) GetRoutes() []route.Route {
	tokenValidatorMiddleware := middleware.KeyAuth(func(key string, c echo.Context) (bool, error) {
		return key == "valid-key", nil
	})

	return []route.Route{
		{
			Method: http.MethodHead,
			Path:   "/api/v1.0",
			Handle: hg.healthCheck,
		},
		{
			Method:      http.MethodPost,
			Path:        "/v1.0/user/unlink",
			Handle:      hg.unlink,
			Middlewares: []echo.MiddlewareFunc{tokenValidatorMiddleware},
		},
		{
			Method:      http.MethodGet,
			Path:        "/v1.0/user/devices",
			Handle:      hg.devices,
			Middlewares: []echo.MiddlewareFunc{tokenValidatorMiddleware},
		},
		{
			Method:      http.MethodPost,
			Path:        "/v1.0/user/devices/query",
			Handle:      hg.query,
			Middlewares: []echo.MiddlewareFunc{tokenValidatorMiddleware},
		},
		{
			Method:      http.MethodPost,
			Path:        "/v1.0/user/devices/action",
			Handle:      hg.action,
			Middlewares: []echo.MiddlewareFunc{tokenValidatorMiddleware},
		},
	}
}
