package restaurantgin

import (
	"g07/component/appctx"
	restaurantbiz "g07/modules/restaurant/biz"
	restaurantmodel "g07/modules/restaurant/model"
	restaurantstorage "g07/modules/restaurant/storage"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateRestaurant(appCtx appctx.AppContext) func(c *gin.Context) {
	return func(c *gin.Context) {
		var newData restaurantmodel.RestaurantCreate

		if err := c.ShouldBind(&newData); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Dependencies install
		store := restaurantstorage.NewSQLStore(appCtx.GetMainDBConnection())
		biz := restaurantbiz.NewCreateRestaurantBiz(store)

		if err := biz.CreateNewRestaurant(c.Request.Context(), &newData); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"data": newData.Id,
		})
	}

}
