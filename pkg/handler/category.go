package handler

import (
	"encoding/json"
	"fmt"
	"github.com/hibiken/asynq"
	"gitlab.com/merakilab9/j4/pkg/model"
	"gitlab.com/merakilab9/j4/pkg/service"
	"gitlab.com/merakilab9/j4/pkg/utils"
	"gitlab.com/merakilab9/meracore/ginext"
	"log"
	"net/http"
)

type CateHandlers struct {
	service service.CateInterface
	client  *asynq.Client
}

func NewCateHandlers(service service.CateInterface) *CateHandlers {
	client := asynq.NewClient(asynq.RedisClientOpt{Addr: utils.RedisAddr})
	return &CateHandlers{service: service, client: client}
}

func (h *CateHandlers) GetUrlCate(c *ginext.Request) (*ginext.Response, error) {

	rs, err := h.service.GetUrlCate()
	if err != nil {
		return nil, err
	}

	return ginext.NewResponseData(http.StatusOK, rs), nil

}

func (h *CateHandlers) SendAPIToQueue(c *ginext.Request) (*ginext.Response, error) {

	//=================== MssBroker ===================//

	rs, err := h.service.GetUrlCate()
	if err != nil {
		return nil, err
	}

	data4tune := model.RequestData{
		rs,
	}

	payload4tune, err := json.Marshal(data4tune)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	task := asynq.NewTask(utils.APICateDelivery, payload4tune)
	log.Printf(" Create tasks: %v", err)

	info, err := h.client.Enqueue(task)
	if err != nil {
		log.Fatalf("could not enqueue tasks: %v", err)
	}
	log.Printf("enqueued tasks: id=%s queue=%s", info.ID, info.Queue)

	return ginext.NewResponseData(http.StatusOK, rs), nil
}
