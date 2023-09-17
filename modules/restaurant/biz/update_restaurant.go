package restaurantbiz

import (
	"context"
	"errors"
	"g07/common"
	restaurantmodel "g07/modules/restaurant/model"
)

type UpdateRestaurantBiz interface {
	FindDataWithCondition(
		ctx context.Context,
		cond map[string]interface{},
		moreKeys ...string,
	) (*restaurantmodel.Restaurant, error)

	Update(
		ctx context.Context,
		cond map[string]interface{},
		updateData *restaurantmodel.RestaurantUpdate,
		moreKeys ...string,
	) error
}

type updateRestaurantBiz struct {
	store UpdateRestaurantBiz
}

func NewUpdateRestaurantBiz(store UpdateRestaurantBiz) *updateRestaurantBiz {
	return &updateRestaurantBiz{store: store}
}

func (biz *updateRestaurantBiz) UpdateRestaurant(ctx context.Context, id int, data *restaurantmodel.RestaurantUpdate) error {
	oldData, err := biz.store.FindDataWithCondition(ctx, map[string]interface{}{"id": id})

	if err != nil {
		if errors.Is(err, common.ErrDataNotFound) {
			return errors.New("data not found")
		}

		return err
	}

	if oldData.Status == 0 {
		return errors.New("data has been deleted")
	}

	err = biz.store.Update(ctx, map[string]interface{}{"id": id}, data)

	if err != nil {
		return err
	}

	return nil
}
