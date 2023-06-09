package domain

import "github.com/nk75razor/REST_microservices/errs"

type Customer struct {
	Id          string
	Name        string
	City        string
	ZipCode     string
	DateOfBirth string
	Status      string
}

type CustomerRepository interface {
	FindAll() ([]Customer, error)
	ById(string) (*Customer, *errs.AppError)
}
