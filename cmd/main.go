package main

import (
	"TaskManager/internal/config"
	"TaskManager/internal/handlers"
	"TaskManager/internal/repository"
	"TaskManager/internal/service"
	"fmt"
	"github.com/joho/godotenv"
	"go.uber.org/zap"
	"os"
	"os/signal"
	"syscall"
)

func main() {

	cfg := config.MustLoad()

	logger, _ := zap.NewDevelopment()
	defer logger.Sync()

	if err := godotenv.Load(".env"); err != nil {
		logger.Fatal("Error loading .env file", zap.Error(err))
	}

	db, err := repository.NewMongo(os.Getenv("MONGODB_URI"))
	if err != nil {
		logger.Fatal("Failed to connect to MongoDB", zap.Error(err))
	}

	repo := repository.NewRepository(db, cfg.MongoDB)
	services := service.NewService(repo)
	handlers := handlers.NewHandler(services, logger)

	e := handlers.InitRoutes(logger)

	go func() {
		logger.Info(fmt.Sprintf("Listening on port %s", cfg.Port))
		if err := e.Start(cfg.Port); err != nil {
			logger.Fatal("Failed to start server", zap.Error(err))
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	logger.Info("Shutting down server gracefully")

}
