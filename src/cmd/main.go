package main

import (
	"currency/app"
	"currency/cmd/common"
	"currency/internal/config"
	"currency/internal/logs"
)

func main() {
	common.Execute()
	logs.Init(&config.Currency.Logs)
	app.Serve()
}
