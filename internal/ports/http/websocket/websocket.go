package websocket

import (
	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"
)

var (
	upgrader              = websocket.Upgrader{}
	errSomethingWentWrong = echo.NewHTTPError(500, "something went wrong")
)

func (hg *HandlerGroup) websocket(c echo.Context) error {
	ctx := c.Request().Context()

	conn, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		log.Ctx(ctx).Warn().Err(err).Msg("failed to upgrade connection")

		return err
	}

	err = hg.app.NewWSClient(ctx, conn)
	if err != nil {
		log.Ctx(ctx).Warn().Err(err).Msg("failed to create new ws client")

		return c.JSON(500, errSomethingWentWrong)
	}

	return nil
}
