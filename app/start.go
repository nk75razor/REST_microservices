package app

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/nk75razor/REST_microservices/domain"
	"github.com/nk75razor/REST_microservices/service"
)

func Start() {
	//mux := http.NewServeMux()
	router := mux.NewRouter()

	// wiring
	//ch := CustomerHandlers{service.NewCustomerService(domain.NewCustomerRepositoryStub())}
	ch := CustomerHandlers{service.NewCustomerService(domain.NewCustomerRepositoryDB())}
	router.HandleFunc("/greet", greet) // this line define routes
	router.HandleFunc("/custom", ch.getAllCustomers)
	router.HandleFunc("/customer/{customer_id:[0-9]+}", getcustomer)
	router.HandleFunc("/createcustomer", CreateCustomer) //.Methods(http.MethodPost)
	router.HandleFunc("/customers/{customer_id:[0-9]+}", ch.GetCustomer)
	http.ListenAndServe("localhost:8000", router) //starting the server

}
