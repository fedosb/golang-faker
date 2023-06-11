package golang_faker

import (
	"fmt"
	"github.com/stretchr/testify/suite"
	"golang-faker/pkg/faker"
	"testing"
)

type FakerSampleSuite struct {
	suite.Suite
}

func (suite *FakerSampleSuite) Test_Faker() {

	cases := map[string]func(){
		"Sample Create bool value": func() {
			var myFaker faker.Faker[bool]
			myFaker = faker.NewFaker[bool]()

			res := myFaker.Create()
			fmt.Printf("   | ---> Result is: %+v\n", *res)
		},
		"Sample Create int value": func() {
			var myFaker faker.Faker[int]
			myFaker = faker.NewFaker[int]()

			res := myFaker.Create()
			fmt.Printf("   | ---> The int result is: %+v\n", *res)
		},
		"Sample Create int8 value": func() {
			var myFaker faker.Faker[int8]
			myFaker = faker.NewFaker[int8]()

			res := myFaker.Create()
			fmt.Printf("   | ---> The int8 result is: %+v\n", *res)
		},
		"Sample Create int16 value": func() {
			var myFaker faker.Faker[int16]
			myFaker = faker.NewFaker[int16]()

			res := myFaker.Create()
			fmt.Printf("   | ---> The int16 result is: %+v\n", *res)
		},
		"Sample Create int32 value": func() {
			var myFaker faker.Faker[int32]
			myFaker = faker.NewFaker[int32]()

			res := myFaker.Create()
			fmt.Printf("   | ---> The int32 result is: %+v\n", *res)
		},
		"Sample Create int64 value": func() {
			var myFaker faker.Faker[int64]
			myFaker = faker.NewFaker[int64]()

			res := myFaker.Create()
			fmt.Printf("   | ---> The int64 result is: %+v\n", *res)
		},
		"Sample Create uint value": func() {
			var myFaker faker.Faker[uint]
			myFaker = faker.NewFaker[uint]()

			res := myFaker.Create()
			fmt.Printf("   | ---> The uint result is: %+v\n", *res)
		},
		"Sample Create uint8 value": func() {
			var myFaker faker.Faker[uint8]
			myFaker = faker.NewFaker[uint8]()

			res := myFaker.Create()
			fmt.Printf("   | ---> The uint8 result is: %+v\n", *res)
		},
		"Sample Create uint16 value": func() {
			var myFaker faker.Faker[uint16]
			myFaker = faker.NewFaker[uint16]()

			res := myFaker.Create()
			fmt.Printf("   | ---> The uint16 result is: %+v\n", *res)
		},
		"Sample Create uint32 value": func() {
			var myFaker faker.Faker[uint32]
			myFaker = faker.NewFaker[uint32]()

			res := myFaker.Create()
			fmt.Printf("   | ---> The uint32 result is: %+v\n", *res)
		},
		"Sample Create uint64 value": func() {
			var myFaker faker.Faker[uint64]
			myFaker = faker.NewFaker[uint64]()

			res := myFaker.Create()
			fmt.Printf("   | ---> The uint64 result is: %+v\n", *res)
		},
		"Sample Create float32 value": func() {
			var myFaker faker.Faker[float32]
			myFaker = faker.NewFaker[float32]()

			res := myFaker.Create()
			fmt.Printf("   | ---> The float32 result is: %+v\n", *res)
		},
		"Sample Create float64 value": func() {
			var myFaker faker.Faker[float64]
			myFaker = faker.NewFaker[float64]()

			res := myFaker.Create()
			fmt.Printf("   | ---> The float64 result is: %+v\n", *res)
		},
		"Sample Create complex64 value": func() {
			var myFaker faker.Faker[complex64]
			myFaker = faker.NewFaker[complex64]()

			res := myFaker.Create()
			fmt.Printf("   | ---> The complex64 result is: %+v\n", *res)
		},
		"Sample Create complex128 value": func() {
			var myFaker faker.Faker[complex128]
			myFaker = faker.NewFaker[complex128]()

			res := myFaker.Create()
			fmt.Printf("   | ---> The complex128 result is: %+v\n", *res)
		},
		"Sample Create string value": func() {
			var myFaker faker.Faker[string]
			myFaker = faker.NewFaker[string]()

			res := myFaker.Create()
			fmt.Printf("   | ---> The complex128 result is: %+v\n", *res)
		},
	}

	for name, cs := range cases {
		suite.Run(name, cs)
	}
}

func TestFakerSampleSuite(t *testing.T) {
	suite.Run(t, new(FakerSampleSuite))
}
