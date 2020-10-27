package befaft

import (
	"github.com/ActiveState/cli/internal/captain"
	"github.com/ActiveState/cli/internal/runners/run"
	"github.com/ActiveState/cli/pkg/project"
)

type BefAft struct {
	Project *project.Project
}

func New(p *project.Project) *BefAft {
	return &BefAft{
		Project: p,
	}
}

func (ba *BefAft) Wrap(next captain.ExecuteFunc) captain.ExecuteFunc {
	return func(cmd *captain.Command, args []string) error {
		runEvent := run.NewEvent(ba.Project.Events())

		if err := runEvent.Run(args, project.BeforeCmd); err != nil {
			return err // TODO: ctx
		}

		if err := next(cmd, args); err != nil {
			return err // TODO: ctx
		}

		if err := runEvent.Run(args, project.AfterCmd); err != nil {
			return err // TODO: ctx
		}

		return nil
	}
}
