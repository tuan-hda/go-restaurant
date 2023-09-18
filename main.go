package main

import (
	"g07/common"
	"g07/component/appctx"
	restaurantgin "g07/modules/restaurant/transport/gin"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"net/http"
	"os"
)

type Restaurant struct {
	common.SQLModel
	Name    string `json:"name" gorm:"column:name;"`
	Address string `json:"address" gorm:"column:addr;"`
}

func (Restaurant) TableName() string { return "restaurants" }

type RestaurantCreate struct {
	common.SQLModel
	Name    *string `json:"name" gorm:"column:name;"`
	Address *string `json:"address" gorm:"column:addr;"`
}

func (RestaurantCreate) TableName() string { return Restaurant{}.TableName() }

type RestaurantUpdate struct {
	common.SQLModel
	Name    *string `json:"name" gorm:"column:name;"`
	Address *string `json:"address" gorm:"column:addr;"`
}

func (RestaurantUpdate) TableName() string { return Restaurant{}.TableName() }

func main() {
	dsn := os.Getenv("DB_CONN_STR")
	db, _ := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	db = db.Debug()

	appCtx := appctx.NewAppContext(db)

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	v1 := r.Group("api/v1")
	{
		restaurants := v1.Group("/restaurants")
		{
			restaurants.POST("", restaurantgin.CreateRestaurant(appCtx))

			restaurants.GET("/:id", restaurantgin.GetRestaurant(appCtx))

			restaurants.GET("", restaurantgin.ListRestaurant(appCtx))

			restaurants.PUT("/:id", restaurantgin.UpdateRestaurant(appCtx))

			restaurants.DELETE("/:id", restaurantgin.DeleteRestaurant(appCtx))
		}
	}

	r.Run()
}
