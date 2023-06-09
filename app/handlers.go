package app

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/nk75razor/REST_microservices/service"
)

// type Customer struct {
// 	Name    string `json:"full_name"  xml:"half_name" `
// 	City    string
// 	ZipCode string
// }

type CustomerHandlers struct {
	service service.CustomerService
}

func greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello World")
}

func (ch *CustomerHandlers) getAllCustomers(w http.ResponseWriter, r *http.Request) {
	// customers := []Customer{
	// 	{"ashish ", " newdelhi", "93939"},
	// 	{"nayan", "jaipur", "302039"},
	// }
	customers, _ := ch.service.GetAllCustomer()
	if r.Header.Get("content-Type") == "application/xml" {
		w.Header().Add("Content-Type", "appliction/xml")
		xml.NewEncoder(w).Encode(customers)
	} else {
		w.Header().Add("Content-Type", "appliction/json")
		json.NewEncoder(w).Encode(customers)
	}
	//json.NewEncoder(w).Encode(customers)

}

func getcustomer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Fprint(w, vars["customer_id"])
}

func CreateCustomer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "post request received")
}

func (ch *CustomerHandlers) GetCustomer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["customer_id"]
	customer, err := ch.service.GetCustomer(id)

	if err != nil {
		//WriteResponse(w, err.Code, err.AsMessage())
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(err.Code)
		//fmt.Fprint(w, err.Message)
		json.NewEncoder(w).Encode(err.AsMessage())
	} else {
		//WriteResponse(w, http.StatusOK, customer)
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(customer)

	}
}

func WriteResponse(w http.ResponseWriter, code int, data interface{}) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		panic(err)
	}

}
