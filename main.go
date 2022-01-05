package main

import (
	"log"
	"net/http"
	"os"

	"github.com/RetailPulse/apis"
	"github.com/joho/godotenv"
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
	//registering routes
	router := apis.NewRouter()

	//connecting to the port
	log.Println("server is listening on port ", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
