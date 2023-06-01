package apps

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"net/http"
)

type Customer struct {
	Name    string `json:"full_name"  xml:"half_name" `
	City    string
	ZipCode string
}

func greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello World")
}

func getAllCustomers(w http.ResponseWriter, r *http.Request) {
	customers := []Customer{
		{"ashish ", " newdelhi", "93939"},
		{"nayan", "jaipur", "302039"},
	}
	if r.Header.Get("content-Type") == "application/xml" {
		w.Header().Add("Content-Type", "appliction/xml")
		xml.NewEncoder(w).Encode(customers)
	} else {
		w.Header().Add("Content-Type", "appliction/json")
		json.NewEncoder(w).Encode(customers)
	}
	//json.NewEncoder(w).Encode(customers)

}
