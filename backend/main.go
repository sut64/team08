package main

import (
	"github.com/sut64/team08/controller"

	"github.com/sut64/team08/entity"
	"github.com/sut64/team08/middlewares"

	"github.com/gin-gonic/gin"
)

func main() {
	entity.SetupDatabase()

	r := gin.Default()
	r.Use(CORSMiddleware())

	api := r.Group("")
	{
		protected := api.Use(middlewares.Authorizes())
		{
			// Ambulance Routes
			protected.GET("/ambulances", controller.ListAmbulances)
			protected.GET("/ambulance/:id", controller.GetAmbulance)
			protected.POST("/ambulances", controller.CreateAmbulance)
			protected.PATCH("/ambulances", controller.UpdateAmbulance)
			protected.DELETE("/ambulances/:id", controller.DeleteAmbulance)

			// Customer Routes
			protected.GET("/statuses", controller.ListStatuses)
			protected.GET("/status/:id", controller.GetStatus)
			protected.POST("/statuses", controller.CreateStatus)
			protected.PATCH("/statuses", controller.UpdateStatus)
			protected.DELETE("/statuses/:id", controller.DeleteStatus)

			// RoomPayment Routes
			protected.GET("/ambulancetypes", controller.ListAmbulancesTypes)
			protected.GET("/ambulancetype/:id", controller.GetAmbulanceType)
			protected.POST("/ambulancetypes", controller.CreateAmbulanceType)
			protected.PATCH("/ambulancetypes", controller.UpdateAmbulanceType)
			protected.DELETE("/ambulancetypes/:id", controller.DeleteAmbulanceType)

			// Employee Routes
			protected.GET("/employees", controller.ListEmployees)
			protected.GET("/employee/:id", controller.GetEmployee)
			//protected.POST("/employees", controller.CreateEmployee)
			protected.PATCH("/employees", controller.UpdateEmployee)
			protected.DELETE("/employees/:id", controller.DeleteEmployee)

			// AmbulanceArrival Routes
			protected.GET("/amnluncearrivals", controller.ListAmbulanceArrivals)
			protected.GET("/amnluncearrival/:id", controller.GetAmbulanceArrival)
			protected.POST("/amnluncearrivals", controller.CreateAmbulanceArrival)
			protected.PATCH("/amnluncearrivals", controller.UpdateAmbulanceArrival)
			protected.DELETE("/amnluncearrivals/:id", controller.DeleteAmbulanceArrival)

			// AmbulanceOnDuty Routes
			protected.GET("/ambulanceonduties", controller.ListAmbulanceOnDutys)
			protected.GET("/amnlunceonduty/:id", controller.GetAmbulanceOnDuty)
			protected.GET("/ambulanceonduties/ambulance", controller.ListAmbulanceOnDutyAmbulance)
			//protected.POST("/ambulanceondutise", controller.CreateAmbulanceOnDuty)
			//protected.PATCH("/ambulanceondutise", controller.UpdateAmbulanceOnDuty)
			//protected.DELETE("/ambulanceondutise/:id", controller.DeleteAmbulanceOnDuty)

			// Patient Routes
			protected.GET("/patients", controller.ListPatients)
			protected.GET("/patient/:id", controller.GetPatient)
			//protected.POST("/patients", controller.CreatePatient)
			protected.PATCH("/patients", controller.UpdatePatient)
			protected.DELETE("/patients/:id", controller.DeletePatient)
		}
	}

	// emp Routes
	r.POST("/employees", controller.CreateEmployee)

	// Authentication Routes
	r.POST("/login", controller.Login)

	// Run the server
	r.Run()
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
