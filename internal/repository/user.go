package repository

import (
	"github.com/hyuti/pocketbase-clean-template/ent"
	"github.com/hyuti/pocketbase-clean-template/pkg/entity/model"
)

type (
	urr = *ent.UserReadRepository
	ucr = *ent.UserCreateRepository
	uur = *ent.UserUpdateRepository
	udr = *ent.UserDeleteRepository
)

type userRepository struct {
	urr
	ucr
	uur
	udr
}

func NewUserRepository(client *ent.Client) ModelRepository[
	*model.User, *model.UserOrderInput, *model.UserWhereInput, *model.UserCreateInput, *model.UserUpdateInput,
] {
	if client == nil {
		panic("client is required")
	}
	return &userRepository{
		ent.NewUserReadRepository(client),
		ent.NewUserCreateRepository(client, false),
		ent.NewUserUpdateRepository(client, false),
		ent.NewUserDeleteRepository(client, false),
	}
}
