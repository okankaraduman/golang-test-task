package amqprpc

import (
	"github.com/okankaraduman/golang-test-task/pkg/logger"
	"github.com/okankaraduman/golang-test-task/pkg/rabbitmq/rmq_rpc/server"
)

// NewRouter -.
func NewRouter(l logger.Interface) map[string]server.CallHandler {
	routes := make(map[string]server.CallHandler)
	{
		newChatRoutes(routes, l)
	}

	return routes
}
