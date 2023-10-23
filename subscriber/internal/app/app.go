package app

import (
	"fmt"
	"github.com/gost1k337/wb_demo/subscriber/config"
	"github.com/gost1k337/wb_demo/subscriber/internal/cache"
	"github.com/gost1k337/wb_demo/subscriber/internal/handlers"
	"github.com/gost1k337/wb_demo/subscriber/internal/repository"
	"github.com/gost1k337/wb_demo/subscriber/internal/service"
	"github.com/gost1k337/wb_demo/subscriber/internal/subscriber"
	"github.com/gost1k337/wb_demo/subscriber/pkg/httpserver"
	"github.com/gost1k337/wb_demo/subscriber/pkg/logging"
	"github.com/gost1k337/wb_demo/subscriber/pkg/postgres"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
)

func Run(cfg *config.Config) {
	logger := logging.NewLogger(cfg)

	db, err := postgres.New(cfg.Db.DSN)
	if err != nil {
		logger.Fatalf("db: %v", err)
	}

	if err = db.Ping(); err != nil {
		logger.Fatalf("db conn: %v", err)
	}
	logger.Info("Connected to postgres")

	c := cache.NewCache()

	logger.Info("Initializing repositories...")
	repos := repository.NewRepositories(db, logger)

	logger.Info("Initializing services...")
	services := service.NewServices(&service.Deps{
		Cache: c,
		Repos: repos,
	}, logger)

	sc, err := subscriber.NewStanSub(services, logger, cfg.Nats.ClusterID, cfg.Nats.ClientID, net.JoinHostPort(cfg.Nats.Host, cfg.Nats.Port))
	if err != nil {
		logger.Fatalf("stan conn: %v", err)
	}
	logger.Info("Connected to stan server")

	sub, err := sc.SubscribeToChannel(cfg.Nats.Channel)
	if err != nil {
		log.Fatalf("sub: %v", err)
	}
	defer sub.Unsubscribe()
	logger.Infof("Subscriber to %s channel", cfg.Nats.Channel)

	logger.Info("Initializing handlers...")
	h := handlers.New(services, logger)

	httpServer := httpserver.New(h.HTTP(), httpserver.Port(cfg.App.Port))

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case s := <-interrupt:
		logger.Info("app - Run - signal: " + s.String())
	case err := <-httpServer.Notify():
		logger.Error(fmt.Errorf("app - Run - httpServer.Notify: %w", err))
	}

	logger.Info("Shutting down...")

	err = httpServer.Shutdown()

	if err != nil {
		logger.Error(fmt.Errorf("app - Run - httpServer.Shutdown: %w", err))
	}
}
