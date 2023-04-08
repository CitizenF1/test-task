package main

import (
	"log"
	"os"
	"rest-api/internal/app"
)

func main() {
	app := app.NewApp()

	if err := app.Run(os.Getenv("APP_PORT")); err != nil {
		log.Fatalf("%s", err.Error())
	}
}
