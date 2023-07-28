package pg

import (
	"gitlab.com/merakilab9/j4/pkg/model"
)

func (r *RepoPG) GetUrlShopid() ([]model.Shopid, error) {
	var shopid []model.Shopid

	if err := r.DB.Find(&shopid).Error; err != nil {
		return nil, err
	}

	return shopid, nil
}
