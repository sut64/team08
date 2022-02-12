package entity

import (
	"testing"
	"time"

	"github.com/asaskevich/govalidator"
	. "github.com/onsi/gomega"
)

func TestAmbulanceCheckPass(t *testing.T) {
	g := NewGomegaWithT(t)

	// ข้อมูลถูกต้องหมดทุก field
	ambulanceCheck := AmbulanceCheck{
		DateTime: time.Now(),
		DocCode:  "AAA123",
		Severity: 2,
		Note:     "This is a test",
	}

	ok, err := govalidator.ValidateStruct(ambulanceCheck)

	g.Expect(ok).To(BeTrue())

	g.Expect(err).To(BeNil())
}

func TestAmbulanceCheckDocCodeMatches(t *testing.T) {
	g := NewGomegaWithT(t)

	fixtures := []string{
		"",       // ว่าง
		"DADD34", // ตัวอักษรเกิน
		"123458", // ไม่มีตัวอักษร
		"D12379", //ตัวเลขเกิน
		"D1267",  //ตัวเลขไม่ครบ
		"D",      //ไม่มีตัวเลข
	}

	for _, fixture := range fixtures {
		ambulanceCheck := AmbulanceCheck{
			DateTime: time.Now(),
			DocCode:  fixture,
			Severity: 2,
			Note:     "This is a test",
		}

		ok, err := govalidator.ValidateStruct(ambulanceCheck)

		g.Expect(ok).ToNot(BeTrue())

		g.Expect(err).ToNot(BeNil())
	}
}

func TestAmbulanceCheckDateTimeMustBeToday(t *testing.T) {
	g := NewGomegaWithT(t)

	ambulanceCheck := AmbulanceCheck{
		DateTime: time.Now().Add(24 * time.Hour), // อนาคต, fail
		DocCode:  "AAA123",
		Severity: 2,
		Note:     "This is a test",
	}

	ok, err := govalidator.ValidateStruct(ambulanceCheck)

	g.Expect(ok).ToNot(BeTrue())

	g.Expect(err).ToNot(BeNil())

	g.Expect(err.Error()).To(Equal("Time must be current date"))
}
func TestAmbulanceCheckSeverityInRange(t *testing.T) {
	g := NewGomegaWithT(t)

	ambulanceCheck := AmbulanceCheck{
		DateTime: time.Now(),
		DocCode:  "AAA123",
		Severity: 0,
		Note:     "This is a test",
	}

	ok, err := govalidator.ValidateStruct(ambulanceCheck)

	g.Expect(ok).ToNot(BeTrue())

	g.Expect(err).ToNot(BeNil())

	g.Expect(err.Error()).To(Equal("Level must be between 1-3"))
}
