package routes

import (
	"gin-framework/controllers"

	"github.com/gin-gonic/gin"
)

func StartServer() *gin.Engine {
	router := gin.Default()

	router.POST("/cars", controllers.CreateCar)

	router.GET("/carss", controllers.GetCars)


	return router
}