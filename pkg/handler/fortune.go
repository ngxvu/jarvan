package handler

import (
	"encoding/json"
	"gitlab.com/merakilab9/meracore/ginext"
	"gitlab.com/merakilab9/meracrawler/fortune/pkg/service"
	"net/http"
)

type FortuneHandler struct {
	service service.FortuneInterface
}

func NewFortuneHandlers(service service.FortuneInterface) *FortuneHandler {
	return &FortuneHandler{service: service}
}

func (h *FortuneHandler) ProcessURLs(r *ginext.Request) (*ginext.Response, error) {
	client := &http.Client{}
	var urls []string
	err := json.NewDecoder(r.GinCtx.Request.Body).Decode(&urls)
	if err != nil {
		return nil, ginext.NewError(http.StatusBadRequest, err.Error())
	}

	if len(urls) == 0 {
		return nil, ginext.NewError(http.StatusBadRequest, err.Error())
	}

	rs, err := h.service.ProcessURLs(r.GinCtx, client, urls)
	if err != nil {
		return nil, err
	}
	return ginext.NewResponseData(http.StatusOK, rs), nil
}

func (h *FortuneHandler) HelloWorld(r *ginext.Request) (*ginext.Response, error) {
	return ginext.NewResponseData(http.StatusOK, "hello"), nil
}
