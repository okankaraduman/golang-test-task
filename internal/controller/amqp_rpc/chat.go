package amqprpc

import (
	"fmt"

	"github.com/okankaraduman/golang-test-task/internal/entity"
	"github.com/okankaraduman/golang-test-task/pkg/logger"
	"github.com/streadway/amqp"

	"github.com/okankaraduman/golang-test-task/pkg/rabbitmq/rmq_rpc/server"
)

type chatRoutes struct {
	l logger.Interface
}

func newChatRoutes(routes map[string]server.CallHandler, l logger.Interface) {
	r := &chatRoutes{l}
	{
		routes["subscribeMessages"] = r.subscribeAndPublish()
	}
}

type Messages struct {
	Messages []entity.Pair `json:"history"`
}

/*
func (r *chatRoutes) getMessages() server.CallHandler {
	return func(d *amqp.Delivery) (interface{}, error) {

		if err != nil {
			return nil, fmt.Errorf("amqp_rpc - chatRoutes - getMessages - r.translationUseCase.History: %w", err)
		}

		response := Messages{translations}

		return response, nil
	}
}

*/

func (r *chatRoutes) subscribeAndPublish() server.CallHandler {
	return func(d *amqp.Channel) (interface{}, error) {

		channel,err  := d.Consume("message")

		if err != nil {
			return nil, fmt.Errorf("amqp_rpc - chatRoutes - subscribeAndPublish - r.translationUseCase.History: %w", err)
		}

		channel.

		return nil, nil
	}

}
