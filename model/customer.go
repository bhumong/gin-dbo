package model

import (
	"math"
	"strconv"
)

type Customer struct {
	BaseModel
	Name      string `gorm:"size:255;not null"`
	Email     string `gorm:"size:255;unique;not null"`
	Phone     string `gorm:"size:25;not null"`
	Address   string `gorm:"size:255;not null"`
	CreatedBy string `gorm:"size:36;not null;index"`
}

type CustomerWithUser struct {
	Customer
	CreatedBy string
}

type CustomerPaginateRequestDTO struct {
	Page       string
	PageSize   string
	FilterName string
}

type CustomerRequestDTO struct {
	Name    string `json:"name" binding:"required"`
	Email   string `json:"email" binding:"required,email"`
	Phone   string `json:"phone" binding:"required,min=8"`
	Address string `json:"address" binding:"required"`
}

type CustomerUpdateRequestDTO struct {
	Name    string `json:"name"`
	Phone   string `json:"phone"`
	Address string `json:"address"`
}

func (dto *CustomerRequestDTO) ToModel() Customer {
	return Customer{
		Name:    dto.Name,
		Email:   dto.Email,
		Phone:   dto.Phone,
		Address: dto.Address,
	}
}

func (dto *CustomerPaginateRequestDTO) ToCustomerPaginateDTO() CustomerPaginateDTO {
	page, err := strconv.Atoi(dto.Page)
	if err != nil || page <= 0 {
		page = 1
	}
	pageSize, err := strconv.Atoi(dto.PageSize)
	if err != nil || pageSize <= 0 {
		pageSize = 10
	}
	if pageSize > 100 {
		pageSize = 100
	}
	return CustomerPaginateDTO{
		Page:                       page,
		PageSize:                   pageSize,
		CustomerPaginateRequestDTO: *dto,
	}
}

type CustomerPaginateDTO struct {
	CustomerPaginateRequestDTO
	Page     int
	PageSize int
}

func (q CustomerPaginateDTO) ToPaginationResponse(total int64) Paginate {
	return Paginate{
		Page:      q.Page,
		PageSize:  q.PageSize,
		PageCount: int(math.Ceil(float64(total) / float64(q.PageSize))),
		Total:     total,
	}
}

type CustomerPaginateResponseDTO struct {
	Customers []CustomerWithUser `json:"customers"`
	Paginate  Paginate           `json:"paginate"`
}
