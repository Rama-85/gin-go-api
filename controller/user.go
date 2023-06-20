package controller

import (
	"gin-go-api/config"
	"gin-go-api/models"

	"github.com/gin-gonic/gin"
)

func GetDays(c *gin.Context) {
	days := []models.Calendar{}
	config.DB.Find(&days)
	c.JSON(200, &days)
}

func CreateMonth(c *gin.Context) {
	var month models.Calendar
	config.DB.Create(&month)
	c.BindJSON(&month)
}

func DeleteMonth(c *gin.Context) {
	var month models.Calendar
	config.DB.Where("id=?", c.Param("id")).Delete(&month)
	c.JSON(200, &month)
}

func UpdateMonth(c *gin.Context) {
	var umonth models.Calendar
	config.DB.Where("id=?", c.Param("id")).First(&umonth)
	c.BindJSON(&umonth)
	config.DB.Save(&umonth)
	c.JSON(200, &umonth)
}
