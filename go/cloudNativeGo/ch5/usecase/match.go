package usecase

import (
	"context"
	"github.com/GordonChen13/learn-examples/go/cloudNativeGo/ch5/models"
)

type Match interface {
	Create(ctx context.Context, name string) (res *models.Match, err error)
	GetByName(ctx context.Context, name string) (res *models.Match, err error)
	GetById(ctx context.Context, id string) (res *models.Match, err error)
}
