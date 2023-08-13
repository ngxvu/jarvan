package handler

import (
	"encoding/json"
	"fmt"
	"github.com/hibiken/asynq"
	"gitlab.com/merakilab9/j4/pkg/model"
	"gitlab.com/merakilab9/j4/pkg/service"
	"gitlab.com/merakilab9/j4/pkg/utils"
)

type CateHandlersQueue struct {
	service service.CateInterface
}

func NewSendAPIToQueueHandlers(service service.CateInterface) *CateHandlersQueue {
	return &CateHandlersQueue{service: service}
}

func (h *CateHandlersQueue) NewAPIDeliveryTask() (*asynq.Task, error) {

	rs, err := h.service.GetUrlCate()
	if err != nil {
		return nil, err
	}
	data4tune := model.RequestData{
		rs,
	}
	payload4tune, err := json.Marshal(&data4tune)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return asynq.NewTask(utils.APICateDelivery, payload4tune), nil
}
