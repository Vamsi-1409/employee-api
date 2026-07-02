package main

import (
	"net/http"

	"employee-api/routes"
        "employee-api/config"
        "employee-api/logger"
)

func main() {

	router := routes.SetupRouter()
        cfg := config.Load()

        logger.Info("Server started on part " + cfg.Port)

        err := http.ListenAndServe(":"+cfg.Port, router)
        if err != nil {
	   logger.Error(err.Error())
        }
}
