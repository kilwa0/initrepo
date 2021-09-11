package github

import (
	"context"
	"github.com/google/go-github/v34/github"
	"log"
)

func ListUserRepos(client *github.Client, user string, perpage int) []string {
	ctx := context.Background()
	opts := &github.RepositoryListOptions{
		ListOptions: github.ListOptions{PerPage: perpage, Page: 1},
	}

	var allReposName []string
	var allRepos []*github.Repository
	for {
		repos, resp, err := client.Repositories.List(ctx, user, opts)
		if err != nil {
			log.Printf("[ERROR] %s", err)
		}
		allRepos = append(allRepos, repos...)
		opts.Page = resp.NextPage

		if resp.NextPage == 0 {
			break
		}
	}
	for index := range allRepos {
		user, _, err := client.Users.Get(ctx, "")
		if err != nil {
			log.Printf("[ERROR] %s", err)
		}
		if allRepos[index].GetFullName() != user.GetLogin()+"/"+user.GetLogin() {
			allReposName = append(allReposName, allRepos[index].GetFullName())
		}
	}
	return allReposName
}

func GetUser(client *github.Client, user string) *github.User {
	ctx := context.Background()
	// var username string
	userGet, _, err := client.Users.Get(ctx, user)
	if err != nil {
		log.Printf("[ERROR] %s", err)
	}
	return userGet
}
