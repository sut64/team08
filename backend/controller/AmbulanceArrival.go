package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sut64/team08/entity"
)

func CreateAmbulanceArrival(c *gin.Context) {
	var ambulancearrival entity.AmbulanceArrival
	if err := c.ShouldBindJSON(&ambulancearrival); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := entity.DB().Create(&ambulancearrival).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": ambulancearrival})
}

func GetAmbulanceArrival(c *gin.Context) {
	var ambulancearrival entity.AmbulanceArrival
	id := c.Param("id")
	if err := entity.DB().Raw("SELECT * FROM ambulancearrivals WHERE id = ?", id).Scan(&ambulancearrival).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": ambulancearrival})
}

func ListAmbulanceArrivals(c *gin.Context) {
	var ambulancearrivals []entity.AmbulanceArrival
	if err := entity.DB().Raw("SELECT * FROM ambulancearrivals").Scan(&ambulancearrivals).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": ambulancearrivals})
}

func DeleteAmbulanceArrival(c *gin.Context) {
	id := c.Param("id")
	if tx := entity.DB().Exec("DELETE FROM ambulancearrivals WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ambulancearrival not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": id})
}

func UpdateAmbulanceArrival(c *gin.Context) {
	var ambulancearrival entity.AmbulanceArrival
	if err := c.ShouldBindJSON(&ambulancearrival); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if tx := entity.DB().Where("id = ?", ambulancearrival.ID).First(&ambulancearrival); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ambulancearrival not found"})
		return
	}
	if err := entity.DB().Save(&ambulancearrival).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": ambulancearrival})
}
