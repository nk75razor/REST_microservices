package dto

type CustomerResponse struct {
	Id          string `json:"customer_ID"`
	Name        string `json:"name"`
	City        string `json:"city"`
	ZipCode     string `json:"pincode"`
	DateOfBirth string `json:"date_of_birth"`
	Status      string `json:"status"`
}
