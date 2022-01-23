package entity

import (
	"time"

	"gorm.io/gorm"
)

type AmbulanceOnDuty struct {
	gorm.Model

	Annotation string
	Date       time.Time

	// AmbulanceID *uint
	// Ambulance   Ambulance

	// RecorderID *uint
	// Recorder   Employee

	// IncidentID *uint
	// Incident   Incident
}

type AmbulanceArrival struct {
	gorm.Model
	Number_of_people int
	Distance         float32
	DateTime         time.Time

	RecorderID *uint
	//Recorder   Employee

	PatientID *uint
	//Patient	Patient

	AmbulanceOnDutyID *uint
	AmbulanceOnDuty   AmbulanceOnDuty
}
