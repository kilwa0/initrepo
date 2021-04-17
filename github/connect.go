package github

import (
	"context"
	"github.com/google/go-github/v34/github"
	"golang.org/x/oauth2"
	"log"
)

func Connect(token string) *github.Client {
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	tc := oauth2.NewClient(ctx, ts)

	client := github.NewClient(tc)

	return client
}

func ConnectEnterprise(token string, url string) *github.Client {
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	tc := oauth2.NewClient(ctx, ts)

	client, err := github.NewEnterpriseClient(url, url, tc)
	if err != nil {
		log.Printf("[ERROR] %s", err)
	}

	return client
}
