package integration

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"runtime"
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/ActiveState/cli/internal/constants"
	"github.com/ActiveState/cli/internal/environment"
	"github.com/ActiveState/cli/internal/testhelpers/e2e"
	"github.com/ActiveState/termtest"
)

type UpdateGenIntegrationTestSuite struct {
	suite.Suite
}

func (suite *UpdateGenIntegrationTestSuite) TestUpdateBits() {
	root := environment.GetRootPathUnsafe()

	ext := ".tar.gz"
	exe := ""
	if runtime.GOOS == "windows" {
		ext = ".zip"
		exe = ".exe"
	}
	platform := runtime.GOOS + "-" + runtime.GOARCH

	archivePath := filepath.Join(root, "build/update", constants.BranchName, constants.Version, platform+ext)
	suite.Require().FileExists(archivePath, "Make sure you ran 'state run generate-update'")
	suite.T().Logf("file %s exists\n", archivePath)

	tempPath, err := ioutil.TempDir("", "")
	suite.Require().NoError(err)

	ts := e2e.New(suite.T(), false)
	defer ts.Close()

	var cp *termtest.ConsoleProcess

	if runtime.GOOS == "windows" {
		cp = ts.SpawnCmd("powershell.exe", "-nologo", "-noprofile", "-command",
			fmt.Sprintf("Expand-Archive -Path '%s' -DestinationPath '%s'", archivePath, tempPath))
	} else {
		cp = ts.SpawnCmd("tar", "-C", tempPath, "-xf", archivePath)
	}

	cp.ExpectExitCode(0)

	cp = ts.SpawnCmd(filepath.Join(tempPath, platform+exe), "--version")
	cp.Expect(constants.RevisionHashShort)
	cp.ExpectExitCode(0)
}

func TestUpdateGenIntegrationTestSuite(t *testing.T) {
	suite.Run(t, new(UpdateGenIntegrationTestSuite))
}
