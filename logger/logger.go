package logger

import (
	"log"
	"os"
)

var Logger = log.New(os.Stdout, "", log.Ldate|log.Ltime)

func Info(message string) {
	Logger.Println("[INFO] " + message)
}

func Error(message string) {
	Logger.Println("[ERROR] " + message)
}
