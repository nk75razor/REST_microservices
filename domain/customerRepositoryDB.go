package domain

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/nk75razor/REST_microservices/errs"
	"github.com/nk75razor/REST_microservices/logger"
)

type CustomerRepositoryDb struct {
	client *sqlx.DB
}

func (d *CustomerRepositoryDb) FindAll(status string) ([]Customer, *errs.AppError) {
	//var rows *sql.Rows
	var err error
	customers := make([]Customer, 0)
	if status == "" {
		FindAllSql := "select customer_ID , name ,  city , pincode , date_of_birth, status from customers"
		err = (d).client.Select(&customers, FindAllSql)
		//rows, err = (*d).client.Query(FindAllSql)
	} else {
		FindAllSql := "select customer_ID , name ,  city , pincode , date_of_birth, status from customers where status = ?"
		//rows, err = (*d).client.Query(FindAllSql, status)
		err = d.client.Select(&customers, FindAllSql, status)
	}

	if err != nil {
		logger.Error("Error while quering the customer table " + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected database error")
	}

	// err = sqlx.StructScan(rows, &customers)
	// if err != nil {
	// 	logger.Error("Error while scanning the customer table " + err.Error())
	// 	return nil, errs.NewUnexpectedError("Unexpectd database error")
	// }
	// for rows.Next() {
	// 	var c Customer
	// 	rows.Scan(&c.Id, &c.Name, &c.DateOfBirth, &c.City, &c.ZipCode, &c.Status)
	// 	if err != nil {
	// 		logger.Error("Error while scanning the customer table " + err.Error())
	// 		return nil, errs.NewUnexpectedError("Unexpectd database error")
	// 	}
	// 	customers = append(customers, c)
	// }
	return customers, nil
}

func (d *CustomerRepositoryDb) ById(id string) (*Customer, *errs.AppError) {
	CustomerSql := "select customer_ID , name , city , pincode , date_of_birth , status from customers where customer_ID=?"
	//row := (*d).client.QueryRow(CustomerSql, id)
	var c Customer
	err := d.client.Get(&c, CustomerSql, id)
	//err2 := row.Scan(&c.Id, &c.Name, &c.City, &c.ZipCode, &c.DateOfBirth, &c.Status)

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

func NewCustomerRepositoryDB(dbclient *sqlx.DB) *CustomerRepositoryDb {
	// dbUser := os.Getenv("DB_USER")
	// dbPassword := os.Getenv("DB_PASSWORD")
	// //dbAddr := os.Getenv()
	// //dbPort := os.Getenv()
	// dbName := os.Getenv("DB_NAME")
	// dataSource := fmt.Sprintf("%s:%s@/%s", dbUser, dbPassword, dbName)
	// client, err := sqlx.Open("mysql", dataSource)
	// if err != nil {
	// 	panic(err)
	// }
	// // "root:12345678@/rest_tutorial"
	// // See "Important settings" section.
	// client.SetConnMaxLifetime(time.Minute * 3)
	// client.SetMaxOpenConns(10)
	// client.SetMaxIdleConns(10)
	return &CustomerRepositoryDb{dbclient}
}

//mysql://root:12345678@127.0.0.1?statusColor=005392&env=local&name=Localhost&tLSMode=0&usePrivateKey=false&safeModeLevel=0&advancedSafeModeLevel=0&driverVersion=0
