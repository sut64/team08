package controller

import (
	"testing"
	"time"

	"github.com/asaskevich/govalidator"
	. "github.com/onsi/gomega"
	"github.com/sut64/team08/entity"
)

func TestAmbulanceOnDutyPass(t *testing.T) {
	g := NewGomegaWithT(t)

	// ข้อมูลถูกต้องหมดทุก field
	ambulanceOnDuty := entity.AmbulanceOnDuty{
		Code:       "D00000000",
		OnDutyDate: time.Now(),
		Passenger:  2,
	}
	// ตรวจสอบด้วย govalidator
	ok, err := govalidator.ValidateStruct(ambulanceOnDuty)

	// ok ต้องเป็น true แปลว่าไม่มี error
	g.Expect(ok).To(BeTrue())

	// err เป็นค่า nil แปลว่าไม่มี error
	g.Expect(err).To(BeNil())
}

func TestAmbulanceOnDutyCodeNotBlank(t *testing.T) {
	g := NewGomegaWithT(t)

	// ข้อมูลถูกต้องหมดทุก field
	ambulanceOnDuty := entity.AmbulanceOnDuty{
		Code:       "",
		OnDutyDate: time.Now(),
		Passenger:  2,
	}
	// ตรวจสอบด้วย govalidator
	ok, err := govalidator.ValidateStruct(ambulanceOnDuty)

	// ok ต้องเป็น true แปลว่าไม่มี error
	g.Expect(ok).ToNot(BeTrue())

	// err เป็นค่า nil แปลว่าไม่มี error
	g.Expect(err).ToNot(BeNil())
}

func TestAmbulanceOnDutyCodeMatches(t *testing.T) {
	g := NewGomegaWithT(t)

	fixtures := []string{
		"A12345678",  // ตัวอักษรผิด
		"DA12345678", // ตัวอักษรเกิน
		"12345678",   // ไม่มีตัวอักษร

		"D123456789", //ตัวเลขเกิน
		"D1234567",   //ตัวเลขไม่ครบ
		"D",          //ไม่มีตัวเลข
	}

	for _, fixture := range fixtures {
		ambulanceOnDuty := entity.AmbulanceOnDuty{
			Code:       fixture,
			OnDutyDate: time.Now(),
			Passenger:  2,
		}
		// ตรวจสอบด้วย govalidator
		ok, err := govalidator.ValidateStruct(ambulanceOnDuty)

		// ok ต้องเป็น true แปลว่าไม่มี error
		g.Expect(ok).ToNot(BeTrue())

		// err เป็นค่า nil แปลว่าไม่มี error
		g.Expect(err).ToNot(BeNil())
	}
}

func TestAmbulanceOnDutyDateToDay(t *testing.T) {
	g := NewGomegaWithT(t)

	// ข้อมูลถูกต้องหมดทุก field
	ambulanceOnDuty := entity.AmbulanceOnDuty{
		Code:       "D12345678",
		OnDutyDate: time.Now().Add(24 * time.Hour),
		Passenger:  2,
	}
	// ตรวจสอบด้วย govalidator
	ok, err := govalidator.ValidateStruct(ambulanceOnDuty)

	// ok ต้องเป็น true แปลว่าไม่มี error
	g.Expect(ok).ToNot(BeTrue())

	// err เป็นค่า nil แปลว่าไม่มี error
	g.Expect(err).ToNot(BeNil())
}

func TestAmbulanceOnDutyPassengerNotZero(t *testing.T) {
	g := NewGomegaWithT(t)

	// ข้อมูลถูกต้องหมดทุก field
	ambulanceOnDuty := entity.AmbulanceOnDuty{
		Code:       "D12345678",
		OnDutyDate: time.Now(),
		Passenger:  0,
	}
	// ตรวจสอบด้วย govalidator
	ok, err := govalidator.ValidateStruct(ambulanceOnDuty)

	// ok ต้องเป็น true แปลว่าไม่มี error
	g.Expect(ok).ToNot(BeTrue())

	// err เป็นค่า nil แปลว่าไม่มี error
	g.Expect(err).ToNot(BeNil())
}
