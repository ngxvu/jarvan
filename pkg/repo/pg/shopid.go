package pg

import (
	"gitlab.com/merakilab9/j4/pkg/model"
	"strconv"
)

func (r *RepoPG) GetUrlShopid() ([]model.ShopIdUrl, error) {
	r.CreateShopidURL()
	var shopid []model.ShopIdUrl

	if err := r.DB.Find(&shopid).Error; err != nil {
		return nil, err
	}

	return shopid, nil
}

func (r *RepoPG) CreateShopidURL() ([]string, error) {
	var cates []model.CateCrawl
	var urls []string

	if err := r.DB.Find(&cates).Error; err != nil {
		return nil, err
	}

	// Lặp qua danh sách Cate lấy các catid và tạo URL
	for _, cate := range cates {
		urlsFromCate := GetURLFromCate(cate)
		urls = append(urls, urlsFromCate...)
	}

	// Lưu các URL vào bảng Shopid
	for _, url := range urls {
		shopid := model.ShopIdUrl{Url: url}
		if err := r.DB.Create(&shopid).Error; err != nil {
			return nil, err
		}
	}

	return urls, nil
}

func GetURLFromCate(cate model.CateCrawl) []string {
	var urls []string

	// Lấy catid từ mô hình Cate và tạo URL cho từng catid
	url := "https://shopee.vn/api/v4/official_shop/get_shops?category_id=" + strconv.Itoa(cate.Catid) + "&limit=25&offset=0"
	url1 := "https://shopee.vn/api/v4/official_shop/get_shops?category_id=" + strconv.Itoa(cate.Catid) + "&limit=25&offset=25"
	url2 := "https://shopee.vn/api/v4/official_shop/get_shops?category_id=" + strconv.Itoa(cate.Catid) + "&limit=25&offset=50"
	url3 := "https://shopee.vn/api/v4/official_shop/get_shops?category_id=" + strconv.Itoa(cate.Catid) + "&limit=25&offset=75"
	url4 := "https://shopee.vn/api/v4/official_shop/get_shops?category_id=" + strconv.Itoa(cate.Catid) + "&limit=25&offset=100"
	url5 := "https://shopee.vn/api/v4/official_shop/get_shops?category_id=" + strconv.Itoa(cate.Catid) + "&limit=25&offset=125"
	url6 := "https://shopee.vn/api/v4/official_shop/get_shops?category_id=" + strconv.Itoa(cate.Catid) + "&limit=25&offset=150"
	url7 := "https://shopee.vn/api/v4/official_shop/get_shops?category_id=" + strconv.Itoa(cate.Catid) + "&limit=25&offset=175"
	url8 := "https://shopee.vn/api/v4/official_shop/get_shops?category_id=" + strconv.Itoa(cate.Catid) + "&limit=25&offset=200"
	url9 := "https://shopee.vn/api/v4/official_shop/get_shops?category_id=" + strconv.Itoa(cate.Catid) + "&limit=25&offset=225"
	url10 := "https://shopee.vn/api/v4/official_shop/get_shops?category_id=" + strconv.Itoa(cate.Catid) + "&limit=25&offset=250"

	urls = append(urls, url, url1, url2, url3, url4, url5, url6, url7, url8, url9, url10)

	// Lặp qua các con của Cate
	for _, child := range cate.Children {
		urlsFromChild := GetURLFromCate(child)
		urls = append(urls, urlsFromChild...)
	}

	return urls
}
