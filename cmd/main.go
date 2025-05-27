package main

import (
	"TaskManager/internal/config"
	"go.uber.org/zap"
)

func main() {

	cfg := config.MustLoad()

	logger, _ := zap.NewDevelopment()
	defer logger.Sync()

}
