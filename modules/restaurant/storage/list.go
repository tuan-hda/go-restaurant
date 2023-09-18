package restaurantstorage

import (
	"context"
	"g07/common"
	restaurantmodel "g07/modules/restaurant/model"
)

func (s *sqlStore) ListDataWithCondition(
	ctx context.Context,
	filter *restaurantmodel.Filter,
	paging *common.Paging,
	moreKeys ...string,
) ([]restaurantmodel.Restaurant, error) {
	db := s.db

	var result []restaurantmodel.Restaurant

	if filter.UserId > 0 {
		db = db.Where("owner_id = ?", filter.UserId)
	}

	if err := db.Table(restaurantmodel.Restaurant{}.TableName()).Count(&paging.Total).Error; err != nil {
		return nil, err
	}

	db = db.Where("status not in (0)")

	offset := (paging.Page - 1) * paging.Limit

	if err := db.Limit(paging.Limit).Offset(offset).Order("id desc").Find(&result).Error; err != nil {
		return nil, err
	}

	return result, nil
}
