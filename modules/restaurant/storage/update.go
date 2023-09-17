package restaurantstorage

import (
	"context"
	restaurantmodel "g07/modules/restaurant/model"
)

func (s *sqlStore) Update(
	ctx context.Context,
	cond map[string]interface{},
	updateData *restaurantmodel.RestaurantUpdate,
	moreKeys ...string,
) error {
	db := s.db

	if err := db.Where(cond).Updates(updateData).Error; err != nil {
		return err
	}

	return nil
}
