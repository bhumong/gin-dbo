package customer

import (
	"context"

	"github.com/bhumong/dbo-test/model"
	"github.com/bhumong/dbo-test/util/querybuilder"
	"gorm.io/gorm"
)

type CustomerRepository struct {
	db *gorm.DB
}

type Count struct {
	Count int64
}

func New(db *gorm.DB) *CustomerRepository {
	return &CustomerRepository{
		db: db,
	}
}

func (r *CustomerRepository) Paginate(ctx context.Context, dto model.CustomerPaginateDTO, userId string) ([]model.CustomerWithUser, int64, error) {
	var customers []model.CustomerWithUser
	var count int64

	queryCount := r.db.Table("customers c").
		Select("count(c.id)").
		Joins("left join users u on c.created_by = u.id")

	query := r.db.WithContext(ctx).
		Scopes(querybuilder.Paginate(dto.Page, dto.PageSize)).
		Table("customers c").
		Select("c.id, c.name, c.email, c.phone, c.address, u.name created_by, c.created_at, c.updated_at").
		Joins("left join users u on c.created_by = u.id")

	if dto.FilterName != "" {
		queryCount.Where("c.name = ?", dto.FilterName)
		query.Where("c.name = ?", dto.FilterName)
	}

	resultCount := queryCount.Scan(&count)
	if resultCount.Error != nil {
		return nil, 0, resultCount.Error
	}
	result := query.Scan(&customers)
	if result.Error != nil {
		return nil, 0, result.Error
	}

	return customers, count, nil
}

func (r *CustomerRepository) GetCustomer(ctx context.Context, customerId string) (*model.Customer, error) {
	customer := model.Customer{}
	result := r.db.WithContext(ctx).First(&customer, "id = ?", customerId)

	if result.Error != nil {
		return nil, result.Error
	}

	return &customer, nil
}

func (r *CustomerRepository) GetCustomerByEmail(ctx context.Context, email string) (*model.Customer, error) {
	customer := model.Customer{}
	result := r.db.WithContext(ctx).First(&customer, "email = ?", email)

	if result.Error != nil {
		return nil, result.Error
	}

	return &customer, nil
}

func (r *CustomerRepository) CreateCustomer(ctx context.Context, dto model.CustomerRequestDTO, userId string) (*model.Customer, error) {
	customer := dto.ToModel()
	customer.CreatedBy = userId

	result := r.db.WithContext(ctx).Create(&customer)
	if result.Error != nil {
		return nil, result.Error
	}

	return &customer, nil
}

func (r *CustomerRepository) UpdateCustomer(ctx context.Context, customer model.Customer) error {
	result := r.db.WithContext(ctx).Model(&customer).Updates(customer)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *CustomerRepository) DeleteCustomer(ctx context.Context, customerId string) error {
	result := r.db.WithContext(ctx).Table("customers").Where("id = ?", customerId).Delete(customerId)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
