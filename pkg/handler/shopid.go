package handler

import (
	"encoding/json"
	"github.com/hibiken/asynq"
	"gitlab.com/merakilab9/j4/pkg/service"
	"gitlab.com/merakilab9/j4/pkg/utils"
	"gitlab.com/merakilab9/meracore/ginext"
	"log"
	"net/http"
)

type ShopIdHandlers struct {
	service service.ShopIdInterface
	client  *asynq.Client
}

func NewShopIdHandlers(service service.ShopIdInterface) *ShopIdHandlers {
	client := asynq.NewClient(asynq.RedisClientOpt{Addr: utils.RedisAddr})
	return &ShopIdHandlers{service: service, client: client}
}

func (h *ShopIdHandlers) GetUrlShopId(c *ginext.Request) (*ginext.Response, error) {

	rs, err := h.service.GetUrlShopId()
	if err != nil {
		return nil, err
	}

	return ginext.NewResponseData(http.StatusOK, rs), nil
}

func (h *ShopIdHandlers) GetUrlShopDetails(c *ginext.Request) (*ginext.Response, error) {

	rs, err := h.service.GetUrlShopDetails()
	if err != nil {
		return nil, err
	}

	return ginext.NewResponseData(http.StatusOK, rs), nil
}

func (h *ShopIdHandlers) SendShopDetailsAPIToQueue(c *ginext.Request) (*ginext.Response, error) {

	rs, err := h.service.GetUrlShopDetails()
	if err != nil {
		return nil, err
	}

	for _, v := range rs {
		data4tune := v.Url
		payload4tune, _ := json.Marshal(data4tune)
		task := asynq.NewTask(utils.APIShopDetailsDelivery, payload4tune)
		log.Printf(" Create tasks: %v", err)

		info, err := h.client.Enqueue(task)
		if err != nil {
			log.Fatalf("could not enqueue tasks: %v", err)
		}
		log.Printf("enqueued tasks: id=%s queue=%s", info.ID, info.Queue)
	}

	return ginext.NewResponseData(http.StatusOK, rs), nil
}
