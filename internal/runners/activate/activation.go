package activate

import (
	"fmt"
	"os"
	"os/signal"
	"path/filepath"
	rt "runtime"
	"syscall"

	"github.com/ActiveState/cli/internal/fileevents"
	"github.com/ActiveState/cli/internal/locale"
	"github.com/ActiveState/cli/internal/logging"
	"github.com/ActiveState/cli/internal/output"
	"github.com/ActiveState/cli/internal/virtualenvironment"
	"github.com/ActiveState/cli/pkg/platform/runtime"
	"github.com/ActiveState/cli/pkg/project"
	"github.com/ActiveState/cli/pkg/projectfile"
)

func (r *Activate) activateAndWait(proj *project.Project, runtime *runtime.Runtime) error {
	logging.Debug("Activating and waiting")
	
	err := os.Chdir(filepath.Dir(proj.Source().Path()))
	if err != nil {
		return err
	}

	venv := virtualenvironment.New(runtime)
	venv.OnUseCache(func() { r.out.Notice(locale.T("using_cached_env")) })

	fail := venv.Setup(true)
	if fail != nil {
		return locale.WrapError(fail, "error_could_not_activate_venv", "Could not activate project. If this is a private project ensure that you are authenticated.")
	}

	ve, err := venv.GetEnv(false, filepath.Dir(projectfile.Get().Path()))
	if err != nil {
		return locale.WrapError(err, "error_could_not_activate_venv", "Could not retrieve environment information.")
	}

	// If we're not using plain output then we should just dump the environment information
	if r.out.Type() != output.PlainFormatName {
		if r.out.Type() == output.EditorV0FormatName {
			fmt.Println("[activated-JSON]")
		}
		r.out.Print(ve)
		return nil
	}

	ignoreWindowsInterrupts()

	r.subshell.SetEnv(ve)
	if fail = r.subshell.Activate(r.out); fail != nil {
		return locale.WrapError(err, "error_could_not_activate_subshell", "Could not activate a new subshell.")
	}

	fe, err := fileevents.New(proj)
	if err != nil {
		return locale.WrapError(err, "err_activate_fileevents", "Could not start file event watcher.")
	}
	defer fe.Close()

	fail = <-r.subshell.Failures()
	if fail != nil {
		return locale.WrapError(fail, "error_in_active_subshell", "Failure encountered in active subshell")
	}

	r.out.Notice(locale.T("info_deactivated", proj))

	return nil
}

func ignoreWindowsInterrupts() {
	if rt.GOOS == "windows" {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT)
		go func() {
			for range c {
			}
		}()
	}
}
