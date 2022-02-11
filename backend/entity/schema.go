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

	Symptom      string
	SymptomLevel uint
	Datetime     time.Time

	PatientID *uint
	Patient   Patient `gorm:"references:id"`

	RecorderID *uint
	Recorder   Employee `gorm:"references:id"`

	IncidentID *uint
	Incident   Incident `gorm:"references:id"`
}

type Ambulance struct {
	gorm.Model
	CarNumber    uint
	Registration string
	DateTime     time.Time

	StatusID *uint
	Status   Status `gorm:"references:id"`

	AmbulanceTypeID *uint
	AmbulanceType   AmbulanceType `gorm:"references:id"`

	EmployeeID *uint
	Employee   Employee `gorm:"references:id"`
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
	DateTime time.Time

	DocCode  string
	Severity uint
	Note     string

	//AmbulanceID ทำหน้าที่เป็น FK
	AmbulanceID *uint
	Ambulance   Ambulance `gorm:"references:id"`

	//RecorderID ทำหน้าที่เป็น FK
	RecorderID *uint
	Recorder   Employee `gorm:"references:id"`

	ProblemID *uint
	Problem   Problem `gorm:"references:id"`
}

type Illness struct {
	gorm.Model
	Value    string
	Incident []Incident `gorm:"foreignKey:IllnessID"`
}

type Incident struct {
	gorm.Model
	Title         string    `valid:"required~Title cannot be blank"`
	Informer      string    `valid:"alpha~Informer cannot be number, required~Informer cannot be blank"`
	Numberpatient int       `valid:"positive~Numberpatient cannot be Negative, required~Numberpatient cannot be Zero"`
	Location      string    `valid:"required~Location cannot be blank"`
	Datetime      time.Time `valid:"future~DateTime must be in the future"`

	EmployeeID  *uint
	Employee    Employee
	IllnessID   *uint
	Illness     Illness
	UrgencyID   *uint
	Urgency     Urgency
	Assessments []Assessment `gorm:"foreignKey:IncidentID" valid:"-"`
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

	govalidator.CustomTypeTagMap.Set("future", func(i interface{}, context interface{}) bool {
		t := i.(time.Time)
		return t.After(time.Now().Add(-1 * time.Hour))
	})

	govalidator.CustomTypeTagMap.Set("positive", func(i interface{}, context interface{}) bool {
		n := i.(int)
		return n >= 1
	})

}
