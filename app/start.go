package app

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"github.com/nk75razor/REST_microservices/domain"
	"github.com/nk75razor/REST_microservices/service"
)

func SanityCheck() {
	if os.Getenv("SERVER_ADDRESS") == "" ||
		os.Getenv("SERVER_PORT") == "" {
		log.Fatal("Enivronment variable not defined....")
	}

}
func Start() {
	//mux := http.NewServeMux()
	SanityCheck()
	router := mux.NewRouter()

	// wiring
	//ch := CustomerHandlers{service.NewCustomerService(domain.NewCustomerRepositoryStub())}
	dbClient := getDbClient()
	customerRepositoryDb := domain.NewCustomerRepositoryDB(dbClient)
	accountRepositoryDb := domain.NewAccountRepositoryDB(dbClient)

	ch := CustomerHandlers{service.NewCustomerService(customerRepositoryDb)}
	ah := AccountHandler{service.NewAccountService(accountRepositoryDb)}
	router.HandleFunc("/greet", greet) // this line define routes
	router.HandleFunc("/custom", ch.getAllCustomers)
	router.HandleFunc("/customer/{customer_id:[0-9]+}", getcustomer)
	router.HandleFunc("/createcustomer", CreateCustomer) //.Methods(http.MethodPost)
	router.HandleFunc("/customers/{customer_id:[0-9]+}", ch.GetCustomer)
	router.HandleFunc("/customers/{customer_id:[0-9]+}/account", ah.NewAccount).Methods(http.MethodPost)
	address := os.Getenv("SERVER_ADDRESS")
	port := os.Getenv("SERVER_PORT")

	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%s", address, port), router)) //starting the server
	//log.Fatal(http.ListenAndServe("root:12345678@/rest_tutorial"))
}
func getDbClient() *sqlx.DB {
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	//dbAddr := os.Getenv()
	//dbPort := os.Getenv()
	dbName := os.Getenv("DB_NAME")
	dataSource := fmt.Sprintf("%s:%s@/%s", dbUser, dbPassword, dbName)
	client, err := sqlx.Open("mysql", dataSource)
	if err != nil {
		panic(err)
	}
	// "root:12345678@/rest_tutorial"
	// See "Important settings" section.
	client.SetConnMaxLifetime(time.Minute * 3)
	client.SetMaxOpenConns(10)
	client.SetMaxIdleConns(10)
	return client
}
