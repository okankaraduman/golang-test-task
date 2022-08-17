package v1

import (
	"net/http"

	"github.com/okankaraduman/golang-test-task/pkg/logger"
	"github.com/okankaraduman/golang-test-task/utils"
)

type chat struct {
	l    logger.Interface
	repo repository.KeyValueRepo
}

// @Summary     Chat
// @Description pushes received information to a RabbitMQ Queue
// @ID          pushMessage
// @Tags  	    pushMessage
// @Accept      json
// @Produce     plain
// @Success     200 {string} string "Successfully pushed the message"
// @Failure     500  {string} string  "Database Problems"
// @Router      /message [post]
func (k *chat) pushMessage(w http.ResponseWriter, r *http.Request) {
	//data := map[string]server.CallHandler

	resp := utils.Response{Resp: w}
	resp.Text(http.StatusNotFound, "Not implemented", "text/plain")

}

// @Summary     Chat
// @Description gets  information from a RabbitMQ Queue
// @ID          pushMessage
// @Tags  	    pushMessage
// @Accept      json
// @Produce     plain
// @Success     200 {string} string "Successfully pushed the message"
// @Failure     500  {string} string  "Database Problems"
// @Router      /message/list [get]
func (k *chat) getMessage(w http.ResponseWriter, r *http.Request) {

	resp := utils.Response{Resp: w}
	resp.Text(http.StatusNotFound, "Not implemented", "text/plain")

}
