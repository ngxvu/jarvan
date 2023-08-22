package service

import (
	"fmt"
	"gitlab.com/merakilab9/j4/pkg/model"
	"gitlab.com/merakilab9/j4/pkg/repo/pg"
	"gitlab.com/merakilab9/j4/pkg/utils"
)

type ShopIdService struct {
	repo pg.PGInterface
}

func NewShopIdService(repo pg.PGInterface) ShopIdInterface {
	return &ShopIdService{repo: repo}
}

type ShopIdInterface interface {
	GetUrlShopId() ([]model.ShopIdUrl, error)
}

func (s *ShopIdService) GetUrlShopId() ([]model.ShopIdUrl, error) {
	result, err := utils.CategoryCrawler()
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	s.repo.SaveCate(result)
	cats, err := s.repo.GetUrlShopid()
	if err != nil {
		return nil, fmt.Errorf("internal server")
	}
	return cats, nil
}
