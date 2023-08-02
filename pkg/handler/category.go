package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"gitlab.com/merakilab9/j4/pkg/service"
	"gitlab.com/merakilab9/meracore/ginext"
	"io/ioutil"
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

	//
	url := "localhost:9002/api/v1/cate/parse-to-json"
	method := "POST"

	jsonData, err := json.Marshal(rs)
	if err != nil {
		return nil, err
	}

	client := &http.Client{}
	req, err := http.NewRequest(method, url, bytes.NewBuffer(jsonData))

	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	fmt.Println(string(body))

	return ginext.NewResponseData(http.StatusOK, rs), nil

}
