package restaurantmodel

import (
	"errors"
	"g07/common"
)

type Restaurant struct {
	common.SQLModel
	Name    string `json:"name" gorm:"column:name;"`
	Address string `json:"address" gorm:"column:addr;"`
}

func (Restaurant) TableName() string { return "restaurants" }

var (
	ErrNameCannotBeBlank    = errors.New("name cannot be blank")
	ErrAddressCannotBeBlank = errors.New("address cannot be blank")
)
