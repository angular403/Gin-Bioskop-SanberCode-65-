package routes

import (
	"github.com/angular403/gin-bioskop/controller"
	"github.com/gin-gonic/gin"
)

func UserRoute(router *gin.Engine){
	router.GET("/", controller.GetUsers)
	router.POST("/", controller.CreateUsers)
	router.PUT("/", controller.UpdateUsers)
	router.DELETE("/", controller.DeleteUsers)
}