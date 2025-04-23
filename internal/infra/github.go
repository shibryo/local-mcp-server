package infra

import (
	"context"
	"fmt"

	"github.com/google/go-github/v71/github"
	"github.com/jferrl/go-githubauth"
	"golang.org/x/oauth2"
)

type GitHub struct {
	clinet *github.Client
}

func NewGitHub(token string) *GitHub {
	privateKey := []byte(token)

	appTokenSource, err := githubauth.NewApplicationTokenSource(1112, privateKey)
	if err != nil {
		fmt.Println("Error creating application token source:", err)
		return nil
	}

	installationTokenSource := githubauth.NewInstallationTokenSource(1113, appTokenSource)

	// oauth2.NewClient uses oauth2.ReuseTokenSource to reuse the token until it expires.
	// The token will be automatically refreshed when it expires.
	// InstallationTokenSource has the mechanism to refresh the token when it expires.
	httpClient := oauth2.NewClient(context.Background(), installationTokenSource)

	client := github.NewClient(httpClient)
	return &GitHub{client}
}

func (g *GitHub) CreateIssue(owner, repo, title, body string, tag *[]string) error {
	issueRequest := &github.IssueRequest{
		Title:  github.String(title),
		Body:   github.String(body),
		Labels: tag,
	}

	issue, _, err := g.clinet.Issues.Create(context.Background(), owner, repo, issueRequest)
	if err != nil {
		return err
	}

	fmt.Printf("Issue created: %s\n", *issue.HTMLURL)
	return nil
}
