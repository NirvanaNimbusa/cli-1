package cmd

import (
	"os"
	"os/exec"
	"path/filepath"

	"github.com/spf13/viper"

	"github.com/ActiveState/cli/internal/fileutils"
	"github.com/ActiveState/cli/internal/osutils"
	"github.com/ActiveState/cli/internal/output"
	"github.com/ActiveState/cli/internal/subshell/sscommon"
	"github.com/ActiveState/cli/pkg/project"
)

var escaper *osutils.ShellEscape

func init() {
	escaper = osutils.NewBatchEscaper()
}

// SubShell covers the subshell.SubShell interface, reference that for documentation
type SubShell struct {
	binary          string
	rcFile          *os.File
	cmd             *exec.Cmd
	env             map[string]string
	errs            chan error
	activateCommand *string
}

// Shell - see subshell.SubShell
func (v *SubShell) Shell() string {
	return "cmd"
}

// Binary - see subshell.SubShell
func (v *SubShell) Binary() string {
	return v.binary
}

// SetBinary - see subshell.SubShell
func (v *SubShell) SetBinary(binary string) {
	v.binary = binary
}

// WriteUserEnv - see subshell.SubShell
func (v *SubShell) WriteUserEnv(env map[string]string, envType sscommon.EnvData, userScope bool) error {
	cmdEnv := NewCmdEnv(userScope)

	// Clean up old entries
	oldEnv := viper.GetStringMap(envType.Key)
	for k, v := range oldEnv {
		if err := cmdEnv.unset(k, v.(string)); err != nil {
			return err
		}
	}

	// Store new entries
	viper.Set(envType.Key, env)

	for k, v := range env {
		value := v
		if k == "PATH" {
			path, err := cmdEnv.Get("PATH")
			if err != nil {
				return err
			}
			if path != "" {
				path = ";" + path
			}

			value = v + path
		}

		// Set key/value in the user environment
		err := cmdEnv.set(k, value)
		if err != nil {
			return err
		}
	}

	osutils.PropagateEnv()
	return nil
}

// SetupShellRcFile - subshell.SubShell
func (v *SubShell) SetupShellRcFile(targetDir string, env map[string]string, namespace project.Namespaced) error {
	env = sscommon.EscapeEnv(env)
	return sscommon.SetupShellRcFile(filepath.Join(targetDir, "shell.bat"), "config_global.bat", env, namespace)
}

// SetEnv - see subshell.SetEnv
func (v *SubShell) SetEnv(env map[string]string) {
	v.env = env
}

// SetActivateCommand - see subshell.SetActivateCommand
func (v *SubShell) SetActivateCommand(cmd string) {
	v.activateCommand = &cmd
}

// Quote - see subshell.Quote
func (v *SubShell) Quote(value string) string {
	return escaper.Quote(value)
}

// Activate - see subshell.SubShell
func (v *SubShell) Activate(out output.Outputer) error {
	env := sscommon.EscapeEnv(v.env)
	var err error
	if v.rcFile, err = sscommon.SetupProjectRcFile("config.bat", ".bat", env, out); err != nil {
		return err
	}

	shellArgs := []string{"/K", v.rcFile.Name()}
	if v.activateCommand != nil {
		if err := fileutils.AppendToFile(v.rcFile.Name(), []byte("\r\n"+*v.activateCommand+"\r\nexit")); err != nil {
			return err
		}
	}

	cmd := exec.Command("cmd", shellArgs...)

	v.errs = sscommon.Start(cmd)
	v.cmd = cmd
	return nil
}

// Errors returns a channel for receiving errors related to active behavior
func (v *SubShell) Errors() <-chan error {
	return v.errs
}

// Deactivate - see subshell.SubShell
func (v *SubShell) Deactivate() error {
	if !v.IsActive() {
		return nil
	}

	if err := sscommon.Stop(v.cmd); err != nil {
		return err
	}

	v.cmd = nil
	return nil
}

// Run - see subshell.SubShell
func (v *SubShell) Run(filename string, args ...string) error {
	return sscommon.RunFuncByBinary(v.Binary())(osutils.EnvMapToSlice(v.env), filename, args...)
}

// IsActive - see subshell.SubShell
func (v *SubShell) IsActive() bool {
	return v.cmd != nil && (v.cmd.ProcessState == nil || !v.cmd.ProcessState.Exited())
}
