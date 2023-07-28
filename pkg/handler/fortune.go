package handler

import (
	"gitlab.com/merakilab9/meracore/ginext"
	"gitlab.com/merakilab9/meracore/logger"
	"gitlab.com/merakilab9/meracrawler/fortune/pkg/service"
	"gitlab.com/merakilab9/meracrawler/fortune/pkg/utils"
	"net/http"
)

type FortuneHandler struct {
	service service.FortuneInterface
}

func NewFortuneHandlers(service service.FortuneInterface) *FortuneHandler {
	return &FortuneHandler{service: service}
}

func (h *FortuneHandler) ProcessURLsCate(r *ginext.Request) (*ginext.Response, error) {
	log := logger.WithCtx(r.GinCtx, "ProcessURLsCate")

	var urls []string
	if err := r.GinCtx.ShouldBindJSON(&urls); err != nil {
		log.Println("Can not Decode Urls.", err)
		return nil, ginext.NewError(http.StatusInternalServerError, utils.MessageError()[http.StatusBadRequest])
	}

	if len(urls) == 0 {
		log.Fatal("No URLs provided.")
		return nil, ginext.NewError(http.StatusNotFound, "")
	}
	client := &http.Client{}
	rs, err := h.service.ProcessURLsCate(client, urls)
	if err != nil {
		log.Fatal("Fail to Process Service URLsCate.")
		return nil, err
	}
	return ginext.NewResponseData(http.StatusOK, rs), nil
}
