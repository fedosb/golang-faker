package golang_faker

import (
	"fmt"
	"github.com/stretchr/testify/suite"
	"golang-faker/pkg/faker"
	"testing"
	"unsafe"
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

				myFaker := faker.NewFaker[bool]()
				res := myFaker.Create()
				fmt.Printf("   | ---> Result is: %+v\n", *res)
			},
		},
		{
			name: "Sample Create int value",
			runFunc: func() {

				myFaker := faker.NewFaker[int]()

				res := myFaker.Create()
				fmt.Printf("   | ---> The int result is: %+v\n", *res)
			},
		},
		{
			name: "Sample Create int8 value",
			runFunc: func() {

				myFaker := faker.NewFaker[int8]()

				res := myFaker.Create()
				fmt.Printf("   | ---> The int8 result is: %+v\n", *res)
			},
		},
		{
			name: "Sample Create int16 value",
			runFunc: func() {

				myFaker := faker.NewFaker[int16]()

				res := myFaker.Create()
				fmt.Printf("   | ---> The int16 result is: %+v\n", *res)
			},
		},
		{
			name: "Sample Create int32 value",
			runFunc: func() {

				myFaker := faker.NewFaker[int32]()

				res := myFaker.Create()
				fmt.Printf("   | ---> The int32 result is: %+v\n", *res)
			},
		},
		{
			name: "Sample Create int64 value",
			runFunc: func() {

				myFaker := faker.NewFaker[int64]()

				res := myFaker.Create()
				fmt.Printf("   | ---> The int64 result is: %+v\n", *res)
			},
		},
		{
			name: "Sample Create uint value",
			runFunc: func() {

				myFaker := faker.NewFaker[uint]()

				res := myFaker.Create()
				fmt.Printf("   | ---> The uint result is: %+v\n", *res)
			},
		},
		{
			name: "Sample Create uint8 value",
			runFunc: func() {

				myFaker := faker.NewFaker[uint8]()

				res := myFaker.Create()
				fmt.Printf("   | ---> The uint8 result is: %+v\n", *res)
			},
		},
		{
			name: "Sample Create uint16 value",
			runFunc: func() {

				myFaker := faker.NewFaker[uint16]()

				res := myFaker.Create()
				fmt.Printf("   | ---> The uint16 result is: %+v\n", *res)
			},
		},
		{
			name: "Sample Create uint32 value",
			runFunc: func() {

				myFaker := faker.NewFaker[uint32]()

				res := myFaker.Create()
				fmt.Printf("   | ---> The uint32 result is: %+v\n", *res)
			},
		},
		{
			name: "Sample Create uint64 value",
			runFunc: func() {

				myFaker := faker.NewFaker[uint64]()

				res := myFaker.Create()
				fmt.Printf("   | ---> The uint64 result is: %+v\n", *res)
			},
		},
		{
			name: "Sample Create float32 value",
			runFunc: func() {

				myFaker := faker.NewFaker[float32]()

				res := myFaker.Create()
				fmt.Printf("   | ---> The float32 result is: %+v\n", *res)
			},
		},
		{
			name: "Sample Create float64 value",
			runFunc: func() {

				myFaker := faker.NewFaker[float64]()

				res := myFaker.Create()
				fmt.Printf("   | ---> The float64 result is: %+v\n", *res)
			},
		},
		{
			name: "Sample Create complex64 value",
			runFunc: func() {

				myFaker := faker.NewFaker[complex64]()

				res := myFaker.Create()
				fmt.Printf("   | ---> The complex64 result is: %+v\n", *res)
			},
		},
		{
			name: "Sample Create complex128 value",
			runFunc: func() {

				myFaker := faker.NewFaker[complex128]()

				res := myFaker.Create()
				fmt.Printf("   | ---> The complex128 result is: %+v\n", *res)
			},
		},
		{
			name: "Sample Create string value",
			runFunc: func() {

				myFaker := faker.NewFaker[string]()

				res := myFaker.Create()
				fmt.Printf("   | ---> The string result is: %+v\n", *res)
			},
		},
		{
			name: "Sample Create array value",
			runFunc: func() {

				myFaker := faker.NewFaker[[10]uint]()

				res := myFaker.Create()
				fmt.Printf("   | ---> The array result is: %+v\n", *res)
			},
		},
		{
			name: "Sample Create slice value",
			runFunc: func() {

				myFaker := faker.NewFaker[[][10]uint]()

				res := myFaker.Create()
				fmt.Printf("   | ---> The slice result is: %+v\n", *res)
			},
		},
		{
			name: "Sample Create map value",
			runFunc: func() {

				myFaker := faker.NewFaker[map[string]complex64]()

				res := myFaker.Create()
				fmt.Printf("   | ---> The map result is: %+v\n", *res)
			},
		},
		{
			name: "Sample Create uintptr value",
			runFunc: func() {

				myFaker := faker.NewFaker[uintptr]()

				res := myFaker.Create()
				fmt.Printf("   | ---> The uintptr result is: %+v\n", *res)
			},
		},
		{
			name: "Sample Create pointer value",
			runFunc: func() {

				myFaker := faker.NewFaker[*uint8]()

				res := myFaker.Create()
				fmt.Printf("   | ---> The pointer result is: %+v\n", *res)
			},
		},
		{
			name: "Sample Create func value",
			runFunc: func() {

				myFaker := faker.NewFaker[func(foo int, bar uint)]()

				res := myFaker.Create()
				fmt.Printf("   | ---> The func result is: %+v\n", res)
			},
		},
		{
			name: "Sample Create chan value",
			runFunc: func() {

				myFaker := faker.NewFaker[chan complex128]()

				res := myFaker.Create()
				fmt.Printf("   | ---> The chan result is: %+v\n", *res)
			},
		},
		{
			name: "Sample Create struct value",
			runFunc: func() {

				type Foo struct {
					FooField []struct {
						FooStructInnerField complex128
					}
				}

				type Bar struct {
					Foo
					BarField map[[20]byte][]*uint8
				}

				type List[T any] struct {
					Data []T
					Meta []struct {
						MetaField string
					}
				}

				myFaker := faker.NewFaker[List[Bar]]()

				res := myFaker.Create()
				fmt.Printf("   | ---> The struct result is: %+v\n", *res)
			},
		},
		{
			name: "Test Invalid chan value",
			runFunc: func() {

				defer func() {
					r := recover()
					suite.NotEqual(nil, r)
				}()

				myFaker := faker.NewFaker[unsafe.Pointer]()

				_ = myFaker.Create()
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
