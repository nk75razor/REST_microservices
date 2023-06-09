package domain

import (
	"database/sql"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/nk75razor/REST_microservices/errs"
)

type CustomerRepositoryDb struct {
	client *sql.DB
}

func (d *CustomerRepositoryDb) FindAll() ([]Customer, error) {

	FindAllSql := "select customer_ID , name ,  city , pincode , date_of_birth, status from customers"
	rows, err := (*d).client.Query(FindAllSql)
	if err != nil {
		log.Println("Error while quering the customer table " + err.Error())
		return nil, err
	}
	customers := make([]Customer, 0)
	for rows.Next() {
		var c Customer
		rows.Scan(&c.Id, &c.Name, &c.DateOfBirth, &c.City, &c.ZipCode, &c.Status)
		if err != nil {
			log.Println("Error while scanning the customer table " + err.Error())
			return nil, err
		}
		customers = append(customers, c)
	}
	return customers, nil
}

func (d CustomerRepositoryDb) ById(id string) (*Customer, *errs.AppError) {
	CustomerSql := "select customer_ID , name , city , pincode , date_of_birth , status from customers"
	row := d.client.QueryRow(CustomerSql, id)
	var c Customer
	err := row.Scan(&c.Id, &c.Name, &c.City, &c.ZipCode, &c.DateOfBirth, &c.Status)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errs.NewNotFoundError("Customer not found")
		} else {
			log.Println("Error while scannin customer", err.Error())
			return nil, errs.NewUnexpectedError("unexpected database Error")
		}

	}
	return &c, nil
}

func NewCustomerRepositoryDB() *CustomerRepositoryDb {
	client, err := sql.Open("mysql", "root:12345678@/rest_tutorial")
	if err != nil {
		panic(err)
	}

	// See "Important settings" section.
	client.SetConnMaxLifetime(time.Minute * 3)
	client.SetMaxOpenConns(10)
	client.SetMaxIdleConns(10)
	return &CustomerRepositoryDb{client}
}

//mysql://root:12345678@127.0.0.1?statusColor=005392&env=local&name=Localhost&tLSMode=0&usePrivateKey=false&safeModeLevel=0&advancedSafeModeLevel=0&driverVersion=0
