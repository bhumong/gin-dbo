package model

import (
	"time"

	"github.com/bhumong/dbo-test/config"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	BaseModel
	Name     string `gorm:"size:255;not null"`
	Email    string `gorm:"size:255;unique;not null"`
	Password string `gorm:"type:varchar;not null"`
}

func (user *User) GeneratePassword(password string) (string, error) {
	encryptedPassword, err := bcrypt.GenerateFromPassword([]byte(password), config.GetConfig().GetInt("bycript.saltLen"))
	if err != nil {
		return "", err
	}
	return string(encryptedPassword), nil
}

func (user *User) BeforeCreate(tx *gorm.DB) error {
	user.ID = uuid.New()
	password, err := user.GeneratePassword(user.Password)
	if err != nil {
		return err
	}
	user.Password = password
	return nil
}

func (user *User) ToResponseDTO() UserResponseDTO {
	return UserResponseDTO{
		ID:        user.ID.String(),
		Name:      user.Name,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}

type UserRequestDTO struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8"`
}

type LoginRequestDTO struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type LoginResponseDTO struct {
	Token string `json:"token"`
}

type UserUpdateDTO struct {
	Name     *string `json:"name"`
	Password *string `json:"password" binding:"min=8"`
}

type UserResponseDTO struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func (dto *UserRequestDTO) ToModel() User {
	return User{
		Name:     dto.Name,
		Email:    dto.Email,
		Password: dto.Password,
	}
}

type CustomClaim struct {
	jwt.Claims
	UserId string
}
