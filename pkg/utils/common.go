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
