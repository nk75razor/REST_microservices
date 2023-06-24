package domain

import (
	"github.com/nk75razor/REST_microservices/dto"
	"github.com/nk75razor/REST_microservices/errs"
)

type Customer struct {
	Id          string `db:"customer_ID"`
	Name        string
	City        string
	ZipCode     string `db:"pincode"`
	DateOfBirth string `db:"date_of_birth"`
	Status      string
}

func (c Customer) StatusAsText() string {
	statusAsText := "active"
	if c.Status == "0" {
		statusAsText = "inactive"
	}
	return statusAsText
}
func (c Customer) ToDto() dto.CustomerResponse {

	response := dto.CustomerResponse{
		Id:          c.Id,
		Name:        c.Name,
		City:        c.City,
		ZipCode:     c.ZipCode,
		DateOfBirth: c.DateOfBirth,
		Status:      c.StatusAsText(),
	}
	return response
}

type CustomerRepository interface {
	FindAll(status string) ([]Customer, *errs.AppError)
	ById(string) (*Customer, *errs.AppError)
}
