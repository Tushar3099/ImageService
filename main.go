package main

import (
	"log"
	"net/http"
	"os"

	"github.com/RetailPulse/apis"
	"github.com/RetailPulse/modules/controller"
	"github.com/joho/godotenv"
	"github.com/rs/cors"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}
	port, ok := os.LookupEnv("PORT")
	if !ok {
		port = "3000"
	}
	controller := controller.New()
	//registering routes
	api := apis.New(controller)

	//connecting to the port
	log.Println("server is listening on port ", port)
	log.Fatal(http.ListenAndServe(":"+port, cors.Default().Handler(api.Router())).Error())
}
