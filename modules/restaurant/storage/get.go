package restaurantstorage

import (
	"context"
	"errors"
	"g07/common"
	restaurantmodel "g07/modules/restaurant/model"
	"gorm.io/gorm"
)

func (s *sqlStore) FindDataWithCondition(
	ctx context.Context,
	cond map[string]interface{},
	moreKeys ...string,
) (*restaurantmodel.Restaurant, error) {
	db := s.db

	var data restaurantmodel.Restaurant

	if err := db.Where(cond).First(&data).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, common.ErrDataNotFound
		}

		return nil, err
	}

	return &data, nil
}
