package main

import (
	"net/http"

	"employee-api/config"
	"employee-api/database"
	"employee-api/logger"
	"employee-api/routes"
)

func main() {

	cfg := config.Load()

	err := database.Connect(
		cfg.DBHost,
		cfg.DBPort,
		cfg.DBUser,
		cfg.DBPassword,
		cfg.DBName,
	)
	if err != nil {
		logger.Error("Failed to connect to database: " + err.Error())
		return
	}

	defer database.DB.Close()

	err = database.CreateTable()
	if err != nil {
		logger.Error("Failed to create employees table: " + err.Error())
		return
	}

	router := routes.SetupRouter()

	logger.Info("Starting server on port " + cfg.Port)

	err = http.ListenAndServe(":"+cfg.Port, router)
	if err != nil {
		logger.Error("Server failed to start: " + err.Error())
	}
}
