package model

type DataShopidCrawled struct {
	Data struct {
		OfficialShops []OfficialShop `json:"official_shops"`
	} `json:"data"`
}

type OfficialShop struct {
	Userid   int    `json:"userid" gorm:"foreignKey:shopid"`
	Username string `json:"username"`
	Shopid   int    `json:"shopid" gorm:"primaryKey"`
}

type ShopIdUrl struct {
	Url string `json:"url" gorm:"primaryKey"`
}

type ShopDetail struct {
	Url string `json:"url" gorm:"primaryKey"`
}
