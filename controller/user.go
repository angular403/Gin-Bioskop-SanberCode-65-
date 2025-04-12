package controller

import (
	"github.com/angular403/gin-bioskop/config"
	"github.com/angular403/gin-bioskop/models"
	"github.com/gin-gonic/gin"
)

func GetUsers(c *gin.Context) {
	users := []models.User{}
	config.DB.Find(&users)
	c.JSON(200, &users)
}

func CreateUsers(c *gin.Context) {

	var users models.User
	c.Bind(users)
	config.DB.Find(&users)
	c.JSON(200, &users)
}

func UpdateUsers(c *gin.Context) {
	var users models.User
	config.DB.Where("id = ?", c.Param("id")).Delete(&users)
	c.Bind(users)
	config.DB.Find(&users)
	c.JSON(200, &users)
}

func DeleteUsers(c *gin.Context) {
	var user models.User
	config.DB.Where("id = ?", c.Param("id")).Delete(&user)
	c.JSON(200, &user)
}
