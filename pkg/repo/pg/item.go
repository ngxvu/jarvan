package pg

import (
	"gitlab.com/merakilab9/j4/pkg/model"
)

func (r *RepoPG) GetUrlItem() ([]model.Item, error) {
	var items []model.Item

	if err := r.DB.Find(&items).Error; err != nil {
		return nil, err
	}

	return items, nil
}
