package service

import (
	"fmt"
	"gitlab.com/merakilab9/j4/pkg/model"
	"gitlab.com/merakilab9/j4/pkg/repo/pg"
)

type ItemService struct {
	repo pg.PGInterface
}

func NewItemService(repo pg.PGInterface) ItemInterface {
	return &ItemService{repo: repo}
}

type ItemInterface interface {
	GetUrlItem() ([]model.Item, error)
}

func (s *ItemService) GetUrlItem() ([]model.Item, error) {
	result, err := s.repo.GetUrlItem()
	if err != nil {
		return nil, fmt.Errorf("internal server")
	}
	return result, nil
}
