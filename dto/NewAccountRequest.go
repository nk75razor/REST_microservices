package dto

import (
	"strings"

	"github.com/nk75razor/REST_microservices/errs"
)

type NewAccountRequest struct {
	CustomerId  string  `json:"Customer_id"`
	AccountType string  `json:"account_type"`
	Amount      float64 `json:"amount"`
}

func (r NewAccountRequest) Validate() *errs.AppError {
	if r.Amount < 5000 {
		return errs.NewValidationError("Top open a new account you need to deposit atleast 5000")
	}
	if strings.ToLower(r.AccountType) != "saving" && strings.ToLower(r.AccountType) != "checking" {
		return errs.NewValidationError("Account type shuld be checking or saving")
	}
	return nil
}
