package entity

import (
	"time"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type Patient struct {
	gorm.Model

	Name             string
	Tel              string
	Email            string
	Assessments      []Assessment       `gorm:"foreignKey:PatientID"`
	AmbulanceArrival []AmbulanceArrival `gorm:"foreignKey:PatientID"`
}

type AmbulanceOnDuty struct {
	gorm.Model

	Code       string    `valid:"matches(^[D]\\d{8}$),required"`
	OnDutyDate time.Time `valid:"today"`
	Passenger  uint      `valid:"required"`

	AmbulanceID *uint
	Ambulance   Ambulance `gorm:"references:id" valid:"-"`

	RecorderID *uint
	Recorder   Employee `gorm:"references:id" valid:"-"`

	IncidentID *uint
	Incident   Incident `gorm:"references:id" valid:"-"`

	AmbulanceArrival []AmbulanceArrival `gorm:"foreignKey:AmbulanceOnDutyID"`
}

type AmbulanceArrival struct {
	gorm.Model
	Number_of_passenger uint
	Distance            float32
	DateTime            time.Time

	RecorderID *uint
	Recorder   Employee `gorm:"references:id"`

	PatientID *uint
	Patient   Patient `gorm:"references:id"`

	AmbulanceOnDutyID *uint           `gorm:"uniqueIndex"`
	AmbulanceOnDuty   AmbulanceOnDuty `gorm:"references:id"`
}

type Assessment struct {
	gorm.Model

	Symptom      string    `valid:"required~Please fill the symptom"`
	SymptomLevel uint      `valid:"int,range(1|3),required~Level must be only (1-3)"`
	Datetime     time.Time `valid:"required,today~Please select current time"`

	PatientID *uint
	Patient   Patient `gorm:"references:id" valid:"-"`

	RecorderID *uint
	Recorder   Employee `gorm:"references:id" valid:"-"`

	IncidentID *uint
	Incident   Incident `gorm:"references:id" valid:"-"`
}

type Ambulance struct {
	gorm.Model
	CarNumber    int       `valid:"required,mtzero"`
	Registration string    `valid:"matches(^[ก-ฮ]{2}\\d{1}$|^[ก-ฮ]{2}\\d{2}$|^[ก-ฮ]{2}\\d{3}$|^[ก-ฮ]{2}\\d{4}$|^1[ก-ฮ]{2}\\d{1}$|^1[ก-ฮ]{2}\\d{2}$|^1[ก-ฮ]{2}\\d{3}$|^1[ก-ฮ]{2}\\d{4}$)"`
	DateTime     time.Time `valid:"today~AmbulanceDate not true"`

	StatusID *uint
	Status   Status `gorm:"references:id" valid:"-"`

	AmbulanceTypeID *uint
	AmbulanceType   AmbulanceType `gorm:"references:id" valid:"-"`

	EmployeeID *uint
	Employee   Employee `gorm:"references:id" valid:"-"`
}
type Status struct {
	gorm.Model
	Detail  string
	Records []Ambulance `gorm:"foreignKey:StatusID"`
}
type AmbulanceType struct {
	gorm.Model
	Name    string
	Detail  string
	Records []Ambulance `gorm:"foreignKey:AmbulanceTypeID"`
}
type Employee struct {
	gorm.Model
	Name             string
	Tel              string
	Email            string `gorm:"uniqueIndex"`
	Password         string
	Records          []Ambulance        `gorm:"foreignKey:EmployeeID"`
	Assessments      []Assessment       `gorm:"foreignKey:RecorderID"`
	Incident         []Incident         `gorm:"foreignKey:EmployeeID"`
	AmbulanceArrival []AmbulanceArrival `gorm:"foreignKey:RecorderID"`
}

type AmbulanceCheck struct {
	gorm.Model
	DateTime time.Time `valid:"today~Time must be current date"`

	DocCode  string `valid:"matches(^[A-Z]{3}\\d{3}$),required~DocCode must be correct in form"`
	Severity int    `valid:"int,range(1|3),required~Level must be between 1-3"`
	Note     string

	//AmbulanceID ทำหน้าที่เป็น FK
	AmbulanceID *uint
	Ambulance   Ambulance `gorm:"references:id" valid:"-"`

	//RecorderID ทำหน้าที่เป็น FK
	RecorderID *uint
	Recorder   Employee `gorm:"references:id" valid:"-"`

	ProblemID *uint
	Problem   Problem `gorm:"references:id" valid:"-"`
}

type Illness struct {
	gorm.Model
	Value    string
	Incident []Incident `gorm:"foreignKey:IllnessID"`
}

type Incident struct {
	gorm.Model
	Title         string
	Informer      string
	Numberpatient uint
	Location      string
	Datetime      time.Time
	EmployeeID    *uint
	Employee      Employee
	IllnessID     *uint
	Illness       Illness
	UrgencyID     *uint
	Urgency       Urgency
	Assessments   []Assessment `gorm:"foreignKey:IncidentID"`
}

type Urgency struct {
	gorm.Model
	Value    string
	Incident []Incident `gorm:"foreignKey:UrgencyID"`
}

type Problem struct {
	gorm.Model
	Name string
}

func init() {
	govalidator.CustomTypeTagMap.Set("today", func(i interface{}, context interface{}) bool {
		t := i.(time.Time)

		if t.Year() == time.Now().Year() {
			if int(t.Month()) == int(time.Now().Month()) {
				if t.Day() == time.Now().Day() {
					return true
				}
			}
		}
		return false
	})
	govalidator.CustomTypeTagMap.Set("mtzero", func(i interface{}, context interface{}) bool {
		c, _ := i.(int)
		return c > 0
	})
}
