package runtime

import (
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/ActiveState/cli/internal/failures"
	"github.com/ActiveState/cli/internal/testhelpers/httpmock"
	"github.com/ActiveState/cli/pkg/platform/api"
	authMock "github.com/ActiveState/cli/pkg/platform/authentication/mock"
	"github.com/ActiveState/cli/pkg/platform/model"
	"github.com/ActiveState/cli/pkg/projectfile"
)

type InternalTestSuite struct {
	suite.Suite

	cacheDir    string
	downloadDir string
	installer   *Installer
	authMock    *authMock.Mock
	httpmock    *httpmock.HTTPMock
}

func (suite *InternalTestSuite) BeforeTest(suiteName, testName string) {
	suite.authMock = authMock.Init()
	suite.authMock.MockLoggedin()
	suite.httpmock = httpmock.Activate(api.GetServiceURL(api.ServiceMono).String())

	pjfile := projectfile.Project{
		Project: model.ProjectURL("string", "string", "00010001-0001-0001-0001-000100010001"),
	}
	pjfile.Persist()

	var err error
	suite.cacheDir, err = ioutil.TempDir("", "")
	suite.Require().NoError(err)

	suite.downloadDir, err = ioutil.TempDir("", "cli-installer-test-download")
	suite.Require().NoError(err)

	var fail *failures.Failure
	suite.installer, fail = NewInstaller(suite.downloadDir, suite.cacheDir, InitDownload(suite.downloadDir))
	suite.Require().NoError(fail.ToError())
	suite.Require().NotNil(suite.installer)
}

func (suite *InternalTestSuite) AfterTest(suiteName, testName string) {
	projectfile.Reset()
	suite.authMock.Close()
	httpmock.DeActivate()
}

func (suite *InternalTestSuite) TestValidateCheckpointNoCommit() {
	pjfile := projectfile.Project{
		Project: model.ProjectURL("string", "string", ""),
	}
	pjfile.Persist()

	fail := suite.installer.validateCheckpoint()
	suite.Equal(FailNoCommits.Name, fail.Type.Name)
}

func (suite *InternalTestSuite) TestValidateCheckpointPrePlatform() {
	httpmock.RegisterWithResponseBody(
		"GET", "/vcs/commits/00010001-0001-0001-0001-000100010001/checkpoint", 200,
		`[{"namespace": "pre-platform-installer"}]`)
	fail := suite.installer.validateCheckpoint()
	suite.Equal(FailPrePlatformNotSupported.Name, fail.Type.Name)
}

func TestInternalTestSuite(t *testing.T) {
	suite.Run(t, new(InternalTestSuite))
}