package utils

import (
	"encoding/json"
	"fmt"
	"gitlab.com/merakilab9/j4/pkg/model"
	"io/ioutil"
	"net/http"
)

func CategoryCrawler() (model.CrawlCate, error) {
	url := "https://shopee.vn/api/v4/pages/get_category_tree"
	method := "GET"

	CrawlClient := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		return model.CrawlCate{}, err
	}
	req.Header.Add("Cookie", "REC_T_ID=f69b8aaf-2566-11ee-9f53-2cea7fa89cee; SPC_F=dvS2x6Y73B3S82eEx63FiiB9xSEFaUB4; SPC_R_T_ID=yTWd8qwe1QKBWAKZuh1DfbvCfa4Dv/0HxLfC1KCSE6npwXfos0xwZIGR4juVWUhAIe68On4uAzgtEP6qLkJ1zdM7pbxxsYGFYXoZ2vcUpAfTaKobuUn8DXiar9P/wl6sHrIR2z2Ogoit/39nDHPdCUfbNtOxv9Gt0B+lUSDl4dc=; SPC_R_T_IV=azg2YUFLZEFhTDRLTllJcg==; SPC_SI=xji2ZAAAAABsaWRKM1BNNxIkCwAAAAAAbGNybUN4bk0=; SPC_T_ID=yTWd8qwe1QKBWAKZuh1DfbvCfa4Dv/0HxLfC1KCSE6npwXfos0xwZIGR4juVWUhAIe68On4uAzgtEP6qLkJ1zdM7pbxxsYGFYXoZ2vcUpAfTaKobuUn8DXiar9P/wl6sHrIR2z2Ogoit/39nDHPdCUfbNtOxv9Gt0B+lUSDl4dc=; SPC_T_IV=azg2YUFLZEFhTDRLTllJcg==")

	res, err := CrawlClient.Do(req)
	if err != nil {
		fmt.Println(err)
		return model.CrawlCate{}, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return model.CrawlCate{}, err
	}

	var result model.CrawlCate
	err = json.Unmarshal(body, &result)
	if err != nil {
		fmt.Println(err)
		return model.CrawlCate{}, err
	}

	return result, nil
}

//func ShopIdCrawl(shopidUrls []model.ShopIdUrl) ([]model.DataShopidCrawled, error) {
//
//	var results []model.DataShopidCrawled
//
//	client := &http.Client{}
//	for _, shop := range shopidUrls {
//		fmt.Println(shop.Url)
//		url := shop.Url
//		method := "GET"
//
//		req, err := http.NewRequest(method, url, nil)
//		if err != nil {
//			fmt.Println(err)
//			return []model.DataShopidCrawled{}, nil
//		}
//		req.Header.Add("Cookie", "REC_T_ID=f69b8aaf-2566-11ee-9f53-2cea7fa89cee; SPC_EC=ZG1oSXVTbXlGM01IRnZidrWPxxkQJAIzjGkR6uX/+PDCB5+xDiWeHXT+EcZvxBrkpSN9Fn7JIeKQyFRh/eL70SvyBY2lMAH2o5Xc6le5K/9ATzGmQcXR1Oi7XQ3eM1zlknZ61xsVCN3D9eTqWlUKDBV7eryo4MFKLBdBDMBI/ZM=; SPC_F=dvS2x6Y73B3S82eEx63FiiB9xSEFaUB4; SPC_R_T_ID=sIQZSunWuJoRYFmG4Tgsoh/97t8fyGbPaqWBcEO807OJMM2TNrNsbBFFNqx3eP8i9FT95zJhrZp8EI+rKhGouEHiLzF7gNozww7uRhYQSnWrGlP8r9MhiCi5ou81qDjBFLyKNyLWyOPAHE10/w/lhm5m+rckQWprtMmNyZwfAio=; SPC_R_T_IV=YzkxMDZCQXd2R0x4UEUwZw==; SPC_SI=ZYPkZAAAAAAzaGpzaHdKRiF9BwAAAAAANWhDeUxsS2I=; SPC_ST=.WWJ1WWNPamlPd2g2ekJFUcdPxH8dAr7Hy8XX7zCQkDr1Mggq8k+uyItDB8tOfAPYV5q+r7CKcMurhtk1x0t7veEXmqWXkOFV/HgPG8gdiwt9W8ZtF9xmZ+98zzt2yVyBEKFZqPX1jQDT7XxbOxQCQXvSlM31XqpjVfLw3Z5FGF9KCn/FC4+w554FLg28k5Oywx0EdngqQyi4mP/VAkL5cQ==; SPC_T_ID=sIQZSunWuJoRYFmG4Tgsoh/97t8fyGbPaqWBcEO807OJMM2TNrNsbBFFNqx3eP8i9FT95zJhrZp8EI+rKhGouEHiLzF7gNozww7uRhYQSnWrGlP8r9MhiCi5ou81qDjBFLyKNyLWyOPAHE10/w/lhm5m+rckQWprtMmNyZwfAio=; SPC_T_IV=YzkxMDZCQXd2R0x4UEUwZw==; SPC_U=153129342")
//
//		res, err := client.Do(req)
//		if err != nil {
//			fmt.Println(err)
//			return []model.DataShopidCrawled{}, nil
//		}
//		defer res.Body.Close()
//
//		body, err := ioutil.ReadAll(res.Body)
//		if err != nil {
//			fmt.Println(err)
//			return []model.DataShopidCrawled{}, nil
//		}
//
//		var result model.DataShopidCrawled
//		err = json.Unmarshal(body, &result)
//		if err != nil {
//			fmt.Println(err)
//			return []model.DataShopidCrawled{}, err
//		}
//
//		results = append(results, result)
//	}
//	return results, nil
//}
