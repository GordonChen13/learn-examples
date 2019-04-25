package repository

import (
	"context"
	"github.com/GordonChen13/learn-examples/go/cloudNativeGo/ch5/models"
)

type Match interface {
	Store(ctx context.Context, match *models.Match) error
	GetById(ctx context.Context, id string) (*models.Match, error)
	GetByName(ctx context.Context, name string) (*models.Match, error)
}
