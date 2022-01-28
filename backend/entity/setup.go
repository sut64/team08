package entity

import (
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"

	"gorm.io/driver/sqlite"
)

var db *gorm.DB

func DB() *gorm.DB {
	return db
}

func SetupDatabase() {
	database, err := gorm.Open(sqlite.Open("SE64.db"), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema

	database.AutoMigrate(
		&AmbulanceOnDuty{},
		&AmbulanceArrival{},
		&Patient{},
		&Assessment{},
		&Ambulance{},
		&AmbulanceType{},
		&Employee{},
		&Status{},
		&AmbulanceCheck{},
		&Illness{},
		&Urgency{},
		&Incident{},
	)

	db = database
	password, err := bcrypt.GenerateFromPassword([]byte("123456"), 14)

	//employee
	db.Model(&Employee{}).Create(&Employee{
		Name:     "นางสาวพร มณีวรรณ",
		Email:    "porn@gmail.com",
		Tel:      "0883322456",
		Password: string(password),
	})
	db.Model(&Employee{}).Create(&Employee{
		Name:     "นายสม จันทร์เพ็ญ",
		Email:    "som@gmail.com",
		Tel:      "0885548900",
		Password: string(password),
	})
	db.Model(&Employee{}).Create(&Employee{
		Name:     "นางสาวกล้วย ไชยวาที",
		Email:    "naruemon@gmail.com",
		Tel:      "0610091572",
		Password: string(password),
	})
	db.Model(&Employee{}).Create(&Employee{
		Name:     "Phupha Bbbb",
		Email:    "bbb@gmail.com",
		Tel:      "0945333333",
		Password: string(password),
	})

	var porn Employee
	var som Employee
	var kluy Employee
	var Phupha Employee

	db.Raw("SELECT * FROM employees WHERE email = ?", "porn@gmail.com").Scan(&porn)
	db.Raw("SELECT * FROM employees WHERE email = ?", "som@gmail.com").Scan(&som)
	db.Raw("SELECT * FROM employees WHERE email = ?", "naruemon@gmail.com").Scan(&kluy)
	db.Raw("SELECT * FROM employees WHERE email = ?", "bbb@gmail.com").Scan(&Phupha)

	//status
	available := Status{
		Detail: "ว่าง",
	}
	db.Model(&Status{}).Create(&available)

	noavailable := Status{
		Detail: "ไม่ว่าง",
	}
	db.Model(&Status{}).Create(&noavailable)

	//AmbulanceType
	Standard := AmbulanceType{
		Name:   "VAN-ALSV",
		Detail: "รถตู้พยาบาลกู้ชีพขั้นสูง มีชุดรางอินทิเกรต, Plasma Generator, Negative Ion Generator, ถังออกซิเจน, เตียงฉุกเฉิน, เก้าอี้ลำเลียงผู้ป่วย, กล่องจ่ายไฟ, สัญญาณเตือนวิทยุ",
	}
	db.Model(&AmbulanceType{}).Create(&Standard)

	Deluxe := AmbulanceType{
		Name:   "Box-Body",
		Detail: "รถกระบะแบบกล่องฉุกเฉิน มีชุดแอร์ฟอกอากาศ, ตู้เก็บถังออกซิเจน, แผงควบคุม, ชุดอุปกรณ์เพดาน, ถังบรรจุออกซิเจน",
	}
	db.Model(&AmbulanceType{}).Create(&Deluxe)

	Suite := AmbulanceType{
		Name:   "Half-Body",
		Detail: "รถกระบะแบบโดมรถตู้ มีชุดแอร์แขวนฟอกอากาศ, ตู้เก็บถังออกซิเจน, แผงควบคุม, ชุดไฟส่องสว่าง L.E.D., ถังบรรจุออกซิเจน",
	}
	db.Model(&AmbulanceType{}).Create(&Suite)

	//Ambulance
	am1 := Ambulance{
		CarNumber:    100,
		Registration: "กข 523",
		DateTime:     time.Now(),
	}
	db.Model(&Ambulance{}).Create(&am1)

	am2 := Ambulance{
		CarNumber:    101,
		Registration: "อบ 7448",
		DateTime:     time.Now(),
	}
	db.Model(&Ambulance{}).Create(&am2)

	// Illness
	death := Illness{
		Value: "ไม่ได้สติ",
	}
	db.Model(&Illness{}).Create(&death)

	// Urgency
	urfast := Urgency{
		Value: "เร็ว (Fast)",
	}
	db.Model(&Urgency{}).Create(&urfast)

	urmedium := Urgency{
		Value: "ปานกลาง (Medium)",
	}
	db.Model(&Urgency{}).Create(&urmedium)

	urslow := Urgency{
		Value: "ช้า (Slow)",
	}
	db.Model(&Urgency{}).Create(&urslow)

	// Incident
	// 1
	db.Model(&Incident{}).Create(&Incident{
		Title:         "ทดสอบ",
		Informer:      "นายปรเมต สมอะไร",
		Numberpatient: 2,
		Location:      "บ้านปรเมต",
		Datetime:      time.Now(),
		Employee:      porn,
		Illness:       death,
		Urgency:       urfast,
	})
}
