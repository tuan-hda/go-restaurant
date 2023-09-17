package restaurantbiz

import (
	"context"
	"errors"
	restaurantmodel "g07/modules/restaurant/model"
	"strings"
)

type CreateRestaurantBiz interface {
	Create(ctx context.Context, data *restaurantmodel.RestaurantCreate) error
}

type createNewRestaurantBiz struct {
	store CreateRestaurantBiz
}

func NewCreateRestaurantBiz(store CreateRestaurantBiz) *createNewRestaurantBiz {
	return &createNewRestaurantBiz{store: store}
}

func (biz *createNewRestaurantBiz) CreateNewRestaurant(ctx context.Context, data *restaurantmodel.RestaurantCreate) error {
	data.Name = strings.TrimSpace(data.Name)

	if data.Name == "" {
		return errors.New("restaurant name cannot be blank")
	}

	data.Address = strings.TrimSpace(data.Address)

	if data.Address == "" {
		return errors.New("restaurant address cannot be blank")
	}

	if err := biz.store.Create(ctx, data); err != nil {
		return err
	}

	return nil
}
