package customer

import (
	"net/http"

	"github.com/bhumong/dbo-test/model"
	errors "github.com/bhumong/dbo-test/pkg/error"
	"github.com/gin-gonic/gin"
)

type CustomerController struct {
	customerService CustomerService
	userService     UserService
}

func New(userService UserService, customerService CustomerService) CustomerController {
	return CustomerController{
		userService:     userService,
		customerService: customerService,
	}
}

func (cc *CustomerController) List(c *gin.Context) {
	claims, isExist := c.Get("jwt.claims")
	if !isExist {
		c.Error(model.ErrAuthClaimsNotFound)
		return
	}
	user, err := cc.userService.GetCurrentUser(c.Request.Context(), claims)
	if err != nil {
		c.Error(err)
		return
	}
	filter := model.CustomerPaginateRequestDTO{
		Page:       c.DefaultQuery("page", "1"),
		PageSize:   c.DefaultQuery("page_size", "10"),
		FilterName: c.Query("name"),
	}
	response, err := cc.customerService.List(c.Request.Context(), filter.ToCustomerPaginateDTO(), user.ID.String())
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(http.StatusOK, response)
}

func (cc *CustomerController) Get(c *gin.Context) {
	claims, isExist := c.Get("jwt.claims")
	if !isExist {
		c.Error(model.ErrAuthClaimsNotFound)
		return
	}
	user, err := cc.userService.GetCurrentUser(c.Request.Context(), claims)
	if err != nil {
		c.Error(err)
		return
	}

	customerId := c.Param("customer_id")
	if customerId == "" {
		c.Error(errors.New("NOT_FOUND", "ID_NOT_PROVIDE", "customer id not provided", http.StatusBadRequest))
		return
	}

	customer, err := cc.customerService.GetCustomer(c.Request.Context(), customerId)
	if err != nil {
		c.Error(err)
		return
	}
	if customer.CreatedBy != user.ID.String() {
		c.Error(errors.New("FORBIDEN", "FORBIDEN_USER", "customer not allowed get by current user", http.StatusForbidden))
	}

	c.JSON(http.StatusOK, customer)
}

func (cc *CustomerController) CreateCustomer(c *gin.Context) {
	claims, isExist := c.Get("jwt.claims")
	if !isExist {
		c.Error(model.ErrAuthClaimsNotFound)
		return
	}

	user, err := cc.userService.GetCurrentUser(c.Request.Context(), claims)
	if err != nil {
		c.Error(err)
		return
	}

	var customerDTO model.CustomerRequestDTO
	err = c.ShouldBindJSON(&customerDTO)
	if err != nil {
		c.Error(err)
		return
	}

	customer, err := cc.customerService.CreateCustomer(c.Request.Context(), customerDTO, user.ID.String())
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, customer)
}

func (cc *CustomerController) UpdateCustomer(c *gin.Context) {
	claims, isExist := c.Get("jwt.claims")
	if !isExist {
		c.Error(model.ErrAuthClaimsNotFound)
		return
	}

	_, err := cc.userService.GetCurrentUser(c.Request.Context(), claims)
	if err != nil {
		c.Error(err)
		return
	}

	customerId := c.Param("customer_id")
	if customerId == "" {
		c.Error(errors.New("NOT_FOUND", "ID_NOT_PROVIDE", "customer id not provided", http.StatusBadRequest))
		return
	}

	var customerDTO model.CustomerUpdateRequestDTO
	err = c.ShouldBindJSON(&customerDTO)
	if err != nil {
		c.Error(err)
		return
	}

	customer, err := cc.customerService.UpdateCustomer(c.Request.Context(), customerDTO, customerId)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, customer)
}

func (cc *CustomerController) DeleteCustomer(c *gin.Context) {
	claims, isExist := c.Get("jwt.claims")
	if !isExist {
		c.Error(model.ErrAuthClaimsNotFound)
		return
	}

	_, err := cc.userService.GetCurrentUser(c.Request.Context(), claims)
	if err != nil {
		c.Error(err)
		return
	}

	customerId := c.Param("customer_id")
	if customerId == "" {
		c.Error(errors.New("NOT_FOUND", "ID_NOT_PROVIDE", "customer id not provided", http.StatusBadRequest))
		return
	}

	err = cc.customerService.DeleteCustomer(c.Request.Context(), customerId)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}
