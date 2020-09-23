package activate

import (
	"path/filepath"
	"strconv"
	"strings"

	"github.com/ActiveState/cli/internal/constants"
	"github.com/ActiveState/cli/internal/fileutils"
	"github.com/ActiveState/cli/internal/output"
	"github.com/ActiveState/cli/pkg/cmdlets/git"
	"github.com/ActiveState/cli/pkg/platform/model"
	"github.com/ActiveState/cli/pkg/project"
	"github.com/ActiveState/cli/pkg/projectfile"
	"github.com/blang/semver"
)

type CheckoutAble interface {
	Run(namespace string, path string) error
}

// Checkout will checkout the given platform project at the given path
// This includes cloning an associatd repository and creating the activestate.yaml
// It does not activate any environment
type Checkout struct {
	repo git.Repository
	output.Outputer
}

func NewCheckout(repo git.Repository, prime primeable) *Checkout {
	return &Checkout{repo, prime.Output()}
}

func (r *Checkout) Run(namespace string, targetPath string) error {
	ns, fail := project.ParseNamespace(namespace)
	if fail != nil {
		return fail
	}

	pj, fail := model.FetchProjectByName(ns.Owner, ns.Project)
	if fail != nil {
		return fail
	}

	commitID := ns.CommitID
	if commitID == nil {
		branch, fail := model.DefaultBranchForProject(pj)
		if fail != nil {
			return fail
		}
		commitID = branch.CommitID
	}

	// Clone the related repo, if it is defined
	if pj.RepoURL != nil {
		fail = r.repo.CloneProject(ns.Owner, ns.Project, targetPath, r.Outputer)
		if fail != nil {
			return fail
		}
	}

	language, err := getLanguage(ns.Owner, ns.Project)
	if err != nil {
		return err
	}

	// Create the config file, if the repo clone didn't already create it
	configFile := filepath.Join(targetPath, constants.ConfigFileName)
	if !fileutils.FileExists(configFile) {
		fail = projectfile.Create(&projectfile.CreateParams{
			Owner:     ns.Owner,
			Project:   ns.Project,
			CommitID:  commitID,
			Directory: targetPath,
			Language:  language,
		})
		if fail != nil {
			return fail
		}
	}

	return nil
}

func getLanguage(owner, project string) (string, error) {
	modelLanguage, fail := model.DefaultLanguageForProject(owner, project)
	if fail != nil {
		return "", fail
	}

	language := modelLanguage.Name
	if strings.ToLower(modelLanguage.Name) == "python" {
		version, err := semver.Parse(modelLanguage.Version)
		if err != nil {
			return language, err
		}
		language = modelLanguage.Name + strconv.FormatUint(version.Major, 10)
	}

	return language, nil
}
