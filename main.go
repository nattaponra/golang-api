package main

import (
	"os"

	"github.com/nattaponra/golangapi/app"
)

func main() {

	app := app.App{}

	app.Initialize(
		os.Getenv("APP_DB_HOST"),
		os.Getenv("APP_DB_PORT"),
		os.Getenv("APP_DB_USERNAME"),
		os.Getenv("APP_DB_PASSWORD"),
		os.Getenv("APP_DB_NAME"))

	app.Run(":8001")
}
