package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sut64/team08/entity"
)

// POST /AmbulanceArrival
func CreateAmbulanceArrival(c *gin.Context) {

	var ambulancearrival entity.AmbulanceArrival
	var employee entity.Employee
	var patient entity.Patient
	var ambulanceonduty entity.AmbulanceOnDuty

	// ผลลัพธ์ที่ได้จากขั้นตอนที่ 7 จะถูก bind เข้าตัวแปร ambulanceArrival
	if err := c.ShouldBindJSON(&ambulancearrival); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 8: ค้นหา AmbulanceOnDuty ด้วย id
	if tx := entity.DB().Where("id = ?", ambulancearrival.AmbulanceOnDutyID).First(&ambulanceonduty); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ambulanceOnDuty not found"})
		return
	}

	// 9: ค้นหา Patient ด้วย id
	if tx := entity.DB().Where("id = ?", ambulancearrival.PatientID).First(&patient); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "patient not found"})
		return
	}

	// 10: ค้นหา employe ด้วย id
	if tx := entity.DB().Where("id = ?", ambulancearrival.RecorderID).First(&employee); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "employee not found"})
		return
	}

	// 11: สร้าง AmbulanceArrival
	aa := entity.AmbulanceArrival{
		Distance:            ambulancearrival.Distance,
		Number_of_passenger: ambulancearrival.Number_of_passenger,
		DateTime:            ambulancearrival.DateTime,
		AmbulanceOnDuty:     ambulanceonduty, // โยงความสัมพันธ์กับ Entity AmbulanceOnDuty ในตาราง AmbulanceArrival
		Patient:             patient,         // โยงความสัมพันธ์กับ Entity Patient
		Recorder:            employee,        // โยงความสัมพันธ์กับ Entity Employee
	}

	// 12: บันทึก
	if err := entity.DB().Create(&aa).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": aa})
}

// GET /AmbulanceArrival/:id
func GetAmbulanceArrival(c *gin.Context) {
	var ambulancearrival entity.AmbulanceArrival
	id := c.Param("id")
	if err := entity.DB().Preload("AmbulanceOnDuty").Preload("Patient").Preload("Recorder").Raw("SELECT * FROM ambulance_arrivals WHERE id = ?", id).Find(&ambulancearrival).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": ambulancearrival})
}

// GET /AmbulanceArrival
func ListAmbulanceArrivals(c *gin.Context) {
	var ambulancearrivals []entity.AmbulanceArrival
	if err := entity.DB().Preload("AmbulanceOnDuty").Preload("Patient").Preload("Recorder").Raw("SELECT * FROM ambulance_arrivals").Find(&ambulancearrivals).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": ambulancearrivals})
}

// Delete /AmbulanceArrivals/:id
func DeleteAmbulanceArrival(c *gin.Context) {
	id := c.Param("id")
	if tx := entity.DB().Exec("DELETE FROM ambulance_arrivals WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Ambulance Arrival not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": id})
}

// PATCH /AmbulanceArrivals
func UpdateAmbulanceArrival(c *gin.Context) {
	var ambulancearrival entity.AmbulanceArrival
	if err := c.ShouldBindJSON(&ambulancearrival); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if tx := entity.DB().Where("id = ?", ambulancearrival.ID).First(&ambulancearrival); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Ambulance Arrival not found"})
		return
	}
	if err := entity.DB().Save(&ambulancearrival).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": ambulancearrival})
}
