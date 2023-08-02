package pg

import (
	"gitlab.com/merakilab9/j4/pkg/model"
	"strconv"
)

func (r *RepoPG) GetUrlShopid() ([]model.Shopid, error) {
	r.CreateShopidURL()
	var shopid []model.Shopid

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

	// Lặp qua danh sách Cate để lấy các catid và tạo URL
	for _, cate := range cates {
		urlsFromCate := GetURLFromCate(cate)
		urls = append(urls, urlsFromCate...)
	}

	// Lưu các URL vào bảng Shopid
	for _, url := range urls {
		shopid := model.Shopid{Url: url}
		if err := r.DB.Create(&shopid).Error; err != nil {
			return nil, err
		}
	}

	return urls, nil
}

func GetURLFromCate(cate model.CateCrawl) []string {
	var urls []string

	// Lấy catid từ mô hình Cate và tạo URL cho từng catid
	url := "https://shopee.vn/api/v4/official_shop/get_shops?category_id=" + strconv.Itoa(cate.Catid)
	urls = append(urls, url)

	// Lặp qua danh sách các con của Cate và gọi đệ quy để lấy các catid từ chúng
	for _, child := range cate.Children {
		urlsFromChild := GetURLFromCate(child)
		urls = append(urls, urlsFromChild...)
	}

	return urls
}
