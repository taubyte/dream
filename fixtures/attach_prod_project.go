package fixtures

import (
	"context"
	"errors"
	"fmt"
	"io"
	"strings"

	commonAuth "bitbucket.org/taubyte/auth/common"
	httpAuthClient "bitbucket.org/taubyte/go-auth-http"
	"github.com/google/go-github/github"
	commonDreamland "github.com/taubyte/dreamland/core/common"
	dreamlandRegistry "github.com/taubyte/dreamland/core/registry"
	"github.com/taubyte/dreamland/helpers"
	"golang.org/x/oauth2"
)

func init() {
	dreamlandRegistry.Fixture("attachProdProject", attachProdProject)
}

// Added this variable so that import can call the attachProdProjectFixture, without having
// to rewrite code
var SharedRepositoryData *httpAuthClient.RawRepoDataOuter

func attachProdProject(u commonDreamland.Universe, params ...interface{}) error {
	if len(params) < 2 {
		return errors.New("attachProdProject expects 2 parameters [project-id] [git-token]")
	}

	projectId := params[0].(string)
	if len(projectId) > 0 {
		helpers.ProjectID = projectId
	}

	gitToken := params[1].(string)
	if len(gitToken) > 0 {
		helpers.GitToken = gitToken
	}

	prodAuthURL := "https://auth.taubyte.com"
	prodClient, err := httpAuthClient.New(u.Context(), httpAuthClient.URL(prodAuthURL), httpAuthClient.Auth(gitToken), httpAuthClient.Unsecure(), httpAuthClient.Provider(helpers.GitProvider))
	if err != nil {
		return err
	}

	project, err := prodClient.GetProjectById(projectId)
	if err != nil {
		return err
	}

	SharedRepositoryData, err = project.Repositories()
	if err != nil {
		return err
	}

	// Override auth method so that projectID is not changed
	commonAuth.GetNewProjectID = func(args ...interface{}) string {
		return projectId
	}

	SharedRepositoryData.Configuration.Id, err = GetRepoId(u.Context(), SharedRepositoryData.Configuration.Fullname, gitToken)
	if err != nil {
		return err
	}

	SharedRepositoryData.Code.Id, err = GetRepoId(u.Context(), SharedRepositoryData.Code.Fullname, gitToken)
	if err != nil {
		return err
	}

	devAuthUrl, err := u.GetURLHttp(u.Auth().Node())
	if err != nil {
		return err
	}

	devClient, err := httpAuthClient.New(u.Context(), httpAuthClient.URL(devAuthUrl), httpAuthClient.Auth(gitToken), httpAuthClient.Provider(helpers.GitProvider))
	if err != nil {
		return err
	}

	devClient.RegisterRepository(SharedRepositoryData.Configuration.Id)
	if err != nil {
		return err
	}

	devClient.RegisterRepository(SharedRepositoryData.Code.Id)
	if err != nil {
		return err
	}

	return project.Create(devClient, SharedRepositoryData.Configuration.Id, SharedRepositoryData.Code.Id)
}

func GetRepoId(ctx context.Context, repoFullName string, token string) (string, error) {
	gitClient := newGithubClient(ctx, token)
	if len(repoFullName) == 0 {
		return "", errors.New("repo not found")
	}

	repoFullnameSplit := strings.Split(repoFullName, "/")
	repo, resp, err := gitClient.Repositories.Get(ctx, repoFullnameSplit[0], repoFullnameSplit[1])
	if err != nil {
		body, err0 := io.ReadAll(resp.Body)
		defer resp.Body.Close()
		if err0 != nil {
			return "", fmt.Errorf("calling github api failed with: [%v & %v]", err, err0)
		}
		return "", fmt.Errorf("getting github repo failed with %v, got response %s", err, string(body))
	}

	return fmt.Sprintf("%d", repo.GetID()), nil
}

var gitClient *github.Client

func newGithubClient(ctx context.Context, token string) *github.Client {
	if gitClient != nil {
		return gitClient
	}
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	tc := oauth2.NewClient(ctx, ts)
	gitClient = github.NewClient(tc)
	return gitClient
}
