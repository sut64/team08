package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sut64/team08/entity"
)

func CreateAssessment(c *gin.Context) {
	var assessment entity.Assessment
	if err := c.ShouldBindJSON(&assessment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := entity.DB().Create(&assessment).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": assessment})
}

func GetAssessment(c *gin.Context) {
	var assessment entity.Assessment
	id := c.Param("id")
	if err := entity.DB().Raw("SELECT * FROM assessments WHERE id = ?", id).Scan(&assessment).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": assessment})
}

func ListAssessment(c *gin.Context) {
	var assessment []entity.Assessment
	if err := entity.DB().Raw("SELECT * FROM assessments").Scan(&assessment).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": assessment})
}

func DeleteAssessment(c *gin.Context) {
	id := c.Param("id")
	if tx := entity.DB().Exec("DELETE FROM assessments WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "assessment not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": id})
}

func UpdateAssessment(c *gin.Context) {
	var assessment entity.Assessment
	if err := c.ShouldBindJSON(&assessment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if tx := entity.DB().Where("id = ?", assessment.ID).First(&assessment); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "assessment not found"})
		return
	}
	if err := entity.DB().Save(&assessment).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": assessment})
}
