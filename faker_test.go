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
			fmt.Printf("   | ---> Result is: %+v\n", *res)
		},
		"Sample Create uint value": func() {
			var myFaker faker.Faker[uint]
			myFaker = faker.NewFaker[uint]()

			res := myFaker.Create()
			fmt.Printf("   | ---> Result is: %+v\n", *res)
		},
		"Sample Create float value": func() {
			var myFaker faker.Faker[float64]
			myFaker = faker.NewFaker[float64]()

			res := myFaker.Create()
			fmt.Printf("   | ---> Result is: %+v\n", *res)
		},
		"Sample Create complex value": func() {
			var myFaker faker.Faker[complex128]
			myFaker = faker.NewFaker[complex128]()

			res := myFaker.Create()
			fmt.Printf("   | ---> Result is: %+v\n", *res)
		},
	}

	for name, cs := range cases {
		suite.Run(name, cs)
	}
}

func TestFakerSampleSuite(t *testing.T) {
	suite.Run(t, new(FakerSampleSuite))
}
