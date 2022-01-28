package entity

import (
	"testing"
	"time"

	"github.com/asaskevich/govalidator"
	. "github.com/onsi/gomega"
)

func TestNumberPatientBeInValid(t *testing.T){
	g := NewGomegaWithT(t)

	ambulancearrival := AmbulanceArrival {
		Number_of_passenger: -6 , // ผิด
		Distance: 2.8,
		DateTime: time.Now(),
	}
	
	ok, err := govalidator.ValidateStruct(ambulancearrival)

	g.Expect(ok).ToNot(BeTrue())

	g.Expect(err).ToNot(BeNil())

	g.Expect(err).ToNot(Equal("Number Patien must be in valid"))
}