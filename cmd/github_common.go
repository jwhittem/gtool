package cmd

import (
	"context"
	"log"

	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

// NewGithubClient will return a new client object, accepts a token string or nil
func NewGithubClient(t string) *github.Client {
	var client *github.Client
	// use personal access token if defined, avoiding low threshold rate limits on github.com
	if t == "" {
		client = github.NewClient(nil)
	} else {
		ts := oauth2.StaticTokenSource(
			&oauth2.Token{AccessToken: t},
		)
		tc := oauth2.NewClient(context.Background(), ts)
		client = github.NewClient(tc)
	}

	Zen, _, _ := client.Zen(context.Background())

	if Zen == "" {
		return nil
	}
	return client
}

// GetAllRepos will return all repos for a user.
func GetAllRepos(client *github.Client, account string) []*github.Repository {
	opt := &github.RepositoryListOptions{
		ListOptions: github.ListOptions{PerPage: 10},
	}

	//get a list of repos
	var allRepos []*github.Repository
	for {
		repos, resp, err := client.Repositories.List(context.Background(), account, opt)
		if err != nil {
			log.Println(err)
		}

		allRepos = append(allRepos, repos...)
		if resp.NextPage == 0 {
			break
		}
		opt.Page = resp.NextPage
	}

	if allRepos == nil {
		return nil
	}
	return allRepos
}
