package main

import (
	"orders.go/m/internal/setup"
)

func main() {
	app := setup.New()

	err := app.Run()
	if err != nil {
		panic(err)
	}
}
