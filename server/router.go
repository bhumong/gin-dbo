package server

import (
	"github.com/bhumong/dbo-test/controller"
	"github.com/bhumong/dbo-test/controller/user"
	"github.com/bhumong/dbo-test/db"
	userRepository "github.com/bhumong/dbo-test/repository/user"
	userService "github.com/bhumong/dbo-test/service/user"

	"github.com/bhumong/dbo-test/middleware"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	router := gin.New()
	gormInstance := db.GetDB()
	router.Use(middleware.DefaultStructuredLogger())
	router.Use(middleware.GlobalErrorHandler())
	router.Use(gin.Recovery())

	health := new(controller.HealthController)

	router.GET("/health", health.Status)

	userRepo := userRepository.New(gormInstance)
	userServ := userService.New(userRepo)
	userController := user.New(userServ)
	router.POST("/users", userController.Create)
	router.POST("/login", userController.Login)

	authGroup := router.Group("", middleware.JwtAuth())
	authGroup.GET("/me", userController.Me)

	return router

}
