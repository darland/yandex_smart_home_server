package app

import (
	"context"

	"github.com/darland/yandex_smart_home_server/internal/domain/dto"
)

type App struct {
}

func New() *App {
	return &App{}
}

func (a App) Unlink(ctx context.Context) error {
	//TODO implement me
	panic("implement me")
}

func (a App) GetDevices(ctx context.Context) (*dto.Payload, error) {
	//TODO implement me
	panic("implement me")
}

func (a App) Query(ctx context.Context, request dto.QueryRequest) (*dto.Payload, error) {
	//TODO implement me
	panic("implement me")
}

func (a App) Action(ctx context.Context, request dto.ActionRequest) (*dto.Payload, error) {
	//TODO implement me
	panic("implement me")
}
