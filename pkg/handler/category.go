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
}

func NewCateHandlers(service service.CateInterface) *CateHandlers {
	return &CateHandlers{service: service}
}

func (h *CateHandlers) GetUrlCate(c *ginext.Request) (*ginext.Response, error) {

	rs, err := h.service.GetUrlCate()
	if err != nil {
		return nil, err
	}

	//data := model.RequestData{
	//	rs,
	//}

	//payload, err := json.Marshal(data)
	//if err != nil {
	//	fmt.Println(err)
	//	return nil, err
	//}

	//url := "http://localhost:7099/api/v1/cate/parse-to-json"
	//method := "POST"
	//
	//client := &http.Client{}
	//req, err := http.NewRequest(method, url, bytes.NewBuffer(payload))
	//if err != nil {
	//	fmt.Println(err)
	//	return nil, err
	//}
	//req.Header.Set("Content-Type", "application/json")
	//
	//res, err := client.Do(req)
	//if err != nil {
	//	fmt.Println(err)
	//	return nil, err
	//}
	//defer res.Body.Close()
	//
	//body, err := ioutil.ReadAll(res.Body)
	//if err != nil {
	//	fmt.Println(err)
	//	return nil, err
	//}
	//fmt.Println(string(body))

	return ginext.NewResponseData(http.StatusOK, rs), nil

}

func (h *CateHandlers) SendAPIToQueue(c *ginext.Request) (*ginext.Response, error) {

	client := asynq.NewClient(asynq.RedisClientOpt{Addr: utils.RedisAddr})
	defer client.Close()

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

	info, err := client.Enqueue(task)
	if err != nil {
		log.Fatalf("could not enqueue tasks: %v", err)
	}
	log.Printf("enqueued tasks: id=%s queue=%s", info.ID, info.Queue)

	return ginext.NewResponseData(http.StatusOK, rs), nil

}
