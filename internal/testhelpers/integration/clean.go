package integration

import (
	"log"
	"os"
	"testing"

	"github.com/ActiveState/cli/pkg/cmdlets/auth"
	auth_model "github.com/ActiveState/cli/pkg/platform/api/mono/mono_client/authentication"
	"github.com/ActiveState/cli/pkg/platform/api/mono/mono_client/projects"
	"github.com/ActiveState/cli/pkg/platform/api/mono/mono_client/users"
	"github.com/ActiveState/cli/pkg/platform/api/mono/mono_models"
	"github.com/ActiveState/cli/pkg/platform/authentication"
)

func cleanUser(t *testing.T, username string) error {
	err := os.Setenv("ACTIVESTATE_API_HOST", "platform.activestate.com")
	if err != nil {
		return err
	}
	defer func() {
		os.Unsetenv("ACTIVESTATE_API_HOST")
	}()

	fail := auth.AuthenticateWithCredentials(&mono_models.Credentials{
		Token: os.Getenv("PLATFORM_API_TOKEN"),
	})
	if fail != nil {
		log.Fatalf("Could not authenticate test cleaning user: %v", fail)
		return fail.ToError()
	}

	err = testSuperuser(t, username)
	if err != nil {
		t.Logf("Authenticated user is not a superuser, not running removal operation. Got error: %v", err)
		return nil
	}

	projects, err := getProjects(username)
	if err != nil {
		return err
	}
	for _, proj := range projects {
		err = deleteProject(username, proj.Name)
		if err != nil {
			return err
		}
	}

	return deleteUser(username)
}

func testSuperuser(t *testing.T, username string) error {
	params := auth_model.NewLoginAsParams()
	params.SetUsername(username)
	ok, err := authentication.Get().Client().Authentication.LoginAs(params, authentication.ClientAuth())
	if err != nil {
		t.Logf("Could not login as user, got: %s", ok.Error)
		return err
	}

	return nil
}

func getProjects(org string) ([]*mono_models.Project, error) {
	params := projects.NewListProjectsParams()
	params.SetOrganizationName(org)
	listProjectsOK, err := authentication.Get().Client().Projects.ListProjects(params, authentication.ClientAuth())
	if err != nil {
		return nil, err
	}

	return listProjectsOK.Payload, nil
}

func deleteProject(org, name string) error {
	params := projects.NewDeleteProjectParams()
	params.SetOrganizationName(org)
	params.SetProjectName(name)

	_, err := authentication.Client().Projects.DeleteProject(params, authentication.ClientAuth())
	if err != nil {
		return err
	}

	return nil
}

func deleteUser(name string) error {
	params := users.NewDeleteUserParams()
	params.SetUsername(name)

	_, err := authentication.Client().Users.DeleteUser(params, authentication.ClientAuth())
	if err != nil {
		return err
	}

	return nil
}
