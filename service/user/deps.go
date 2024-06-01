package user

import (
	"context"

	"github.com/bhumong/dbo-test/model"
)
type UserRepository interface {
	CreateUser(ctx context.Context, dto model.UserRequestDTO) (*model.User, error)
	GetUser(ctx context.Context, id string) (*model.User, error)
	GetUserByEmail(ctx context.Context, email string) (*model.User, error)
	UpdateUser(ctx context.Context, id string, dto model.UserUpdateDTO) error
}
