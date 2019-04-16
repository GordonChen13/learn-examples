package usecase

import (
	"context"
	"github.com/GordonChen13/learn-examples/go/cloudNativeGo/ch5/models"
)

type match interface {
	create(ctx context.Context, match *models.Match) (res *models.Match, err error)
}
