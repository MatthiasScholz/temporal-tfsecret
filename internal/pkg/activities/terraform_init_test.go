package activities

import (
	//"context"
	//"fmt"
	"testing"

	"github.com/stretchr/testify/suite"

	"go.temporal.io/sdk/testsuite"
	//"go.temporal.io/sdk/workflow"
)

var EXAMPLE_ROOT = "examples/root"

// func TestTerraformInit(t *testing.T) {
// 	// Prepare
// 	config := TerraformConfig{EXAMPLE_ROOT, nil}
//
// 	//wf := testsuite.WorkflowTestSuite.NewTestWorkflowEnvironment()
// 	env := testsuite.NewActivityEnvironment()
// 	env.RegisterActivity(TerraformInit, config)
//
// 	// Execute
// 	// success, err := TerraformInit(ctx, config)
// }

// import (
// 	"context"
// 	"errors"
// 	"testing"
//
// 	"github.com/stretchr/testify/mock"
// 	"github.com/stretchr/testify/suite"
//
// 	"go.temporal.io/sdk/activity"
// 	"go.temporal.io/sdk/testsuite"
// )
//
type UnitTestSuite struct {
	suite.Suite
	// WorkflowTestSuite is the test suite to run unit tests for workflow/activity.
	testsuite.WorkflowTestSuite

	env *testsuite.TestActivityEnvironment
}

func (s *UnitTestSuite) SetupTest() {
	s.env = s.NewTestActivityEnvironment()
}

// func (s *UnitTestSuite) AfterTest(suiteName, testName string) {
// 	s.env.AssertExpectations(s.T())
// }

func (s *UnitTestSuite) Test_SimpleWorkflow_Success() {
	// Prepare
	//config := TerraformConfig{EXAMPLE_ROOT, nil}
	config := TerraformConfig{"../../../examples/root", nil}
	s.env.RegisterActivity(TerraformInit)

	result, err := s.env.ExecuteActivity(TerraformInit, config)

	s.True(result.HasValue())
	s.NoError(err)
}

// func (s *UnitTestSuite) Test_SimpleWorkflow_ActivityParamCorrect() {
//         s.env.OnActivity(SimpleActivity, mock.Anything, mock.Anything).Return(
//           func(ctx context.Context, value string) (string, error) {
//                 s.Equal("test_success", value)
//                 return value, nil
//         })
//         s.env.ExecuteWorkflow(SimpleWorkflow, "test_success")
//
//         s.True(s.env.IsWorkflowCompleted())
//         s.NoError(s.env.GetWorkflowError())
// }
//
// func (s *UnitTestSuite) Test_SimpleWorkflow_ActivityFails() {
//         s.env.OnActivity(SimpleActivity, mock.Anything, mock.Anything).Return(
//           "", errors.New("SimpleActivityFailure"))
//         s.env.ExecuteWorkflow(SimpleWorkflow, "test_failure")
//
//         s.True(s.env.IsWorkflowCompleted())
//
//         err := s.env.GetWorkflowError()
//         s.Error(err)
//         var applicationErr *temporal.ApplicationError
//         s.True(errors.As(err, &applicationErr))
//         s.Equal("SimpleActivityFailure", applicationErr.Error())
// }

func TestUnitTestSuite(t *testing.T) {
	suite.Run(t, new(UnitTestSuite))
}
