package entity

import (
	"time"

	"gorm.io/gorm"
)

type Patient struct {
	gorm.Model

	Name  string
	Tel   string
	Email string
}

type AmbulanceOnDuty struct {
	gorm.Model

	Code       string
	OnDutyDate time.Time
	Passenger  uint

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

type Assessment struct {
	gorm.Model

	Symptom      string
	SymptomLevel uint
	Datetime     time.Time

	// PatientID  *uint
	// Patient	  Patient

	// RecorderID *uint
	// Recorder   Employee

	// IncidentID *uint
	// Incident   Incident
}
