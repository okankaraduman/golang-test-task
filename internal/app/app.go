// Package app configures and runs application.
package app

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"
	"github.com/okankaraduman/golang-test-task/config"
	amqprpc "github.com/okankaraduman/golang-test-task/internal/controller/amqp_rpc"
	"github.com/okankaraduman/golang-test-task/pkg/httpserver"
	"github.com/okankaraduman/golang-test-task/pkg/logger"
	"github.com/okankaraduman/golang-test-task/pkg/rabbitmq/rmq_rpc/server"
	"github.com/okankaraduman/golang-test-task/pkg/redis"
)

// Run creates objects via constructors.
func Run(cfg *config.Config) {

	l := logger.New(cfg.Log.Level)

	//RFC Server

	// RabbitMQ RPC Server
	rmqRouter := amqprpc.NewRouter(l)

	rmqServer, err := server.New(cfg.RMQ.URL, cfg.RMQ.ServerExchange, rmqRouter, l)
	if err != nil {
		l.Fatal(fmt.Errorf("app - Run - rmqServer - server.New: %w", err))
	}

	// Redis
	re, err := redis.New(cfg.Re.Host+":"+cfg.Re.Port, redis.MaxPoolSize(cfg.Re.PoolMax))
	if err != nil {
		l.Fatal(fmt.Errorf("app - Run - redis.New: %w", err))
	}
	defer re.Close()

	r := gin.Default()
	v1(r, l)
	httpServer := httpserver.New(r, httpserver.Port(cfg.HTTP.Port))
	//Router

	// Waiting signal
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case s := <-interrupt:
		l.Info("app - Run - signal: " + s.String())
	case err = <-httpServer.Notify():
		l.Error(fmt.Errorf("app - Run - httpServer.Notify: %w", err))
	}

	// Shutdown
	err = httpServer.Shutdown()
	if err != nil {
		l.Error(fmt.Errorf("app - Run - httpServer.Shutdown: %w", err))
	}
	err = rmqServer.Shutdown()
	if err != nil {
		l.Error(fmt.Errorf("app - Run - rmqServer.Shutdown: %w", err))
	}
}
