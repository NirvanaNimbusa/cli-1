package project_test

import (
	"fmt"
	"runtime"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	yaml "gopkg.in/yaml.v2"

	secretsapi "github.com/ActiveState/cli/pkg/platform/api/secrets"
	"github.com/ActiveState/cli/pkg/platform/authentication"
	"github.com/ActiveState/cli/pkg/project"
	"github.com/ActiveState/cli/pkg/projectfile"
)

func init() {
	secretsClient := secretsapi.NewDefaultClient(authentication.Get().BearerToken())
	project.RegisterExpander("variables", project.NewVarPromptingExpander(secretsClient))
}

func loadProject(t *testing.T) *project.Project {
	projectfile.Reset()

	pjFile := &projectfile.Project{}
	contents := strings.TrimSpace(`
project: "https://platform.activestate.com/Expander/general?commitID=00010001-0001-0001-0001-000100010001"
platforms:
  - name: Linux
    os: linux
  - name: Windows
    os: windows
  - name: macOS
    os: macos
events:
  - name: pre
    value: echo 'Hello $variables.foo!'
  - name: post
    value: echo 'Hello $variables.bar!'
constants:
  - name: constant
    value: value
  - name: recursive
    value: recursive $constants.constant
scripts:
  - name: test
    value: make test
  - name: recursive
    value: $scripts.recursive
`)

	err := yaml.Unmarshal([]byte(contents), pjFile)
	assert.Nil(t, err, "Unmarshalled YAML")

	fail := pjFile.Parse()
	assert.NoError(t, fail.ToError())

	pjFile.Persist()

	return project.Get()
}

func TestExpandProjectPlatformOs(t *testing.T) {
	prj := loadProject(t)

	expanded := project.ExpandFromProject("$platform.os", prj)
	assert.NoError(t, project.Failure().ToError(), "Ran without failure")

	if runtime.GOOS != "darwin" {
		assert.Equal(t, runtime.GOOS, expanded, "Expanded platform variable")
	} else {
		assert.Equal(t, "macos", expanded, "Expanded platform variable")
	}
}

func TestExpandProjectScript(t *testing.T) {
	prj := loadProject(t)

	expanded := project.ExpandFromProject("$ $scripts.test", prj)
	assert.NoError(t, project.Failure().ToError(), "Ran without failure")
	assert.Equal(t, "$ make test", expanded, "Expanded simple script")
}

func TestExpandProjectConstant(t *testing.T) {
	prj := loadProject(t)

	expanded := project.ExpandFromProject("$ $constants.constant", prj)
	assert.NoError(t, project.Failure().ToError(), "Ran without failure")
	assert.Equal(t, "$ value", expanded, "Expanded simple constant")

	expanded = project.ExpandFromProject("$ $constants.recursive", prj)
	assert.NoError(t, project.Failure().ToError(), "Ran without failure")
	assert.Equal(t, "$ recursive value", expanded, "Expanded recursive constant")
}

func TestExpandProjectAlternateSyntax(t *testing.T) {
	prj := loadProject(t)

	expanded := project.ExpandFromProject("${platform.os}", prj)
	assert.NoError(t, project.Failure().ToError(), "Ran without failure")
	if runtime.GOOS != "darwin" {
		assert.Equal(t, runtime.GOOS, expanded, "Expanded platform variable")
	} else {
		assert.Equal(t, "macos", expanded, "Expanded platform variable")
	}
}

func TestExpandProjectUnknownCategory(t *testing.T) {
	prj := loadProject(t)

	expanded := project.ExpandFromProject("$unknown.unknown", prj)
	assert.Error(t, project.Failure().ToError(), "Ran with failure")
	assert.Equal(t, "", expanded, "Failed to expand")
	assert.True(t, project.Failure().Type.Matches(project.FailExpandVariableBadCategory), "Handled unknown category")
}

func TestExpandProjectUnknownName(t *testing.T) {
	prj := loadProject(t)

	expanded := project.ExpandFromProject("$platform.unknown", prj)
	assert.Error(t, project.Failure().ToError(), "Ran with failure")
	assert.Equal(t, "", expanded, "Failed to expand")
	assert.True(t, project.Failure().Type.Matches(project.FailExpandVariableBadName), "Handled unknown category")
}

func TestExpandProjectInfiniteRecursion(t *testing.T) {
	prj := loadProject(t)

	expanded := project.ExpandFromProject("$scripts.recursive", prj)
	assert.Error(t, project.Failure().ToError(), "Ran with failure")
	assert.Equal(t, "", expanded, "Failed to expand")
	assert.True(t, project.Failure().Type.Matches(project.FailExpandVariableRecursion), "Handled unknown category")
}

// Tests all possible $platform.[name] variable expansions.
func TestExpandProjectPlatform(t *testing.T) {
	projectFile := &projectfile.Project{}
	contents := strings.TrimSpace(`
project: https://platform.activestate.com/Expander/Plarforms?commitID=00010001-0001-0001-0001-000100010001"
platforms:
  - name: Any
`)

	err := yaml.Unmarshal([]byte(contents), projectFile)
	assert.Nil(t, err, "Unmarshalled YAML")
	projectFile.Persist()
	prj := project.Get()

	for _, name := range []string{"name", "os", "version", "architecture", "libc", "compiler"} {
		project.ExpandFromProject(fmt.Sprintf("$platform.%s", name), prj)
		assert.NoError(t, project.Failure().ToError(), "Ran without failure")
	}
}

func TestExpandDashed(t *testing.T) {
	projectFile := &projectfile.Project{}
	contents := strings.TrimSpace(`
project: "https://platform.activestate.com/Expander/Dashed?commitID=00010001-0001-0001-0001-000100010001"
scripts:
  - name: foo-bar
    value: bar
`)

	err := yaml.Unmarshal([]byte(contents), projectFile)
	assert.Nil(t, err, "Unmarshalled YAML")
	fail := projectFile.Parse()
	assert.NoError(t, fail.ToError())
	projectFile.Persist()
	prj := project.Get()

	expanded := project.ExpandFromProject("- $scripts.foo-bar -", prj)
	assert.NoError(t, project.Failure().ToError(), "Ran without failure")
	assert.Equal(t, "- bar -", expanded)
	projectfile.Reset()
}
