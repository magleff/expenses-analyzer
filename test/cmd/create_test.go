package cmd_test

import (
	"github.com/golang/mock/gomock"
	"github.com/magleff/gobro/cmd"
	"github.com/magleff/gobro/common"
	"github.com/magleff/gobro/entities"
	"github.com/magleff/gobro/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"gopkg.in/mgo.v2/bson"
	"testing"
)

// --- SETUP ---

type CreateCmdTestSuite struct {
	suite.Suite
	MockAccountController *mocks.MockAccountController
	MockBudgetController  *mocks.MockBudgetController
	MockUserSession       *mocks.MockUserSession
	MockController        *gomock.Controller
}

func (suite *CreateCmdTestSuite) SetupTest() {
	suite.MockController = gomock.NewController(suite.T())
	suite.MockAccountController = mocks.NewMockAccountController(suite.MockController)
	suite.MockBudgetController = mocks.NewMockBudgetController(suite.MockController)
	suite.MockUserSession = mocks.NewMockUserSession(suite.MockController)
}

func (suite *CreateCmdTestSuite) TearDownTest() {
	suite.MockController.Finish()
}

func TestCreateCmdTestSuite(t *testing.T) {
	suite.Run(t, new(CreateCmdTestSuite))
}

// --- TESTS ---

func (suite *CreateCmdTestSuite) TestCreateCmdNoType() {
	// Arrange
	user := &entities.User{
		ID: bson.NewObjectId(),
	}

	// Act
	err := cmd.CreateCmdDo([]string{}, suite.MockAccountController, user)

	// Assert
	assert.Error(suite.T(), err, "Should throw an error if the type is not provided.")
}

func (suite *CreateCmdTestSuite) TestCreateAccountCmd() {
	// Arrange
	accountName := "main"
	user := &entities.User{
		ID: bson.NewObjectId(),
	}

	suite.MockAccountController.EXPECT().Create(user.ID, accountName).Return(nil)

	// Act
	err := cmd.CreateCmdDo([]string{common.TYPE_ACCOUNT, accountName},
		suite.MockAccountController, user)

	// Assert
	assert.NoError(suite.T(), err, "Should not throw an error.")
}

func (suite *CreateCmdTestSuite) TestCreateAccountNoName() {
	// Arrange
	user := &entities.User{
		ID: bson.NewObjectId(),
	}

	// Act
	err := cmd.CreateCmdDo([]string{common.TYPE_ACCOUNT}, suite.MockAccountController, user)

	// Assert
	assert.Error(suite.T(), err, "Should throw an error if the name of the account is not provided.")
}

// func (suite *CreateCmdTestSuite) TestCreateBudgetCmd() {
// 	// Arrange
// 	user := &entities.User{
// 		ID: bson.NewObjectId(),
// 	}

// 	suite.MockBudgetController.EXPECT().Create(user.ID).Return(nil)

// 	// Act
// 	err := cmd.CreateCmdDo([]string{common.TYPE_BUDGET}, suite.MockBudgetController, user)

// 	// Assert
// 	assert.NoError(suite.T(), err, "Should not throw an error.")
// }
