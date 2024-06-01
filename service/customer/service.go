package customer

import (
	"context"
	"net/http"

	"github.com/bhumong/dbo-test/model"
	errors "github.com/bhumong/dbo-test/pkg/error"
)

type CustomerService struct {
	customerRepo CustomerRepository
}

func New(customerRepo CustomerRepository) *CustomerService {
	return &CustomerService{
		customerRepo: customerRepo,
	}
}

func (s *CustomerService) List(ctx context.Context, dto model.CustomerPaginateDTO, userId string) (*model.CustomerPaginateResponseDTO, error) {
	parsedDto := dto.ToCustomerPaginateDTO()
	customers, total, err := s.customerRepo.Paginate(ctx, parsedDto, userId)
	if err != nil {
		return nil,
			errors.New("FAILED_PAGINATE", "FAILED_PAGINATE", "fail to get customers data", http.StatusInternalServerError)
	}
	return &model.CustomerPaginateResponseDTO{
		Customers: customers,
		Paginate:  dto.ToPaginationResponse(total),
	}, nil
}

func (s *CustomerService) GetCustomer(ctx context.Context, customerId string) (*model.Customer, error) {
	customer, err := s.customerRepo.GetCustomer(ctx, customerId)
	if err != nil {
		return nil,
			errors.New("CUSTOMER_NOT_FOUND", "customer not found", "customer not found by current id", http.StatusNotFound)
	}

	return customer, nil
}

func (s *CustomerService) CreateCustomer(ctx context.Context, dto model.CustomerRequestDTO, userId string) (*model.Customer, error) {
	existCustomer, _ := s.customerRepo.GetCustomerByEmail(ctx, dto.Email)

	if existCustomer != nil {
		return nil, errors.New("EMAIL_ALREADY_EXIST", "EMAIL_ALREADY_EXIST", "customer email already exist", http.StatusBadRequest)
	}

	customer, err := s.customerRepo.CreateCustomer(ctx, dto, userId)

	if err != nil {
		return nil,
			errors.New("CUSTOMER_NOT_FOUND", "customer not found", "customer not found by current id", http.StatusInternalServerError)
	}
	return customer, nil
}

func (s *CustomerService) UpdateCustomer(ctx context.Context, dto model.CustomerUpdateRequestDTO, customerId string) (*model.Customer, error) {
	customer, err := s.GetCustomer(ctx, customerId)
	if err != nil {
		return nil, err
	}
	if dto.Name != "" {
		customer.Name = dto.Name
	}
	if dto.Phone != "" {
		customer.Phone = dto.Phone
	}
	if dto.Address != "" {
		customer.Address = dto.Address
	}
	err = s.customerRepo.UpdateCustomer(ctx, *customer)
	if err != nil {
		return nil, errors.New("CANNOT_UPDATE_CUSTOMER", "CANNOT_UPDATE_CUSTOMER", "failed to update customer", http.StatusInternalServerError)
	}

	return customer, nil
}

func (s *CustomerService) DeleteCustomer(ctx context.Context, customerId string) error {
	_, err := s.GetCustomer(ctx, customerId)
	if err != nil {
		return err
	}

	err = s.customerRepo.DeleteCustomer(ctx, customerId)
	if err != nil {
		return errors.
			New("CANNOT_SOFT_DELETE_CUSTOMER", "CANNOT_SOFT_DELETE_CUSTOMER", "failed to soft delete customer", http.StatusInternalServerError).
			SetCause(err)
	}
	return nil
}
