package service

import (
	"context"
	"gitlab.com/merakilab9/meracore/ginext"
	"gitlab.com/merakilab9/meracore/logger"
	"gitlab.com/merakilab9/meracrawler/fortune/pkg/utils"
	"net/http"
)

type FortuneService struct {
	client ClientServiceInterface
}

func NewFortuneService(client ClientServiceInterface) FortuneInterface {
	return &FortuneService{client: client}
}

type FortuneInterface interface {
	ProcessURLs(ctx context.Context, client *http.Client, urls []string) (string, error)
}

func (s *FortuneService) ProcessURLs(ctx context.Context, client *http.Client, urls []string) (string, error) {
	log := logger.WithCtx(ctx, "FortuneService.ProcessURLs")
	res, err := s.client.ProcessURLs(client, urls)
	if err != nil {
		log.WithError(err).Error("Error when ProcessURLs")
		return res, ginext.NewError(http.StatusInternalServerError, utils.MessageError()[http.StatusInternalServerError])
	}
	return res, nil
}
