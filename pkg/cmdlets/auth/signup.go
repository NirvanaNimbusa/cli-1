package auth

import (
	"errors"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"

	"github.com/ActiveState/cli/internal/config"
	"github.com/ActiveState/cli/internal/constants"
	"github.com/ActiveState/cli/internal/errs"
	"github.com/ActiveState/cli/internal/output"
	"github.com/ActiveState/cli/pkg/platform/api"

	"github.com/ActiveState/cli/internal/locale"
	"github.com/ActiveState/cli/internal/logging"
	"github.com/ActiveState/cli/internal/prompt"
	"github.com/ActiveState/cli/pkg/platform/api/mono"
	"github.com/ActiveState/cli/pkg/platform/api/mono/mono_client/users"
	"github.com/ActiveState/cli/pkg/platform/api/mono/mono_models"
	"github.com/ActiveState/cli/pkg/platform/authentication"
)

type signupInput struct {
	Name      string
	Email     string
	Username  string
	Password  string
	Password2 string
}

// Signup will prompt the user to create an account
func Signup(out output.Outputer, prompt prompt.Prompter) error {
	input := &signupInput{}

	if authentication.Get().Authenticated() {
		return locale.NewInputError("err_auth_authenticated", "You are already authenticated as: {{.V0}}. You can log out by running `state auth logout`.", authentication.Get().WhoAmI())
	}

	accepted, err := promptTOS(out, prompt)
	if err != nil {
		return err
	}
	if !accepted {
		return locale.NewInputError("tos_not_accepted", "")
	}

	err = promptForSignup(input, out, prompt)
	if err != nil {
		return locale.WrapError(err, "err_prompt_unknown")
	}

	if err = doSignup(input, out); err != nil {
		return err
	}

	if authentication.Get().Authenticated() {
		if err := generateKeypairForUser(input.Password); err != nil {
			return locale.WrapError(err, "keypair_err_save")
		}
	}

	return nil
}

func signupFromLogin(username string, password string, out output.Outputer, prompt prompt.Prompter) error {
	input := &signupInput{}

	input.Username = username
	input.Password = password
	err := promptForSignup(input, out, prompt)
	if err != nil {
		return errs.Wrap(err, "UserInput failure")
	}

	return doSignup(input, out)
}

func downloadTOS() (string, error) {
	resp, err := http.Get(constants.TermsOfServiceURLText)
	if err != nil {
		return "", errs.Wrap(err, "Failed to download the Terms Of Service document.")
	}
	if resp.StatusCode != http.StatusOK {
		return "", errs.New("The server responded with status '%s' when trying to download the Terms Of Service document.", resp.Status)
	}
	defer resp.Body.Close()

	tosPath := filepath.Join(config.ConfigPath(), "platform_tos.txt")
	tosFile, err := os.Create(tosPath)
	if err != nil {
		return "", errs.Wrap(err, "Could not create Terms Of Service file in configuration directory.")
	}
	defer tosFile.Close()

	_, err = io.Copy(tosFile, resp.Body)
	if err != nil {
		return "", errs.Wrap(err, "Failed to write Terms Of Service file contents.")
	}

	return tosPath, nil
}

func promptTOS(out output.Outputer, prompt prompt.Prompter) (bool, error) {
	choices := []string{
		locale.T("tos_accept"),
		locale.T("tos_not_accept"),
		locale.T("tos_show_full"),
	}

	out.Notice(locale.Tr("tos_disclaimer", constants.TermsOfServiceURLLatest))
	choice, err := prompt.Select(locale.Tl("tos", "Terms of Service"), locale.T("tos_acceptance"), choices, locale.T("tos_accept"))
	if err != nil {
		return false, err
	}

	switch choice {
	case locale.T("tos_accept"):
		return true, nil
	case locale.T("tos_not_accept"):
		return false, nil
	case locale.T("tos_show_full"):
		tosFilePath, err := downloadTOS()
		if err != nil {
			return false, locale.WrapError(err, "err_download_tos", "Could not download terms of service file.")
		}

		tos, err := ioutil.ReadFile(tosFilePath)
		if err != nil {
			return false, errs.Wrap(err, "IO failure")
		}
		out.Print(tos)
		return prompt.Confirm("", locale.T("tos_acceptance"), true)
	}

	return false, nil
}

func promptForSignup(input *signupInput, out output.Outputer, prompter prompt.Prompter) error {
	var err error

	if input.Username != "" {
		out.Notice(locale.T("confirm_password_account_creation"))
	} else {
		input.Username, err = prompter.Input("", locale.T("username_prompt_signup"), "", prompt.InputRequired)
		if err != nil {
			return err
		}
		input.Password, err = prompter.InputSecret("", locale.T("password_prompt_signup"), prompt.InputRequired)
		if err != nil {
			return err
		}
	}

	// Must define password validator here as it has to reference the input
	var passwordValidator = func(val interface{}) error {
		value := val.(string)
		if value != input.Password {
			return locale.NewError("InvalidPassword")
		}
		return nil
	}

	input.Password2, err = prompter.InputSecret("", locale.T("password_prompt_confirm"), prompt.InputRequired)
	if err != nil {
		return err
	}
	err = passwordValidator(input.Password2)
	if err != nil {
		return errs.Wrap(err, "InvalidPassword failure")
	}

	input.Name, err = prompter.Input("", locale.T("name_prompt"), "", prompt.InputRequired)
	if err != nil {
		return err
	}

	input.Email, err = prompter.Input("", locale.T("email_prompt"), "", prompt.InputRequired)
	if err != nil {
		return err
	}
	return nil
}

func doSignup(input *signupInput, out output.Outputer) error {
	params := users.NewAddUserParams()
	eulaHelper := true
	params.SetUser(&mono_models.UserEditable{
		Email:        input.Email,
		Username:     input.Username,
		Password:     input.Password,
		Name:         input.Name,
		EULAAccepted: &eulaHelper,
	})
	addUserOK, err := mono.Get().Users.AddUser(params)

	// Error checking
	if err != nil {
		switch err.(type) {
		// Authentication failed due to email already existing (username check already happened at this point)
		case *users.AddUserConflict:
			logging.Error("Encountered add user conflict: %v", err)
			return locale.WrapInputError(err, "err_auth_signup_user_exists", "", api.ErrorMessageFromPayload(err))
		default:
			logging.Error("Encountered unknown error adding user: %v", err)
			return locale.WrapError(err, "err_auth_failed_unknown_cause", "", api.ErrorMessageFromPayload(err))
		}
	}

	err = AuthenticateWithCredentials(&mono_models.Credentials{
		Username: input.Username,
		Password: input.Password,
	})
	if err != nil {
		return err
	}

	out.Notice(locale.T("signup_success", map[string]string{
		"Email": addUserOK.Payload.User.Email,
	}))

	return nil
}

// UsernameValidator verifies whether the chosen username is usable
func UsernameValidator(val interface{}) error {
	value := val.(string)
	params := users.NewUniqueUsernameParams()
	params.SetUsername(value)
	res, err := mono.Get().Users.UniqueUsername(params)
	if err != nil || *res.Payload.Code != int64(200) {
		return errors.New(locale.T("err_username_taken"))
	}
	return nil
}
