package golang_faker

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

type FakerTestSuite struct {
	suite.Suite
}

func (suite *FakerTestSuite) Test_Faker() {

	cases := map[string]func(){
		"Sample test": func() {

		},
	}

	for name, cs := range cases {
		suite.Run(name, cs)
	}
}

func TestFakerTestSuite(t *testing.T) {
	suite.Run(t, new(FakerTestSuite))
}
