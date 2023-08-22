package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	fmt.Println("hello")

	// call all cate id

	url := "https://shopee.vn/api/v4/pages/get_category_tree"
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		fmt.Println(err)
		return
	}

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	listCateChild := []int{}
	catRes := CateRes{}
	err = json.Unmarshal(body, &catRes)

	for _, category := range catRes.Data.CategoryList {
		listCateChild = append(listCateChild, category.Catid)
	}

	fmt.Println("Total cate child:", len(listCateChild))

	totalShop := 0
	listOfficalShopID := []int{}

	for _, catid := range listCateChild {
		url := fmt.Sprintf("https://shopee.vn/api/v4/official_shop/get_shops?category_id=%v", catid)
		method := "GET"
		client := &http.Client{}
		req, err := http.NewRequest(method, url, nil)
		if err != nil {
			fmt.Println(err)
			return
		}
		res, err := client.Do(req)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer res.Body.Close()
		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			fmt.Println(err)
			return
		}
		catRes := ShopRes{}
		err = json.Unmarshal(body, &catRes)
		fmt.Println(catid, "Total shop: ", catRes.Data.Total)
		totalShop += catRes.Data.Total
		for _, v := range catRes.Data.Shops {
			listOfficalShopID = append(listOfficalShopID, v.Shopid)
		}
	}

	fmt.Println("Total shop: ", totalShop)
	fmt.Println("Total offical shop: ", len(listOfficalShopID))

	for _, shopid := range listOfficalShopID {
		url := fmt.Sprintf("https://shopee.vn/api/v4/product/get_shop_info?shopid=%v", shopid)
		method := "GET"
		client := &http.Client{}
		req, err := http.NewRequest(method, url, nil)
		if err != nil {
			fmt.Println(err)
			return
		}
		res, err := client.Do(req)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer res.Body.Close()
		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(string(body))
	}
}

type ShopRes struct {
	Data struct {
		Total int    `json:"total"`
		Shops []Shop `json:"official_shops"`
	}
}
type Shop struct {
	Shopid int `json:"shopid"`
}
type CateRes struct {
	Data struct {
		CategoryList []Category `json:"category_list"`
	}
}

type Category struct {
	Catid    int        `json:"catid"`
	Children []Category `json:"children"`
}
