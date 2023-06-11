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

	cases := []struct {
		name    string
		runFunc func()
	}{
		{
			name: "Sample Create bool value",
			runFunc: func() {
				var myFaker faker.Faker[bool]
				myFaker = faker.NewFaker[bool]()
				res := myFaker.Create()
				fmt.Printf("   | ---> Result is: %+v\n", *res)
			},
		},
		{
			name: "Sample Create int value",
			runFunc: func() {
				var myFaker faker.Faker[int]
				myFaker = faker.NewFaker[int]()

				res := myFaker.Create()
				fmt.Printf("   | ---> The int result is: %+v\n", *res)
			},
		},
		{
			name: "Sample Create int8 value",
			runFunc: func() {
				var myFaker faker.Faker[int8]
				myFaker = faker.NewFaker[int8]()

				res := myFaker.Create()
				fmt.Printf("   | ---> The int8 result is: %+v\n", *res)
			},
		},
		{
			name: "Sample Create int16 value",
			runFunc: func() {
				var myFaker faker.Faker[int16]
				myFaker = faker.NewFaker[int16]()

				res := myFaker.Create()
				fmt.Printf("   | ---> The int16 result is: %+v\n", *res)
			},
		},
		{
			name: "Sample Create int32 value",
			runFunc: func() {
				var myFaker faker.Faker[int32]
				myFaker = faker.NewFaker[int32]()

				res := myFaker.Create()
				fmt.Printf("   | ---> The int32 result is: %+v\n", *res)
			},
		},
		{
			name: "Sample Create int64 value",
			runFunc: func() {
				var myFaker faker.Faker[int64]
				myFaker = faker.NewFaker[int64]()

				res := myFaker.Create()
				fmt.Printf("   | ---> The int64 result is: %+v\n", *res)
			},
		},
		{
			name: "Sample Create uint value",
			runFunc: func() {
				var myFaker faker.Faker[uint]
				myFaker = faker.NewFaker[uint]()

				res := myFaker.Create()
				fmt.Printf("   | ---> The uint result is: %+v\n", *res)
			},
		},
		{
			name: "Sample Create uint8 value",
			runFunc: func() {
				var myFaker faker.Faker[uint8]
				myFaker = faker.NewFaker[uint8]()

				res := myFaker.Create()
				fmt.Printf("   | ---> The uint8 result is: %+v\n", *res)
			},
		},
		{
			name: "Sample Create uint16 value",
			runFunc: func() {
				var myFaker faker.Faker[uint16]
				myFaker = faker.NewFaker[uint16]()

				res := myFaker.Create()
				fmt.Printf("   | ---> The uint16 result is: %+v\n", *res)
			},
		},
		{
			name: "Sample Create uint32 value",
			runFunc: func() {
				var myFaker faker.Faker[uint32]
				myFaker = faker.NewFaker[uint32]()

				res := myFaker.Create()
				fmt.Printf("   | ---> The uint32 result is: %+v\n", *res)
			},
		},
		{
			name: "Sample Create uint64 value",
			runFunc: func() {
				var myFaker faker.Faker[uint64]
				myFaker = faker.NewFaker[uint64]()

				res := myFaker.Create()
				fmt.Printf("   | ---> The uint64 result is: %+v\n", *res)
			},
		},
		{
			name: "Sample Create float32 value",
			runFunc: func() {
				var myFaker faker.Faker[float32]
				myFaker = faker.NewFaker[float32]()

				res := myFaker.Create()
				fmt.Printf("   | ---> The float32 result is: %+v\n", *res)
			},
		},
		{
			name: "Sample Create float64 value",
			runFunc: func() {
				var myFaker faker.Faker[float64]
				myFaker = faker.NewFaker[float64]()

				res := myFaker.Create()
				fmt.Printf("   | ---> The float64 result is: %+v\n", *res)
			},
		},
		{
			name: "Sample Create complex64 value",
			runFunc: func() {
				var myFaker faker.Faker[complex64]
				myFaker = faker.NewFaker[complex64]()

				res := myFaker.Create()
				fmt.Printf("   | ---> The complex64 result is: %+v\n", *res)
			},
		},
		{
			name: "Sample Create complex128 value",
			runFunc: func() {
				var myFaker faker.Faker[complex128]
				myFaker = faker.NewFaker[complex128]()

				res := myFaker.Create()
				fmt.Printf("   | ---> The complex128 result is: %+v\n", *res)
			},
		},
		{
			name: "Sample Create string value",
			runFunc: func() {
				var myFaker faker.Faker[string]
				myFaker = faker.NewFaker[string]()

				res := myFaker.Create()
				fmt.Printf("   | ---> The string result is: %+v\n", *res)
			},
		},
		{
			name: "Sample Create array value",
			runFunc: func() {
				var myFaker faker.Faker[[10]uint]
				myFaker = faker.NewFaker[[10]uint]()

				res := myFaker.Create()
				fmt.Printf("   | ---> The array result is: %+v\n", *res)
			},
		},
		{
			name: "Sample Create slice value",
			runFunc: func() {
				var myFaker faker.Faker[[][10]uint]
				myFaker = faker.NewFaker[[][10]uint]()

				res := myFaker.Create()
				fmt.Printf("   | ---> The slice result is: %+v\n", *res)
			},
		},
		{
			name: "Sample Create map value",
			runFunc: func() {
				var myFaker faker.Faker[map[string]complex64]
				myFaker = faker.NewFaker[map[string]complex64]()

				res := myFaker.Create()
				fmt.Printf("   | ---> The map result is: %+v\n", *res)
			},
		},
	}

	for _, cs := range cases {
		suite.Run(cs.name, cs.runFunc)
	}
}

func TestFakerSampleSuite(t *testing.T) {
	suite.Run(t, new(FakerSampleSuite))
}
