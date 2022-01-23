package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sut64/team08/entity"
)

// POST /incidents
// func CreateIncident(c *gin.Context) {

// 	var incident entity.Incident
// 	var employee entity.Employee
// 	var illness entity.Illness
// 	var urgency entity.Urgency

// 	// ผลลัพธ์ที่ได้จากขั้นตอนที่ 11 จะถูก bind เข้าตัวแปร incident
// 	if err := c.ShouldBindJSON(&incident); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	// 12: ค้นหา Employee ด้วย ID
// 	if tx := entity.DB().Where("id = ?", incident.EmployeeID).First(&employee); tx.RowsAffected == 0 {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "employee not found"})
// 		return
// 	}

// 	// 13: ค้นหา Illness ด้วย ID
// 	if tx := entity.DB().Where("id = ?", incident.IllnessID).First(&illness); tx.RowsAffected == 0 {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "illness not found"})
// 		return
// 	}

// 	// 14: ค้นหา Urgency ด้วย ID
// 	if tx := entity.DB().Where("id = ?", incident.UrgencyID).First(&urgency); tx.RowsAffected == 0 {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "urgency not found"})
// 		return
// 	}

// 	// 16: สร้าง Incident
// 	ic := entity.Incident{
// 		Employee: employee,          // โยงความสัมพันธ์กับ Entity employee
// 		Illness:  illness,           // โยงความสัมพันธ์กับ Entity illness
// 		Urgency:  urgency,           // โยงความสัมพันธ์กับ Entity Urgency
// 		Datetime: incident.Datetime, // ตั้งค่าฟิลด์ Datetime
// 	}

// 	// 17: บันทึก
// 	if err := entity.DB().Create(&ic).Error; err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}
// 	c.JSON(http.StatusOK, gin.H{"data": ic})
// }

// List /incidents
func ListIncidents(c *gin.Context) {
	var incidents []entity.Incident
	if err := entity.DB().Preload("Employee").Preload("Illness").Preload("Urgency").Raw("SELECT * FROM incidents").Find(&incidents).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": incidents})
}
