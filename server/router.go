package server

import (
	"github.com/bhumong/dbo-test/controller"
	customerController "github.com/bhumong/dbo-test/controller/customer"
	userController "github.com/bhumong/dbo-test/controller/user"
	"github.com/bhumong/dbo-test/db"

	customerRepository "github.com/bhumong/dbo-test/repository/customer"
	userRepository "github.com/bhumong/dbo-test/repository/user"

	customerService "github.com/bhumong/dbo-test/service/customer"
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
	userCon := userController.New(userServ)

	v1 := router.Group("/v1")
	v1.POST("/users", userCon.Create)
	v1.POST("/login", userCon.Login)

	authGroup := v1.Group("", middleware.JwtAuth())
	authGroup.GET("/me", userCon.Me)

	customerRepo := customerRepository.New(gormInstance)
	customerServ := customerService.New(customerRepo)
	customerCon := customerController.New(userServ, customerServ)

	cg := authGroup.Group("/customers")
	cg.GET("", customerCon.List)
	cg.GET("/:customer_id", customerCon.Get)
	cg.POST("", customerCon.CreateCustomer)
	cg.PUT("/:customer_id", customerCon.UpdateCustomer)
	cg.DELETE("/:customer_id", customerCon.DeleteCustomer)

	return router

}
