package handler

import (
	"gitlab.com/merakilab9/j4/pkg/service"
	"gitlab.com/merakilab9/meracore/ginext"
	"net/http"
)

type ShopIdHandlers struct {
	service service.ShopIdInterface
}

func NewShopIdHandlers(service service.ShopIdInterface) *ShopIdHandlers {
	return &ShopIdHandlers{service: service}
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
