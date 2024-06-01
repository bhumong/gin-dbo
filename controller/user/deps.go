package user

import (
	"context"

	"github.com/bhumong/dbo-test/model"
)

type UserService interface {
	CreateUser(ctx context.Context, dto model.UserRequestDTO) (*model.User, error)
	Login(ctx context.Context, dto model.LoginRequestDTO) (*model.LoginResponseDTO, error)
	GetCurrentUser(ctx context.Context, userId string) (*model.User, error)
}
