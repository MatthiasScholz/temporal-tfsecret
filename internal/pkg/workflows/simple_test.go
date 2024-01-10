package workflows

import (
	"testing"

	"github.com/stretchr/testify/suite"
	"go.temporal.io/sdk/testsuite"
)

type UnitTestSuite struct {
	suite.Suite
	testsuite.WorkflowTestSuite
}

func TestUnitTestSuite(t *testing.T) {
	suite.Run(t, new(UnitTestSuite))
}

func (s *UnitTestSuite) Test_SampleGreetingsWorkflow() {
	// Prepare
	env := s.NewTestWorkflowEnvironment()
	var a *Activities
	//env.RegisterActivity(a)

	// .Mock activities
	env.OnActivity(a.GetGreeting).Return("Hello", nil)
	env.OnActivity(a.GetName).Return("World", nil)
	env.OnActivity(a.SayGreeting, "Hello", "World").Return("Hello World!", nil)

	// Execute
	env.ExecuteWorkflow(GreetingSample)

	// Check
	s.True(env.IsWorkflowCompleted())
	s.NoError(env.GetWorkflowError())

	env.AssertExpectations(s.T())
}

// import (
// 	//"context"
// 	//"errors"
// 	"testing"
//
// 	//"github.com/stretchr/testify/mock"
// 	"github.com/stretchr/testify/suite"
//
// 	//"go.temporal.io/sdk/activity"
// 	"go.temporal.io/sdk/testsuite"
// )
//
// type UnitTestSuite struct {
// 	suite.Suite
// 	testsuite.WorkflowTestSuite
//
// 	env *testsuite.TestWorkflowEnvironment
// }
//
// func (s *UnitTestSuite) SetupTest() {
// 	s.env = s.NewTestWorkflowEnvironment()
// }
//
// func (s *UnitTestSuite) AfterTest(suiteName, testName string) {
// 	s.env.AssertExpectations(s.T())
// }
//
// func (s *UnitTestSuite) Test_SimpleWorkflow_Success() {
// 	s.env.ExecuteWorkflow(SimpleWorkflow) //, "test_success")
//
// 	s.True(s.env.IsWorkflowCompleted())
// 	s.NoError(s.env.GetWorkflowError())
// }
//
// func (s *UnitTestSuite) Test_SimpleWorkflow_ActivityParamCorrect() {
// 	s.env.OnActivity(SimpleActivity, mock.Anything, mock.Anything).Return(
// 		func(ctx context.Context, value string) (string, error) {
// 			s.Equal("test_success", value)
// 			return value, nil
// 	})
// 	s.env.ExecuteWorkflow(SimpleWorkflow, "test_success")
//
// 	s.True(s.env.IsWorkflowCompleted())
// 	s.NoError(s.env.GetWorkflowError())
// }
//
// func (s *UnitTestSuite) Test_SimpleWorkflow_ActivityFails() {
// 	s.env.OnActivity(SimpleActivity, mock.Anything, mock.Anything).Return(
// 		"", errors.New("SimpleActivityFailure"))
// 	s.env.ExecuteWorkflow(SimpleWorkflow, "test_failure")
//
// 	s.True(s.env.IsWorkflowCompleted())
//
// 	err := s.env.GetWorkflowError()
// 	s.Error(err)
// 	var applicationErr *temporal.ApplicationError
// 	s.True(errors.As(err, &applicationErr))
// 	s.Equal("SimpleActivityFailure", applicationErr.Error())
// }

// func TestUnitTestSuite(t *testing.T) {
// 	suite.Run(t, new(UnitTestSuite))
// }
//
