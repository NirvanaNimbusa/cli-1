package variables

import (
	"fmt"
	"strings"

	"github.com/bndr/gotabulate"
	"github.com/spf13/cobra"

	"github.com/ActiveState/cli/internal/failures"
	"github.com/ActiveState/cli/internal/locale"
	"github.com/ActiveState/cli/internal/logging"
	"github.com/ActiveState/cli/internal/print"
	"github.com/ActiveState/cli/pkg/cmdlets/commands"
	secretsapi "github.com/ActiveState/cli/pkg/platform/api/secrets"
	"github.com/ActiveState/cli/pkg/project"
)

// Command represents the secrets command and its dependencies.
type Command struct {
	config        *commands.Command
	secretsClient *secretsapi.Client

	Args struct {
		Name            string
		Value           string
		ShareUserHandle string
	}
}

// NewCommand creates a new Keypair command.
func NewCommand(secretsClient *secretsapi.Client) *Command {
	c := Command{
		secretsClient: secretsClient,
		config: &commands.Command{
			Name:        "variables",
			Aliases:     []string{"vars"},
			Description: "variables_cmd_description",
		},
	}
	c.config.Run = c.Execute

	c.config.Append(buildGetCommand(&c))
	c.config.Append(buildSetCommand(&c))
	c.config.Append(buildSyncCommand(&c))

	return &c
}

// Config returns the underlying commands.Command definition.
func (cmd *Command) Config() *commands.Command {
	return cmd.config
}

// Execute processes the secrets command.
func (cmd *Command) Execute(_ *cobra.Command, args []string) {
	failure := listAllVariables(cmd.secretsClient)
	if failure != nil {
		failures.Handle(failure, locale.T("variables_err"))
	}
}

// listAllVariables prints a list of all of the variables defined for this project.
func listAllVariables(secretsClient *secretsapi.Client) *failures.Failure {
	prj := project.Get()
	logging.Debug("listing variables for org=%s, project=%s", prj.Owner(), prj.Name())

	rows := [][]interface{}{}
	vars := prj.Variables()
	for _, v := range vars {
		value := ""
		encrypted := "-"
		store := "local"
		shared := "-"
		valOrNil, failure := v.ValueOrNil()
		if failure != nil {
			return failure
		} else if v.IsSecret() {
			if valOrNil == nil {
				value = locale.T("variables_value_secret_undefined")
			} else {
				value = locale.T("variables_value_secret")
			}
			encrypted = locale.T("confirmation")
			if v.IsShared() {
				shared = string(*v.SharedWith())
			}
			store = string(*v.Store())
		} else {
			value = *valOrNil
		}
		rows = append(rows, []interface{}{v.Name(), sanitizeValue(value), encrypted, shared, store})
	}

	t := gotabulate.Create(rows)
	t.SetHeaders([]string{
		locale.T("variables_col_name"),
		locale.T("variables_col_value"),
		locale.T("variables_col_encrypted"),
		locale.T("variables_col_shared"),
		locale.T("variables_col_store"),
	})
	t.SetAlign("left")

	print.Line(t.Render("simple"))

	return nil

}

// sanitizeValue will reduce the string length to 100 characters or the first line of text
func sanitizeValue(v string) string {
	v = strings.TrimSpace(v)
	nlPos := strings.Index(v, "\n")

	if nlPos != -1 && nlPos < 100 {
		v = fmt.Sprintf("%s [...]", v[0:nlPos])
	} else if len(v) > 100 {
		v = fmt.Sprintf("%s [...]", v[0:100])
	}

	return v
}
