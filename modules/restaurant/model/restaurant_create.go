package restaurantmodel

import (
	"g07/common"
	"strings"
)

type RestaurantCreate struct {
	common.SQLModel
	Name    string `json:"name" gorm:"column:name;"`
	Address string `json:"address" gorm:"column:addr;"`
}

func (RestaurantCreate) TableName() string { return Restaurant{}.TableName() }

func (r *RestaurantCreate) Validate() error {
	r.Name = strings.TrimSpace(r.Name)

	if r.Name == "" {
		return ErrNameCannotBeBlank
	}

	r.Address = strings.TrimSpace(r.Address)

	if r.Address == "" {
		return ErrAddressCannotBeBlank
	}

	return nil
}
