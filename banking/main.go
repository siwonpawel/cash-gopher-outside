package main

import (
	"github.com/siwonpawel/cash-gopher-outside/banking/app"
	"github.com/siwonpawel/cash-gopher-outside/banking/logger"
)

func main() {
	logger.Info("Starting the application...")
	app.Start()
}
