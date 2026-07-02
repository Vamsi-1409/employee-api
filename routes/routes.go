package routes

import (
	"employee-api/handlers"

	"github.com/gorilla/mux"
)

func SetupRouter() *mux.Router {

	router := mux.NewRouter()

	router.HandleFunc("/health", handlers.HealthCheck).Methods("GET")
        router.HandleFunc("/employees", handlers.GetEmployees).Methods("GET")
        router.HandleFunc("/employees", handlers.CreateEmployee).Methods("POST")
        router.HandleFunc("/employees/{id}", handlers.GetEmployeeByID).Methods("GET")
        router.HandleFunc("/employees/{id}", handlers.UpdateEmployee).Methods("PUT")
        router.HandleFunc("/employees/{id}", handlers.DeleteEmployee).Methods("DELETE")
	return router
}
