package pg

import (
	"gitlab.com/merakilab9/j4/pkg/model"
)

func (r *RepoPG) GetUrlCate() ([]model.Cate, error) {
	var cates []model.Cate

	if err := r.DB.Find(&cates).Error; err != nil {
		return nil, err
	}

	return cates, nil
}
