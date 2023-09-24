package upload

import (
	"g07/component/appctx"
	"github.com/gin-gonic/gin"
	"net/http"
)
gin

import (
	"g07/component/appctx"
	"github.com/gin-gonic/gin"
	"net/http"
)

func UploadImage(appCtx appctx.AppContext) func(*gin.Context) {
	return func(c *gin.Context) {
		fileHeader, err := c.FormFile("file")

		if err != nil {
			c.JSON(http.StatusBadGateway, gin.H{"error": err.Error()})
			return
		}

		//folder := c.DefaultPostForm("folder", "img")

		file, err := fileHeader.Open()

		if err != nil {
			c.JSON(http.StatusBadGateway, gin.H{"error": err.Error()})
			return
		}

		defer file.Close()

		dataBytes := make([]byte, fileHeader.Size)
		if _, err := file.Read(dataBytes); err != nil {
			c.JSON(http.StatusBadGateway, gin.H{"error": err.Error()})
			return
		}
	}
}
