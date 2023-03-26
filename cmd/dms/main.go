package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/lefes/discord-message-scheduler/internal/config"
	"github.com/lefes/discord-message-scheduler/internal/repository"
	"github.com/lefes/discord-message-scheduler/internal/service"
	"github.com/lefes/discord-message-scheduler/internal/transport/discord"
	"github.com/lefes/discord-message-scheduler/internal/transport/http"
	"github.com/lefes/discord-message-scheduler/pkg/database/badgerdb"
	"github.com/lefes/discord-message-scheduler/pkg/logger"
)

func main() {
	// Config init
	cfg, err := config.Init()
	if err != nil {
		logger.Error(err)

		return
	}

	// BadgerDB init
	db, err := badgerdb.NewClient(cfg.DB)
	if err != nil {
		logger.Error(err)

		return
	}

	defer db.Close()

	// // Repositories and services
	repos := repository.NewRepositories(db)
	services := service.NewServices(service.Deps{Repos: repos})

	// HTTP server
	// TODO: rework handlers registration
	httpServer := http.NewServer(cfg, nil)
	httpServer.RegisterHandlers()

	go func() {
		if err := httpServer.Start(); err != nil {
			logger.Error(err)

			return
		}
	}()

	logger.Info("Server is running...")

	// Discord client
	commands := discord.NewCommands(services)
	handlers := discord.NewHandlers(commands)
	discordClient, err := discord.NewClient(&cfg.Discord)
	if err != nil {
		logger.Error(err)

		return
	}

	if err := discordClient.Start(*handlers, *commands); err != nil {
		logger.Error(err)

		return
	}

	logger.Info("Discord client is running...")

	// Graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	<-quit

	logger.Info("Shutting down...")

	if err := discordClient.Shutdown(); err != nil {
		logger.Error(err)

		return
	}
	logger.Info("Discord client gracefully stopped")

	const timeout = 5 * time.Second

	ctx, shutdown := context.WithTimeout(context.Background(), timeout)
	defer shutdown()

	if err := httpServer.Shutdown(ctx); err != nil {
		logger.Error(err)

		return
	}
	logger.Info("Server gracefully stopped")

	// Close DB
	err = db.Close()
	if err != nil {
		logger.Error(err)

		return
	}

	logger.Info("DB gracefully stopped")

}
