package service

import (
	"io/ioutil"
	"net/http"
)

type FortuneService struct {
	Client *http.Client
}

func NewFortuneService(client *http.Client) FortuneInterface {
	return &FortuneService{Client: client}
}

type FortuneInterface interface {
	ProcessURLsCate(client *http.Client, urls []string) (string, error)
}

func (s *FortuneService) ProcessURLsCate(client *http.Client, urls []string) (string, error) {
	method := "GET"
	var bodies [][]byte
	for _, url := range urls {
		req, err := http.NewRequest(method, url, nil)
		if err != nil {
			return "", err
		}
		res, err := client.Do(req)
		if err != nil {
			return "", err
		}
		defer res.Body.Close()

		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			return "", err
		}
		bodies = append(bodies, body)
	}

	var result string
	for _, body := range bodies {
		result += string(body)
	}
	return result, nil
}
