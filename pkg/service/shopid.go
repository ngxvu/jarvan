package service

import (
	"fmt"
	"gitlab.com/merakilab9/j4/pkg/model"
	"gitlab.com/merakilab9/j4/pkg/repo/pg"
)

type ShopIdService struct {
	repo pg.PGInterface
}

func NewShopIdService(repo pg.PGInterface) ShopIdInterface {
	return &ShopIdService{repo: repo}
}

type ShopIdInterface interface {
	GetUrlShopId() ([]model.Shopid, error)
}

func (s *ShopIdService) GetUrlShopId() ([]model.Shopid, error) {
	result, err := s.repo.GetUrlShopid()
	if err != nil {
		return nil, fmt.Errorf("internal server")
	}
	return result, nil
}
