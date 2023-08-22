package model

type ShopIdUrl struct {
	Url string `json:"url" gorm:"primaryKey"`
}

type DataShopidCrawled struct {
	Data struct {
		Total         int `json:"total"`
		OfficialShops []struct {
			Userid   int    `json:"userid"`
			Username string `json:"username"`
			Shopid   int    `json:"shopid"`
		} `json:"official_shops"`
	} `json:"data"`
}
