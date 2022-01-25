package controller

import (
	"github.com/sut64/team08/entity"

	"github.com/gin-gonic/gin"

	"net/http"
)

func CreateAmbulance(c *gin.Context) {

	var ambulance entity.Ambulance
	var ambulancestypes entity.AmbulanceType
	var status entity.Status
	var employee entity.Employee

	// ผลลัพธ์ที่ได้จากขั้นตอนที่ 7 จะถูก bind เข้าตัวแปร Ambulance
	if err := c.ShouldBindJSON(&ambulance); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 8: ค้นหา AmbulanceType ด้วย id
	if tx := entity.DB().Where("id = ?", ambulance.AmbulanceTypeID).First(&ambulancestypes); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ambulancetype not found"})
		return
	}

	// 9: ค้นหา Status ด้วย i
	if tx := entity.DB().Where("id = ?", ambulance.StatusID).First(&status); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "room not found"})
		return
	}

	// 10: ค้นหา Employee ด้วย id
	if tx := entity.DB().Where("id = ?", ambulance.EmployeeID).First(&employee); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "employee not found"})
		return
	}
	// 11: สร้าง CheckIn
	ami := entity.Ambulance{
		AmbulanceType: ambulancestypes,    // โยงความสัมพันธ์กับ Entity AmbulanceType
		Status:        status,             // โยงความสัมพันธ์กับ Entity Status
		Employee:      employee,           // โยงความสัมพันธ์กับ Entity Employee
		DateTime:      ambulance.DateTime, // ตั้งค่าฟิลด์ Date_time
		CarNumber:     ambulance.CarNumber,
		Registration:  ambulance.Registration,
	}

	// 12: บันทึก
	if err := entity.DB().Create(&ami).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": ami})
}

// GET /user/:id

func GetAmbulance(c *gin.Context) {

	var ambulance entity.Ambulance
	id := c.Param("id")

	if err := entity.DB().Preload("ambulancetype").Preload("status").Preload("employee").Raw("SELECT * FROM ambulance WHERE id = ?", id).Find(&ambulance).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": ambulance})

}

// GET /users

func ListAmbulances(c *gin.Context) {

	var ambulances []entity.Ambulance
	if err := entity.DB().Preload("ambulancetype").Preload("status").Preload("employee").Raw("SELECT * FROM ambulance").Find(&ambulances).Error; err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": ambulances})

}

// DELETE /users/:id

func DeleteAmbulance(c *gin.Context) {

	id := c.Param("id")
	if tx := entity.DB().Exec("DELETE FROM ambulance WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ambulance not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": id})
}

// PATCH /users

func UpdateAmbulance(c *gin.Context) {

	var ambulances entity.Ambulance

	if err := c.ShouldBindJSON(&ambulances); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return

	}

	if tx := entity.DB().Where("id = ?", ambulances.ID).First(&ambulances); tx.RowsAffected == 0 {

		c.JSON(http.StatusBadRequest, gin.H{"error": "user not found"})

		return

	}

	if err := entity.DB().Save(&ambulances).Error; err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return

	}

	c.JSON(http.StatusOK, gin.H{"data": ambulances})

}
