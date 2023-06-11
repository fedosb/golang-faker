# golang-faker
### Object Generator with Random Test Data

This program is designed to generate random test data in Go. It provides convenient functions for generating values of various types, such as `bool`, `int`, `float`, `string`, `struct`, `slice`, `map`, and others.

## Usage

1. Create an instance of Faker for the desired type:
```go
myFaker := faker.NewFaker[int]()
```

2. Use the Create() method to generate a test value of the specified type:
```go
res := myFaker.Create()
fmt.Printf("Result: %+v\n", *res)
```

3. You can use the Faker to generate values of various types:
```go
// Generating a test value of type bool
myBoolFaker := faker.NewFaker[bool]()
boolRes := myBoolFaker.Create()

// Generating a test value of type string
myStringFaker := faker.NewFaker[string]()
stringRes := myStringFaker.Create()

// Generating a test value of type struct
type MyStruct struct {
    Name  string
    Age   int
    Email string
}

myStructFaker := faker.NewFaker[MyStruct]()
structRes := myStructFaker.Create()
```

### Testing
To demonstrate the functionality, the `testify` package is used. The example provided in the `faker_test.go` file showcases the functionality of the Faker and testing objects. You can also add test scenarios to the `cases` variable in the `Test_Faker` function in the `faker_test.go` file. Each demonstration (or test) scenario is represented as a structure with two fields: `name` (the name of the test) and `runFunc` (the function containing the test scenario code).