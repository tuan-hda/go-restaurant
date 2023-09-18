package restaurantmodel

import (
	"g07/common"
	"strings"
)

type RestaurantUpdate struct {
	common.SQLModel
	Name    *string `json:"name" gorm:"column:name;"`
	Address *string `json:"address" gorm:"column:addr;"`
	Status  *int    `json:"-" gorm:"column:status;"`
}

func (RestaurantUpdate) TableName() string { return Restaurant{}.TableName() }

func (r *RestaurantUpdate) Validate() error {
	if strPtr := r.Name; strPtr != nil {
		str := strings.TrimSpace(*r.Name)

		if str == "" {
			return ErrNameCannotBeBlank
		}

		r.Name = &str
	}

	if strPtr := r.Address; strPtr != nil {
		str := strings.TrimSpace(*r.Address)

		if str == "" {
			return ErrAddressCannotBeBlank
		}

		r.Address = &str
	}

	return nil
}
