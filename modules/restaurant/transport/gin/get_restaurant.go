package restaurantgin

import (
	"context"
	"g07/common"
	"g07/component/appctx"
	restaurantbiz "g07/modules/restaurant/biz"
	restaurantmodel "g07/modules/restaurant/model"
	restaurantstorage "g07/modules/restaurant/storage"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

type fakeGetDataStore struct {
}

func (fakeGetDataStore) FindDataWithCondition(ctx context.Context, cond map[string]interface{}, moreKeys ...string) (*restaurantmodel.Restaurant, error) {
	return &restaurantmodel.Restaurant{
		common.SQLModel{
			Id:        1,
			Status:    1,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		"Fake restaurant",
		"Fake address",
	}, nil
}

func GetRestaurant(appCtx appctx.AppContext) func(c *gin.Context) {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		store := restaurantstorage.NewSQLStore(appCtx.GetMainDBConnection())
		biz := restaurantbiz.NewGetRestaurantBiz(store)

		//fakeStore := fakeGetDataStore{}
		//biz := restaurantbiz.NewGetRestaurantBiz(fakeStore)

		data, err := biz.GetRestaurant(c.Request.Context(), id)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"data": data,
		})
	}

}
