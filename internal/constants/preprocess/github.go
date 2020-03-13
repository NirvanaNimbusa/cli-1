package preprocess

import (
	"context"
	"errors"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/ActiveState/cli/internal/constants"
	"github.com/google/go-github/v29/github"
	"golang.org/x/oauth2"
)

const labelPrefix = "version: "
const branchPrefix = "ActiveState:"
const masterBranch = "master"

// GithubIncrementProvider provides methods for getting label values from the Github API
type GithubIncrementProvider struct {
	client *github.Client
}

// NewGithubProvider returns an initialized Github client
func NewGithubProvider(token string) *GithubIncrementProvider {
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	tc := oauth2.NewClient(context.Background(), ts)

	return &GithubIncrementProvider{
		client: github.NewClient(tc),
	}
}

// IncrementBranch returns the increment value string (major, minor, patch) of a
// pull request label for the current pull request
func (g *GithubIncrementProvider) IncrementBranch() (string, error) {
	prNum, err := pullRequestNumber()
	if err != nil {
		return "", err
	}
	if prNum == 0 {
		return patch, nil
	}

	return g.versionLabelPullRequest(prNum)
}

// IncrementMaster returns the version number for the master branch by reading
// the appropriate version file associated with the most recently merged pull request
func (g *GithubIncrementProvider) IncrementMaster() (string, error) {
	pullRequests, err := g.pullRequestList(&github.PullRequestListOptions{
		State:     "closed",
		Sort:      "updated",
		Direction: "desc",
	})
	if err != nil {
		return "", err
	}

	var branchName string
	for _, pullRequest := range pullRequests {
		merged, err := g.isMerged(pullRequest)
		if err != nil {
			return "", err
		}
		if !merged {
			continue
		}
		branchName = strings.TrimPrefix(pullRequest.Head.GetLabel(), branchPrefix)
		break
	}
	if branchName == "" {
		return "", errors.New("could not determine branch name from previosly merged pull requests")
	}

	return getVersionString(branchName)
}

func (g *GithubIncrementProvider) versionLabelPullRequest(number int) (string, error) {
	pullRequest, err := g.pullRequest(number)
	if err != nil {
		return "", err
	}

	label := getLabel(pullRequest.Labels)
	target := strings.TrimPrefix(pullRequest.GetBase().GetLabel(), fmt.Sprintf("%s:", constants.LibraryOwner))
	if target != masterBranch && label == "" {
		return patch, nil
	}

	if label == "" {
		return "", errors.New("no pull request label found")
	}

	return strings.TrimPrefix(label, labelPrefix), nil
}

func (g *GithubIncrementProvider) pullRequestList(options *github.PullRequestListOptions) ([]*github.PullRequest, error) {
	pullReqests, _, err := g.client.PullRequests.List(
		context.Background(),
		constants.LibraryOwner,
		constants.LibraryName,
		options,
	)
	if err != nil {
		return nil, err
	}

	return pullReqests, nil
}

func (g *GithubIncrementProvider) pullRequest(number int) (*github.PullRequest, error) {
	pullRequest, _, err := g.client.PullRequests.Get(context.Background(), constants.LibraryOwner, constants.LibraryName, number)
	if err != nil {
		return nil, err
	}

	return pullRequest, nil
}

func (g *GithubIncrementProvider) isMerged(pullRequest *github.PullRequest) (bool, error) {
	if pullRequest.Number == nil {
		return false, errors.New("could not check if pull request has been merged, invalid pull request received")
	}
	merged, _, err := g.client.PullRequests.IsMerged(
		context.Background(),
		constants.LibraryOwner,
		constants.LibraryName,
		*pullRequest.Number,
	)
	if err != nil {
		return false, err
	}

	return merged, nil
}

func pullRequestNumber() (int, error) {
	// CircleCI
	prInfo := os.Getenv("CI_PULL_REQUEST")
	if prInfo != "" {
		return pullRequestNumberCircle(prInfo)
	}

	// Azure
	prInfo = os.Getenv("SYSTEM_PULLREQUEST_PULLREQUESTNUMBER")
	if prInfo != "" {
		return pullRequestNumberAzure(prInfo)
	}

	// Pull request info not set, we are on a branch but no PR has been created
	// Should still be allowed to build in this state hence we do not return
	// an error here.
	return 0, nil
}

func pullRequestNumberCircle(info string) (int, error) {
	regex := regexp.MustCompile("/pull/[0-9]+")
	match := regex.FindString(info)
	if match == "" {
		return 0, fmt.Errorf("could not determine pull request number from: %s", info)
	}
	num := strings.TrimPrefix(match, "/pull/")
	prNumber, err := strconv.Atoi(num)
	if err != nil {
		return 0, err
	}

	return prNumber, nil
}

func pullRequestNumberAzure(info string) (int, error) {
	regex := regexp.MustCompile("[0-9]+")
	if !regex.MatchString(info) {
		return 0, fmt.Errorf("pull request number contains non-digits, recieved: %s", info)
	}

	prNumber, err := strconv.Atoi(info)
	if err != nil {
		return 0, err
	}

	return prNumber, nil
}

func getLabel(labels []*github.Label) string {
	regex := regexp.MustCompile("version: (major|minor|patch)")

	for _, label := range labels {
		if label.Name != nil && regex.MatchString(*label.Name) {
			return *label.Name
		}
	}

	return ""
}