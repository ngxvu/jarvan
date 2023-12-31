package pg

import (
	"gitlab.com/merakilab9/j4/pkg/model"
	"log"
)

func (r *RepoPG) GetUrlCate() ([]model.CateUrl, error) {

	newCateUrl := model.CateUrl{Url: "https://shopee.vn/api/v4/pages/get_category_tree"}
	r.DB.Create(newCateUrl)

	var cateUrls []model.CateUrl

	if err := r.DB.Find(&cateUrls).Error; err != nil {
		return nil, err
	}

	return cateUrls, nil
}

func (r *RepoPG) SaveCate(result model.CrawlCate) error {
	categories := result.Data.CategoryList

	flattenedCategories := []model.CateCrawl{}
	for _, category := range categories {
		flattenCategories(category, &flattenedCategories)
	}

	for _, category := range flattenedCategories {
		err := r.DB.Create(&category).Error
		if err != nil {
			log.Println(err)
		}

		if len(category.Children) > 0 {
			for _, child := range category.Children {
				child.ParentCatid = category.Catid
				err := r.DB.Create(&child).Error
				if err != nil {
					log.Println(err)
				}
			}
		}
	}

	return nil
}

func flattenCategories(category model.CateCrawl, flattenedCategories *[]model.CateCrawl) {
	*flattenedCategories = append(*flattenedCategories, category)

	for _, child := range category.Children {
		child.ParentCatid = category.Catid
		flattenCategories(child, flattenedCategories)
	}
}
