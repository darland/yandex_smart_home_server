package service

import (
	"context"

	"github.com/gorilla/websocket"
)

type App interface {
	NewWSClient(ctx context.Context, conn *websocket.Conn) error
}
