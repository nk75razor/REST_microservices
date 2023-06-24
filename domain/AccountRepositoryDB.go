package domain

import (
	"strconv"

	"github.com/jmoiron/sqlx"
	"github.com/nk75razor/REST_microservices/errs"
	"github.com/nk75razor/REST_microservices/logger"
)

type AccountRepositoryDB struct {
	client *sqlx.DB
}

func (d AccountRepositoryDB) Save(a Account) (*Account, *errs.AppError) {
	sqlInsert := "INSERT INTO accounts(customer_id, opening_date, account_type,amount,status) values(?,?,?,?,?)"
	result, err := d.client.Exec(sqlInsert, a.AccountId, a.OpeningDate, a.AccountType, a.Amount, a.Status)
	if err != nil {
		logger.Error("Error while creating a new account" + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected error from database")
	}
	id, err := result.LastInsertId()
	if err != nil {
		logger.Error("Error while last insert id for new account:" + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected error from database")
	}
	a.AccountId = strconv.FormatInt(id, 10)
	return &a, nil
}
func NewAccountRepositoryDB(dbClient *sqlx.DB) AccountRepositoryDB {
	return AccountRepositoryDB{dbClient}
}
