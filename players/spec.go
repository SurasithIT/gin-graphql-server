package players

import (
	"context"
	"surasithit/gin-graphql-server/players/model"
)

//go:generate mockgen -source=./store.go -destination=./mocks/store.go -package=mocks "surasithit/gin-graphql-server" Store

type Store interface {
	FindAll(ctx context.Context) ([]*model.Player, error)
	FindByIDs(ctx context.Context, IDs []int) ([]*model.Player, error)
	FindByTeamId(ctx context.Context, teamId int) ([]*model.Player, error)
	FindOne(ctx context.Context, id int) (*model.Player, error)
	Create(ctx context.Context, team *model.Player) (*model.Player, error)
	Update(ctx context.Context, id int, team *model.Player) (*model.Player, error)
	Delete(ctx context.Context, id int) error
}
