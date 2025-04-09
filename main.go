package main

import (
	"github.com/angular403/gin-bioskop/config"
	"github.com/angular403/gin-bioskop/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.New()
	config.Connect()
	routes.UserRoute(router)
	router.Run(":8080")
}