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
