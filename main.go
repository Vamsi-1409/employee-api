package main

import (
	"log"
	"net/http"

	"employee-api/routes"
)

func main() {

	router := routes.SetupRouter()

	log.Println("Server started on :8080")

	log.Fatal(http.ListenAndServe(":8080", router))
}
