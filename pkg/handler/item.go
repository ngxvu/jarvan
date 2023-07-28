package handler

import (
	"gitlab.com/merakilab9/j4/pkg/service"
	"gitlab.com/merakilab9/meracore/ginext"
	"net/http"
)

type ItemHandlers struct {
	service service.ItemInterface
}

func NewItemHandlers(service service.ItemInterface) *ItemHandlers {
	return &ItemHandlers{service: service}
}

func (h *ItemHandlers) GetUrlItem(c *ginext.Request) (*ginext.Response, error) {

	rs, err := h.service.GetUrlItem()
	if err != nil {
		return nil, err
	}

	return ginext.NewResponseData(http.StatusOK, rs), nil
}
