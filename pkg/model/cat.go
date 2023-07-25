package model

type ShopeeCrawl struct {
	Data struct {
		CategoryList []struct {
			Catid       int    `json:"catid"`
			ParentCatid int    `json:"parent_catid"`
			Name        string `json:"name"`
			DisplayName string `json:"display_name"`
			Image       string `json:"image"`
			Level       int    `json:"level"`
			Children    []struct {
				Catid       int         `json:"catid"`
				ParentCatid int         `json:"parent_catid"`
				Name        string      `json:"name"`
				DisplayName string      `json:"display_name"`
				Image       string      `json:"image"`
				Level       int         `json:"level"`
				Children    interface{} `json:"children"`
			}
		} `json:"category_list"`
	} `json:"data"`
}
