package main

import (
	"test-api-task/internal/app"
	"test-api-task/internal/config"
)

func main() {
	run()
}

func run() {
	cfg := config.MustLoad()

	app, err := app.NewApp(*cfg)
	if err != nil {
		panic(err)
	}

	app.StartHTTPServer()
}
