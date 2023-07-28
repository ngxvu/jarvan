package handler

import (
	"gitlab.com/merakilab9/j4/pkg/service"
	"gitlab.com/merakilab9/meracore/ginext"
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

	return ginext.NewResponseData(http.StatusOK, rs), nil

	//url := "localhost:9002/api/v1/cate/parse-to-json"
	//url.Do()

	fortuneAPIURL := "localhost:9002/api/v1/cate/parse-to-json"
	http.Get(fortuneAPIURL)
	return ginext.NewResponseData(http.StatusOK, rs), nil
}
