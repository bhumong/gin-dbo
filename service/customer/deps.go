package customer

import (
	"context"

	"github.com/bhumong/dbo-test/model"
)

type CustomerRepository interface {
	Paginate(ctx context.Context, dto model.CustomerPaginateDTO, userId string) ([]model.CustomerWithUser, int64, error)
	GetCustomer(ctx context.Context, customerId string) (*model.Customer, error)
	GetCustomerByEmail(ctx context.Context, email string) (*model.Customer, error)
	CreateCustomer(ctx context.Context, dto model.CustomerRequestDTO, userId string) (*model.Customer, error)
	UpdateCustomer(ctx context.Context, dto model.Customer) error
	DeleteCustomer(ctx context.Context, customerId string) error
}
