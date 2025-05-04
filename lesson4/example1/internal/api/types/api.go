package types

import (
	"context"

	"github.com/alinurmyrzakhanov/lessons/lesson4/example1/configs"
)

type Api interface {
	Start(ctx context.Context, cancel context.CancelFunc)
	Config() *configs.ApiConfig
}
