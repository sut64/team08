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

type Assessment struct {
	gorm.Model

	Symptom      string
	SymptomLevel uint
	Datetime     time.Time

	git PatientID  *uint
	Patient	  Patient

	// RecorderID *uint
	// Recorder   Employee

	// IncidentID *uint
	// Incident   Incident
}
