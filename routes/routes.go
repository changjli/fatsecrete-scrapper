package routes

import (
	fatsecret_controller "scrapper/atom/fatsecrete/controllers"
	yandex_image_controller "scrapper/atom/yandex_image/controllers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupRoutes() *gin.Engine {
	route := gin.Default()

	// bisa langsung atau disini
	route.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	api := route.Group("/api")

	api.POST("/fatsecrete", fatsecret_controller.InsertByName)
	api.GET("/yandex-image", yandex_image_controller.InsertAllV2)

	return route
}
