package service

import (
	"encoding/json"
	"fmt"
	"gitlab.com/merakilab9/j4/pkg/model"
	"gitlab.com/merakilab9/j4/pkg/repo/pg"
	"gitlab.com/merakilab9/j4/pkg/utils"
	"io/ioutil"
	"net/http"
)

type ShopIdService struct {
	repo pg.PGInterface
}

func NewShopIdService(repo pg.PGInterface) ShopIdInterface {
	return &ShopIdService{repo: repo}
}

type ShopIdInterface interface {
	GetUrlShopId() ([]model.ShopIdUrl, error)
	GetUrlShopDetails() ([]model.ShopDetail, error)
}

func (s *ShopIdService) GetUrlShopId() ([]model.ShopIdUrl, error) {
	//b1
	result, err := utils.CategoryCrawler()
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	//b2
	s.repo.SaveCate(result)
	//b3
	urls, err := s.repo.GetCateid()
	//b4
	s.repo.CreateShopidURL(urls)
	//b5
	cats, err := s.repo.GetUrlShopid()
	if err != nil {
		return nil, fmt.Errorf("internal server")
	}
	return cats, nil
}

func (s *ShopIdService) GetUrlShopDetails() ([]model.ShopDetail, error) {
	//b1
	result, err := s.ShopIdCrawl()
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	//b2
	for _, data := range result {
		err := s.repo.SaveShopID(data)
		if err != nil {
			fmt.Println(err)
		}
	}
	//b3
	urls, err := s.repo.GetShopID()
	//b4
	s.repo.CreateShopDetailsURL(urls)
	//b5
	shops, err := s.repo.GetUrlShopDetails()
	if err != nil {
		return nil, fmt.Errorf("internal server")
	}
	return shops, nil
}

func (s *ShopIdService) ShopIdCrawl() ([]model.DataShopidCrawled, error) {

	var results []model.DataShopidCrawled

	client := &http.Client{}
	cats, err := s.repo.GetUrlShopid()
	if err != nil {
		return nil, fmt.Errorf("internal server")
	}
	for _, shop := range cats {
		url := shop.Url
		method := "GET"

		req, err := http.NewRequest(method, url, nil)
		if err != nil {
			fmt.Println(err)
			return []model.DataShopidCrawled{}, nil
		}
		req.Header.Add("Cookie", "REC_T_ID=f69b8aaf-2566-11ee-9f53-2cea7fa89cee; SPC_EC=ZG1oSXVTbXlGM01IRnZidrWPxxkQJAIzjGkR6uX/+PDCB5+xDiWeHXT+EcZvxBrkpSN9Fn7JIeKQyFRh/eL70SvyBY2lMAH2o5Xc6le5K/9ATzGmQcXR1Oi7XQ3eM1zlknZ61xsVCN3D9eTqWlUKDBV7eryo4MFKLBdBDMBI/ZM=; SPC_F=dvS2x6Y73B3S82eEx63FiiB9xSEFaUB4; SPC_R_T_ID=sIQZSunWuJoRYFmG4Tgsoh/97t8fyGbPaqWBcEO807OJMM2TNrNsbBFFNqx3eP8i9FT95zJhrZp8EI+rKhGouEHiLzF7gNozww7uRhYQSnWrGlP8r9MhiCi5ou81qDjBFLyKNyLWyOPAHE10/w/lhm5m+rckQWprtMmNyZwfAio=; SPC_R_T_IV=YzkxMDZCQXd2R0x4UEUwZw==; SPC_SI=ZYPkZAAAAAAzaGpzaHdKRiF9BwAAAAAANWhDeUxsS2I=; SPC_ST=.WWJ1WWNPamlPd2g2ekJFUcdPxH8dAr7Hy8XX7zCQkDr1Mggq8k+uyItDB8tOfAPYV5q+r7CKcMurhtk1x0t7veEXmqWXkOFV/HgPG8gdiwt9W8ZtF9xmZ+98zzt2yVyBEKFZqPX1jQDT7XxbOxQCQXvSlM31XqpjVfLw3Z5FGF9KCn/FC4+w554FLg28k5Oywx0EdngqQyi4mP/VAkL5cQ==; SPC_T_ID=sIQZSunWuJoRYFmG4Tgsoh/97t8fyGbPaqWBcEO807OJMM2TNrNsbBFFNqx3eP8i9FT95zJhrZp8EI+rKhGouEHiLzF7gNozww7uRhYQSnWrGlP8r9MhiCi5ou81qDjBFLyKNyLWyOPAHE10/w/lhm5m+rckQWprtMmNyZwfAio=; SPC_T_IV=YzkxMDZCQXd2R0x4UEUwZw==; SPC_U=153129342")

		res, err := client.Do(req)
		if err != nil {
			fmt.Println(err)
			return []model.DataShopidCrawled{}, nil
		}
		defer res.Body.Close()

		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			fmt.Println(err)
			return []model.DataShopidCrawled{}, nil
		}

		var result model.DataShopidCrawled
		err = json.Unmarshal(body, &result)
		if err != nil {
			fmt.Println(err)
			return []model.DataShopidCrawled{}, err
		}

		results = append(results, result)
	}
	return results, nil
}
