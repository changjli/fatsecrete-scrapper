package yandex_image_controller

import (
	"net/http"
	"scrapper/atom/yandex_image"

	"github.com/gin-gonic/gin"
)

func InsertAllV2(c *gin.Context) {
	// validasi binding

	err := yandex_image.InsertAllServiceV2()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status":  http.StatusBadRequest,
			"data":    nil,
			"message": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status":  http.StatusOK,
			"data":    nil,
			"message": "sucess create user",
		})
	}
}
