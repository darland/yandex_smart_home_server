package service

import (
	"context"

	"github.com/darland/yandex_smart_home_server/internal/domain/dto"
)

type App interface {
	Unlink(context.Context) error
	GetDevices(context.Context) (*dto.Payload, error)
	Query(context.Context, dto.QueryRequest) (*dto.Payload, error)
	Action(context.Context, dto.ActionRequest) (*dto.Payload, error)
}
