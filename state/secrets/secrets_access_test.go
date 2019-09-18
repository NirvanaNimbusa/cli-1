package secrets_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/ActiveState/cli/internal/environment"
	"github.com/ActiveState/cli/internal/failures"
	"github.com/ActiveState/cli/internal/locale"
	"github.com/ActiveState/cli/internal/testhelpers/exiter"
	"github.com/ActiveState/cli/internal/testhelpers/httpmock"
	"github.com/ActiveState/cli/internal/testhelpers/secretsapi_test"
	"github.com/ActiveState/cli/pkg/platform/api"
	secretsapi "github.com/ActiveState/cli/pkg/platform/api/secrets"
	authMock "github.com/ActiveState/cli/pkg/platform/authentication/mock"
	"github.com/ActiveState/cli/pkg/projectfile"
	"github.com/ActiveState/cli/state/secrets"
	"github.com/stretchr/testify/suite"
)

type SecretsAccessTestSuite struct {
	suite.Suite

	projectFile *projectfile.Project

	secretsClient *secretsapi.Client
	secretsMock   *httpmock.HTTPMock
	platformMock  *httpmock.HTTPMock
	authMock      *authMock.Mock
}

func (suite *SecretsAccessTestSuite) BeforeTest(suiteName, testName string) {
	locale.Set("en-US")
	failures.ResetHandled()
	projectfile.Reset()

	root, err := environment.GetRootPath()
	suite.Require().NoError(err, "Should detect root path")
	err = os.Chdir(filepath.Join(root, "state", "secrets", "testdata", "access"))
	suite.Require().NoError(err, "Should chdir")

	secretsClient := secretsapi_test.NewDefaultTestClient("bearing123")
	suite.Require().NotNil(secretsClient)
	suite.secretsClient = secretsClient

	suite.secretsMock = httpmock.Activate(secretsClient.BaseURI)
	suite.platformMock = httpmock.Activate(api.GetServiceURL(api.ServiceMono).String())

	suite.authMock = authMock.Init()
	suite.authMock.MockLoggedin()
}

func (suite *SecretsAccessTestSuite) TestExecute_NoAccess() {
	cmd := secrets.NewCommand(suite.secretsClient)

	suite.platformMock.RegisterWithCode("GET", "/organizations/AccessOrg", 200)
	suite.platformMock.RegisterWithCode("GET", "/organizations/AccessOrg/members", 200)

	ex := exiter.New()
	cmd.Config().Exiter = ex.Exit

	exitCode := ex.WaitForExit(func() {
		cmd.Config().Execute()
	})
	suite.Equal(1, exitCode, "expected exit code to match")
}

func TestSecretsAccessTestSuite(t *testing.T) {
	suite.Run(t, new(SecretsAccessTestSuite))
}
