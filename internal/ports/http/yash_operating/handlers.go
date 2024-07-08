package yash_operating

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/darland/yandex_smart_home_server/internal/domain/dto"
)

// https://yandex.ru/dev/dialogs/smart-home/doc/en/reference/check
func (hg *HandlerGroup) healthCheck(e echo.Context) error {
	return e.String(http.StatusOK, "OK")
}

// https://yandex.ru/dev/dialogs/smart-home/doc/en/reference/unlink
func (hg *HandlerGroup) unlink(e echo.Context) error {
	xRequestID := e.Request().Header.Get("X-Request-ID")

	err := hg.app.Unlink(e.Request().Context())
	if err != nil {
		return e.JSON(http.StatusInternalServerError, echo.ErrInternalServerError)
	}

	return e.JSON(http.StatusOK, map[string]string{
		"request_id": xRequestID,
	})
}

// https://yandex.ru/dev/dialogs/smart-home/doc/en/reference/get-devices
func (hg *HandlerGroup) devices(e echo.Context) error {
	payload, err := hg.app.GetDevices(e.Request().Context())
	if err != nil {
		return e.JSON(http.StatusInternalServerError, echo.ErrInternalServerError)
	}

	return e.JSON(http.StatusOK, &dto.DevicesResponse{
		Payload:   payload,
		RequestID: e.Request().Header.Get("X-Request-ID"),
	})
}

// https://yandex.ru/dev/dialogs/smart-home/doc/en/reference/post-devices-query
func (hg *HandlerGroup) query(e echo.Context) error {
	var request dto.QueryRequest
	if err := e.Bind(&request); err != nil {
		return e.JSON(http.StatusBadRequest, echo.ErrBadRequest)
	}

	response, err := hg.app.Query(e.Request().Context(), request)
	if err != nil {
		return e.JSON(http.StatusInternalServerError, echo.ErrInternalServerError)
	}

	return e.JSON(http.StatusOK, &dto.DevicesResponse{
		RequestID: e.Request().Header.Get("X-Request-ID"),
		Payload:   response,
	})
}

// https://yandex.ru/dev/dialogs/smart-home/doc/en/reference/post-action
func (hg *HandlerGroup) action(e echo.Context) error {
	var request dto.ActionRequest
	if err := e.Bind(&request); err != nil {
		return e.JSON(http.StatusBadRequest, echo.ErrBadRequest)
	}

	response, err := hg.app.Action(e.Request().Context(), request)
	if err != nil {
		return e.JSON(http.StatusInternalServerError, echo.ErrInternalServerError)
	}

	return e.JSON(http.StatusOK, &dto.DevicesResponse{
		RequestID: e.Request().Header.Get("X-Request-ID"),
		Payload:   response,
	})
}
