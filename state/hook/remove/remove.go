package remove

import (
	"github.com/ActiveState/ActiveState-CLI/internal/failures"
	"github.com/ActiveState/ActiveState-CLI/internal/locale"
	"github.com/ActiveState/ActiveState-CLI/internal/print"
	"github.com/ActiveState/ActiveState-CLI/pkg/cmdlets/commands"
	"github.com/ActiveState/ActiveState-CLI/pkg/cmdlets/hooks"
	"github.com/ActiveState/ActiveState-CLI/pkg/projectfile"
	"gopkg.in/AlecAivazis/survey.v1"

	"github.com/ActiveState/ActiveState-CLI/internal/logging"
	"github.com/spf13/cobra"
)

// Args hold the arg values passed through the command line
var Args struct {
	Identifier string
}

// Used for testing
var testPromptResultOverride string

// Command remove, sub command of hook
var Command = &commands.Command{
	Name:        "remove",
	Description: "hook_remove_description",
	Run:         Execute,

	Arguments: []*commands.Argument{
		&commands.Argument{
			Name:        "arg_hook_remove_identifier",
			Description: "arg_hook_remove_identifier_description",
			Variable:    &Args.Identifier,
		},
	},
}

// Execute the hook remove command
// Adds a statement to be run on the given hook
func Execute(cmd *cobra.Command, args []string) {
	logging.Debug("Execute `hook remove`")

	project, err := projectfile.Get()
	if err != nil {
		failures.Handle(err, locale.T("hook_remove_cannot_remove", Args))
		return
	}

	var removed *projectfile.Hook
	removed = removebyHash(project, Args.Identifier)

	if removed == nil {
		hashedHooks, err := hooks.HashHooksFiltered(project.Hooks, []string{Args.Identifier})
		if err != nil {
			failures.Handle(err, locale.T("hook_remove_cannot_remove"))
			return
		}

		numOfHooksFound := len(hashedHooks)
		if numOfHooksFound == 1 && Args.Identifier != "" {
			removed = removeByName(project)
		} else if numOfHooksFound > 0 {
			removed = removeByPrompt(project) // under construction
		} else {
			failures.Handle(failures.User.New(locale.T("err_hook_cannot_find")), "")
		}
	}

	if removed == nil {
		print.Warning(locale.T("hook_remove_cannot_remove"))
	} else {
		hash, _ := removed.Hash()
		print.Info(locale.T("hook_removed", map[string]interface{}{"Hook": removed, "Hash": hash}))
	}
}

//  Cycle through the configured hooks, hash then remove hook if matches, save, exit
func removebyHash(project *projectfile.Project, hashToRemove string) *projectfile.Hook {
	newHooks := project.Hooks
	var removed *projectfile.Hook
	for i, hook := range newHooks {
		hash, err := hook.Hash()
		if hashToRemove == hash {
			newHooks := append(newHooks[:i], newHooks[i+1:]...)
			project.Hooks = newHooks
			removed = &hook
			break
		} else if err != nil {
			logging.Warning("Failed to remove hook '%v': %v", hashToRemove, err)
			print.Warning(locale.T("hook_remove_cannot_remove", map[string]interface{}{"Hookname": hashToRemove, "Error": err}))
		}
	}
	project.Save()
	return removed
}

func removeByName(project *projectfile.Project) *projectfile.Hook {
	newHooks := project.Hooks
	var removed *projectfile.Hook
	for i, hook := range newHooks {
		if Args.Identifier == hook.Name {
			newHooks := append(newHooks[:i], newHooks[i+1:]...)
			project.Hooks = newHooks
			removed = &hook
			break
		}
	}
	project.Save()
	return removed
}

func removeByPrompt(project *projectfile.Project) *projectfile.Hook {
	var removed *projectfile.Hook

	options, optionsMap, err := hooks.PromptOptions(project, Args.Identifier)
	if err != nil {
		failures.Handle(err, locale.T("err_hook_cannot_list"))
	}

	prompt := &survey.Select{
		Message: locale.T("prompt_hook_choose_remove"),
		Options: options,
	}

	result := ""
	err = survey.AskOne(prompt, &result, nil)

	// For tests we want to override the result as we cannot process prompts from within a test
	if testPromptResultOverride != "" {
		result = testPromptResultOverride
	}

	if err != nil && testPromptResultOverride == "" {
		failures.Handle(err, locale.T("err_invalid_input"))
		return removed
	}

	hash, exists := optionsMap[result]
	print.Formatted("\nresult: %v\n", result)
	print.Formatted("\nmap: %v\n", optionsMap)
	if result == "" || !exists {
		failures.Handle(failures.User.New(locale.T("err_hook_cannot_find")), "")
		return removed
	}

	print.Line()
	return removebyHash(project, hash)
}
