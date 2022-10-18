package main_test

import (
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/milligan22963/passphrase/cmd"
)

type MainTestSuite struct {
	suite.Suite
}

func (suite *MainTestSuite) TestRootCmd() {
	cmd.Execute()
}

func TestMainTestSuite(t *testing.T) {
	suite.Run(t, new(MainTestSuite))
}
