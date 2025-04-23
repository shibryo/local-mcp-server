package infra

import (
	"context"
	"fmt"

	"github.com/google/go-github/v71/github"
)

type GitHub struct {
	clinet *github.Client
}

func NewGitHub(token string) *GitHub {
	client := github.NewClient(nil).WithAuthToken(token)

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
