package user

import (
	"context"
	"net/http"

	"github.com/bhumong/dbo-test/model"
	errors "github.com/bhumong/dbo-test/pkg/error"
	"github.com/bhumong/dbo-test/utils"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	userRepo UserRepository
}

func New(userRepo UserRepository) *UserService {
	return &UserService{
		userRepo: userRepo,
	}
}

func (s *UserService) CreateUser(ctx context.Context, dto model.UserRequestDTO) (*model.User, error) {
	existUser, _ := s.userRepo.GetUserByEmail(ctx, dto.Email)

	if existUser != nil {
		return nil, errors.New("EMAIL_ALREADY_EXIST", "EMAIL_ALREADY_EXIST", "user already exist", http.StatusBadRequest)
	}

	user, err := s.userRepo.CreateUser(ctx, dto)
	if err != nil {
		return nil, errors.New("INT_USER_CNT", "USER CANNOT CREATE", "internal: user cannot create", http.StatusBadRequest).
			SetCause(err)
	}
	return user, err
}

func (s *UserService) Login(ctx context.Context, dto model.LoginRequestDTO) (*model.LoginResponseDTO, error) {
	existUser, err := s.userRepo.GetUserByEmail(ctx, dto.Email)
	if err != nil || existUser == nil {
		return nil, errors.New("USER_NOT_FOUND", "USER_NOT_FOUND", "user not found", http.StatusBadRequest).
			SetCause(err)
	}

	err = bcrypt.CompareHashAndPassword([]byte(existUser.Password), []byte(dto.Password))
	if err == bcrypt.ErrMismatchedHashAndPassword {
		return nil,
			errors.New("INVALID_PASSWORD", "INVALID_PASSWORD", "password given is invalid", http.StatusBadRequest).
				SetCause(err)
	}
	if err != nil {
		return nil,
			errors.New("BCRYPT_FAIL", "BCRYPT_FAIL", "bycrypt failed check hash", http.StatusInternalServerError).
				SetCause(err)
	}

	authJwt := utils.GetJwt()
	tokenString, err := authJwt.NewWithClaims(jwt.MapClaims{
		"userId": existUser.ID,
	})
	if err != nil {
		return nil,
			errors.New("CANNOT_GENERATE_TOKEN", "CANNOT_GENERATE_TOKEN", "cannot generate token", http.StatusInternalServerError).
				SetCause(err)
	}
	return &model.LoginResponseDTO{
		Token: tokenString,
	}, nil
}

func (s *UserService) GetCurrentUser(ctx context.Context, userId string) (*model.User, error) {
	existUser, err := s.userRepo.GetUser(ctx, userId)
	if err != nil || existUser == nil {
		return nil, errors.New("USER_NOT_FOUND", "USER_NOT_FOUND", "user not found", http.StatusNotFound).
			SetCause(err)
	}
	return existUser, nil
}
