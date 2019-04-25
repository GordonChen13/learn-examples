package impl

import (
	"github.com/GordonChen13/learn-examples/go/cloudNativeGo/ch5/models"
	"github.com/GordonChen13/learn-examples/go/cloudNativeGo/ch5/repository"
	"golang.org/x/net/context"
)

type MatchUseCase struct {
	Repository repository.Match
}

func NewMatchUseCase(repository repository.Match) *MatchUseCase {
	return &MatchUseCase{
		Repository:repository,
	}
}

func (m *MatchUseCase) Create(ctx context.Context, name string) (*models.Match, error) {
	id := models.NewMatchId()
	match := &models.Match{
		Id:id,
		Name:name,
		Moves:nil,
	}
	err := m.Repository.Store(ctx, match)

	if err != nil {
		return nil,err
	}

	return match, nil
}


