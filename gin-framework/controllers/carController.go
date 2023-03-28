package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Car struct {
	CarId string `json:"car_id"`
	Brand string `json:"brand"`
	Model string `json:"model"`
	Price int    `json:"price"`
}

var Cars = []Car{}

func CreateCar(ctx *gin.Context) {
	var newCar Car

	if err :=  ctx.ShouldBindJSON(&newCar); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	newCar.CarId = fmt.Sprintf("c%d", len(Cars)+1)
	Cars = append(Cars, newCar)

	ctx.JSON(http.StatusCreated, gin.H{
		"car": newCar,
		"message": "success create data",
	})
}

func GetCars(ctx *gin.Context) {
	var newCars Car

	// if err :=  ctx.ShouldBindJSON(&newCars); err != nil {
	// 	ctx.AbortWithError(http.StatusBadRequest, err)
	// 	return
	// }

	Cars = append(Cars, newCars)

	ctx.JSON(http.StatusOK, gin.H{
		"data": newCars,
		"message": "success get data",
	})
}
