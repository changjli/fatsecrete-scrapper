package fatsecret_controller

import (
	"net/http"
	"scrapper/atom/fatsecrete"

	"github.com/gin-gonic/gin"
)

func InsertByName(c *gin.Context) {
	// validasi binding
	var req fatsecrete.InsertByNameReq
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"data":    nil,
			"message": err.Error(),
		})
		return
	}

	err = fatsecrete.InsertByNameService(req)
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
