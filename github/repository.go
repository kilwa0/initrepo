package github

import (
	"context"
	"github.com/google/go-github/v34/github"
	"log"
	"os"
	"path/filepath"
)

const (
	Separator     = os.PathSeparator
	ListSeparator = os.PathListSeparator
)

func CreateRepo(client *github.Client, org string) *github.Repository {
	ctx := context.Background()

	dir, err := os.Getwd()
	if err != nil {
		log.Printf("[ERROR] %s", err)
	}

	repoPath := filepath.Base(dir)

	repo := &github.Repository{
		Name:    github.String(repoPath),
		Private: github.Bool(false),
	}

	create, _, err := client.Repositories.Create(ctx, org, repo)
	if err != nil {
		log.Printf("[ERROR] %s", err)
	}
	return create
}

func DeleteRepo(client *github.Client, org string) *github.Response {
	ctx := context.Background()

	if org == "" {
		user, _, err := client.Users.Get(ctx, "")
	if err != nil {
		log.Printf("[ERROR] %s", err)
	}
	org = *user.Login

	}

	dir, err := os.Getwd()
	if err != nil {
		log.Printf("[ERROR] %s", err)
	}

	repoPath := filepath.Base(dir)

	del, err := client.Repositories.Delete(ctx, org, repoPath)
	if err != nil {
		log.Printf("[ERROR] %s", err)
	}
	return del
}
