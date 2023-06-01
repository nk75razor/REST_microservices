package apps

import "net/http"

func start() {
	http.HandleFunc("/greet", greet) // this line define routes
	http.HandleFunc("/custom", getAllCustomers)

	http.ListenAndServe("localhost:8000", nil) //starting the server

}
