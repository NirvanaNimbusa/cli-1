package activate

import (
	"os"
	"os/exec"
	"os/signal"
	"path"
	"runtime"
	"syscall"
	"time"

	"github.com/ActiveState/cli/internal/config"
	"github.com/ActiveState/cli/internal/constants"
	"github.com/ActiveState/cli/internal/failures"
	"github.com/ActiveState/cli/internal/hail"
	"github.com/ActiveState/cli/internal/locale"
	"github.com/ActiveState/cli/internal/logging"
	"github.com/ActiveState/cli/internal/print"
	"github.com/ActiveState/cli/internal/subshell"
	"github.com/ActiveState/cli/internal/virtualenvironment"
	"github.com/ActiveState/cli/pkg/project"
)

type activationLoopFunc func(targetPath string, activator activateFunc) error

func activationLoop(targetPath string, activator activateFunc) error {
	// activate should be continually called while returning true
	// looping here provides a layer of scope to handle printing output
	var proj *project.Project
	var retErr error
	for {
		var fail *failures.Failure
		proj, fail = project.FromPath(targetPath)
		if fail != nil {
			// The default failure returned by the project package is a big too vague, we want to give the user
			// something more actionable for the context they're in
			return failures.FailUserInput.New("err_project_notexist_asyaml")
		}
		print.Info(locale.T("info_activating_state", proj))

		err := os.Chdir(targetPath)
		if err != nil {
			return err
		}

		if constants.BranchName != constants.StableBranch {
			print.Stderr().Warning(locale.Tr("unstable_version_warning", constants.BugTrackerURL))
		}

		loop, err := activator(proj.Owner(), proj.Name(), proj.Source().Path())
		if err != nil {
			retErr = err
			break
		}
		if !loop {
			break
		}

		print.Info(locale.T("info_reactivating", proj))
	}

	print.Bold(locale.T("info_deactivated", proj))

	return retErr
}

type activateFunc func(owner, name, srcPath string) (bool, error)

// activate will activate the venv and subshell. It is meant to be run in a loop
// with the return value indicating whether another iteration is warranted.
func activate(owner, name, srcPath string) (bool, error) {
	venv := virtualenvironment.Get()
	venv.OnDownloadArtifacts(func() { print.Line(locale.T("downloading_artifacts")) })
	venv.OnInstallArtifacts(func() { print.Line(locale.T("installing_artifacts")) })
	fail := venv.Activate()
	if fail != nil {
		failures.Handle(fail, locale.T("error_could_not_activate_venv"))
		return false, nil
	}

	ignoreWindowsInterrupts()

	subs, err := subshell.Activate()
	if err != nil {
		failures.Handle(err, locale.T("error_could_not_activate_subshell"))
		return false, nil
	}

	done := make(chan struct{})
	defer close(done)
	fname := path.Join(config.ConfigPath(), constants.UpdateHailFileName)

	hails, fail := hail.Open(done, fname)
	if fail != nil {
		failures.Handle(fail, locale.T("error_unable_to_monitor_pulls"))
		return false, nil
	}

	return listenForReactivation(venv.ActivationID(), hails, subs)
}

func ignoreWindowsInterrupts() {
	if runtime.GOOS == "windows" {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT)
		go func() {
			for range c {
			}
		}()
	}
}

type subShell interface {
	Deactivate() *failures.Failure
	Failures() <-chan *failures.Failure
}

func listenForReactivation(id string, rcvs <-chan *hail.Received, subs subShell) (bool, error) {
	for {
		select {
		case rcvd, ok := <-rcvs:
			if !ok {
				logging.Error("hailing channel closed")
				return false, nil
			}

			if rcvd.Fail != nil {
				logging.Error("error in hailing channel: %s", rcvd.Fail)
				continue
			}

			if !idsValid(id, rcvd.Data) {
				continue
			}

			// A subshell will have triggered this case; Wait for
			// output completion before deactivating. The nature of
			// this issue is unclear at this time.
			time.Sleep(time.Second)

			if fail := subs.Deactivate(); fail != nil {
				failures.Handle(fail, locale.T("error_deactivating_subshell"))
				return false, nil
			}

			return true, nil

		case fail, ok := <-subs.Failures():
			if !ok {
				logging.Error("subshell failure channel closed")
				return false, nil
			}

			if fail != nil {
				err := fail.ToError()
				if eerr, ok := err.(*exec.ExitError); ok {
					return false, eerr
				}

				fail.Silent = true
				failures.Handle(fail, locale.T("error_in_active_subshell"))
				//failures.Handle(fail, "")
			}

			return false, nil
		}
	}
}

func idsValid(currID string, rcvdID []byte) bool {
	return currID != "" && len(rcvdID) > 0 && currID == string(rcvdID)
}
