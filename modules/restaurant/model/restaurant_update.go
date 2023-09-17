package restaurantmodel

import "g07/common"

type RestaurantUpdate struct {
	common.SQLModel
	Name    *string `json:"name" gorm:"column:name;"`
	Address *string `json:"address" gorm:"column:addr;"`
}

func (RestaurantUpdate) TableName() string { return Restaurant{}.TableName() }
