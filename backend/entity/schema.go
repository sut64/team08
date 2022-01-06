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
