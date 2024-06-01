package user

import (
	"net/http"

	"github.com/bhumong/dbo-test/model"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type UserController struct {
	userService UserService
}

func New(userService UserService) UserController {
	return UserController{
		userService: userService,
	}
}

func (uc *UserController) Create(c *gin.Context) {
	var userDTO model.UserRequestDTO
	err := c.ShouldBindJSON(&userDTO)
	if err != nil {
		c.Error(err)
		return
	}
	user, err := uc.userService.CreateUser(c.Request.Context(), userDTO)

	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(http.StatusOK, user.ToResponseDTO())
}

func (uc *UserController) Login(c *gin.Context) {
	var loginDTO model.LoginRequestDTO
	err := c.ShouldBindJSON(&loginDTO)
	if err != nil {
		c.Error(err)
		return
	}
	response, err := uc.userService.Login(c.Request.Context(), loginDTO)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, response)
}

func (uc *UserController) Me(c *gin.Context) {
	claims, isExist := c.Get("jwt.claims")
	if !isExist {
		c.Error(model.ErrAuthClaimsNotFound)
		return
	}
	jwtClaim, ok := claims.(jwt.MapClaims)
	if !ok {
		c.Error(model.ErrAuthClaimsNotFound)
		return
	}
	userId, ok := jwtClaim["userId"].(string)
	if !ok {
		c.Error(model.ErrAuthClaimsNotFound)
		return
	}
	user, err := uc.userService.GetCurrentUser(c.Request.Context(), userId)
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(http.StatusOK, user.ToResponseDTO())
}
