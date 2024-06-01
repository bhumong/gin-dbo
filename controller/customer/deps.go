package customer

import (
	"context"

	"github.com/bhumong/dbo-test/model"
)

type CustomerService interface {
	List(ctx context.Context, dto model.CustomerPaginateDTO, userId string) (*model.CustomerPaginateResponseDTO, error)
	GetCustomer(ctx context.Context, customerId string) (*model.Customer, error)
	CreateCustomer(ctx context.Context, dto model.CustomerRequestDTO, userId string) (*model.Customer, error)
	UpdateCustomer(ctx context.Context, dto model.CustomerUpdateRequestDTO, customerId string) (*model.Customer, error)
	DeleteCustomer(ctx context.Context, customerId string) error
}

type UserService interface {
	GetCurrentUser(ctx context.Context, claims any) (*model.User, error)
}
