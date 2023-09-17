package main

import (
	"g07/common"
	restaurantgin "g07/modules/restaurant/transport/gin"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"net/http"
	"os"
	"strconv"
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
	//
	//log.Println(db, err)

	//newRestaurant := Restaurant{Name: "200lab", Address: "Somewhere"}
	//if err := db.Create(&newRestaurant).Error; err != nil {
	//	log.Println(err)
	//}
	//
	//log.Println("Inserted ID", newRestaurant.Id)

	//var oldRes Restaurant
	//if err := db.Select("id, name").Where(map[string]interface{}{"id": 4}).First(&oldRes).Error; err != nil {
	//	log.Println(err)
	//}
	//
	//log.Println(oldRes)
	//
	//var listRes []Restaurant
	//
	//if err := db.Limit(10).Find(&listRes).Error; err != nil {
	//	log.Println(err)
	//}
	//
	//log.Println(listRes)
	//
	//dataUpdate := Restaurant{
	//	Name: "Tani Coffee",
	//}
	//
	//if err := db.Where("id = ?", 3).Updates(&dataUpdate).Error; err != nil {
	//	log.Println(err)
	//}

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
			restaurants.POST("", restaurantgin.CreateRestaurant(db))

			restaurants.GET("/:id", restaurantgin.GetRestaurant(db))

			restaurants.GET("", func(c *gin.Context) {
				var data []Restaurant

				type Paging struct {
					Page  int `json:"page" form:"page"`
					Limit int `json:"limit" form:"limit"`
				}

				var paging Paging

				if err := c.ShouldBind(&paging); err != nil {
					c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
					return
				}

				if paging.Page < 1 {
					paging.Page = 1
				}

				if paging.Limit <= 0 {
					paging.Limit = 10
				}

				offset := (paging.Page - 1) * paging.Limit

				if err := db.Offset(offset).Limit(paging.Limit).Order("id desc").Find(&data).Error; err != nil {
					c.JSON(http.StatusInternalServerError, gin.H{
						"error": err.Error(),
					})
					return
				}

				c.JSON(http.StatusOK, gin.H{
					"data": data,
				})
			})

			restaurants.PUT("/:id")

			restaurants.DELETE("/:id", func(c *gin.Context) {
				id, err := strconv.Atoi(c.Param("id"))

				if err != nil {
					c.JSON(http.StatusBadRequest, gin.H{
						"error": err.Error(),
					})
					return
				}

				if err := db.Table(Restaurant{}.TableName()).Where("id = ?", id).Delete(nil).Error; err != nil {
					c.JSON(http.StatusInternalServerError, gin.H{
						"error": err.Error(),
					})
					return
				}

				c.JSON(http.StatusOK, gin.H{
					"data": true,
				})
			})
		}
	}

	r.Run()
}
