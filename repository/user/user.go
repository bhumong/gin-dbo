package user

import (
	"context"

	"github.com/bhumong/dbo-test/model"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func New(db *gorm.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (r *UserRepository) CreateUser(ctx context.Context, dto model.UserRequestDTO) (*model.User, error) {
	user := dto.ToModel()
	result := r.db.WithContext(ctx).Create(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func (r *UserRepository) GetUser(ctx context.Context, id string) (*model.User, error) {
	user := model.User{}
	result := r.db.WithContext(ctx).First(&user, "id = ?", id)

	if result.Error != nil {
		return nil, result.Error
	}

	return &user, nil
}

func (r *UserRepository) GetUserByEmail(ctx context.Context, email string) (*model.User, error) {
	user := model.User{}
	result := r.db.WithContext(ctx).First(&user, "email = ?", email)

	if result.Error != nil {
		return nil, result.Error
	}

	return &user, nil
}

func (r *UserRepository) UpdateUser(ctx context.Context, id string, dto model.UserUpdateDTO) error {
	user := model.User{}
	userUpdate := model.User{}

	if dto.Name == nil && dto.Password == nil {
		return nil
	}

	if dto.Name != nil {
		userUpdate.Name = *dto.Name
	}
	if dto.Password != nil {
		password, err := userUpdate.GeneratePassword(*dto.Password)
		if err != nil {
			return err
		}
		userUpdate.Password = password
	}

	result := r.db.Model(&user).Where("id = ?", id).Updates(userUpdate)

	return result.Error
}
