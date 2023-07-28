package service

import (
	"fmt"
	"gitlab.com/merakilab9/j4/pkg/model"
	"gitlab.com/merakilab9/j4/pkg/repo/pg"
)

type CateService struct {
	repo pg.PGInterface
}

func NewCateService(repo pg.PGInterface) CateInterface {
	return &CateService{repo: repo}
}

type CateInterface interface {
	GetUrlCate() ([]model.Cate, error)
}

func (s *CateService) GetUrlCate() ([]model.Cate, error) {
	result, err := s.repo.GetUrlCate()
	if err != nil {
		return nil, fmt.Errorf("internal server")
	}
	return result, nil
}
