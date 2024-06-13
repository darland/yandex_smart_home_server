package yash_operating

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// https://yandex.ru/dev/dialogs/smart-home/doc/en/reference/check
func (hg *HandlerGroup) healthCheck(e echo.Context) error {
	return e.String(http.StatusOK, "OK")
}

// https://yandex.ru/dev/dialogs/smart-home/doc/en/reference/unlink
func (hg *HandlerGroup) unlink(e echo.Context) error {
	return e.String(http.StatusOK, "OK")
}

// https://yandex.ru/dev/dialogs/smart-home/doc/en/reference/get-devices
func (hg *HandlerGroup) devices(e echo.Context) error {
	return e.String(http.StatusOK, "OK")
}

// https://yandex.ru/dev/dialogs/smart-home/doc/en/reference/post-devices-query
func (hg *HandlerGroup) query(e echo.Context) error {
	return e.String(http.StatusOK, "OK")
}

// https://yandex.ru/dev/dialogs/smart-home/doc/en/reference/post-action
func (hg *HandlerGroup) action(e echo.Context) error {
	return e.String(http.StatusOK, "OK")
}
