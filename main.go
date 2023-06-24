package main

import (
	"github.com/nk75razor/REST_microservices/app"
	"github.com/nk75razor/REST_microservices/logger"
)

func main() {
	//fmt.Println("Starting our application")
	logger.Info("starting the application")
	app.Start()

}
