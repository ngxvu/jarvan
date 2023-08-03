package model

type CrawlCate struct {
	Data struct {
		CategoryList []CateCrawl `json:"category_list"`
	} `json:"data"`
}

type CateCrawl struct {
	Catid       int         `json:"catid" gorm:"primaryKey"`
	ParentCatid int         `json:"parent_catid"`
	Level       int         `json:"level"`
	Children    []CateCrawl `json:"children" gorm:"foreignKey:catid"`
}

type CateUrl struct {
	Url string `json:"url" gorm:"primaryKey"`
}

type RequestData struct {
	Data []CateUrl `json:"data"`
}
