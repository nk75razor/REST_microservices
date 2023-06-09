package domain

type CustomerRepositoryStub struct {
	customers []Customer
}

func (s CustomerRepositoryStub) FindAll() ([]Customer, error) {
	return s.customers, nil
}

func NewCustomerRepositoryStub() CustomerRepositoryStub {
	customers := []Customer{
		{"1001", "Nayan", "Jaipur", "101901", "20-04-2002", "1"},
		{"1002", "saumya", "NewDelhi", "101901", "01-02-2006", "1"},
	}
	return CustomerRepositoryStub{customers}
}
