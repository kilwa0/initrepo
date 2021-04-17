package main

import (
	"flag"
	"fmt"
	meta "github.com/google/go-github/v34/github"
	"github.com/kilwa0/initrepo/github"
)

func main() {
	var user, token, action, org, host string
	flag.StringVar(&token, "token", "", "Github secret token")
	flag.StringVar(&user, "user", "", "Github User")
	flag.StringVar(&org, "organization", "", "Gihub Organization, if empty defaults to user")
	flag.StringVar(&action, "action", "repositories", "usage")
	flag.StringVar(&host, "host", "", "Github Host")
	flag.Parse()

	if host == "" {
		client := github.Connect(token)
		if action == "repositories" {
			var repositories []string
			repositories = github.ListUserRepos(client, user, 10)
			fmt.Println(repositories)
		} else if action == "create" {
			var repositories *meta.Repository
			repositories = github.CreateRepo(client, org)
			fmt.Println(repositories.GetFullName())
		} else if action == "delete" {
			var repositories *meta.Response
			repositories = github.DeleteRepo(client, org)
			if repositories.StatusCode == 204 {
				fmt.Printf("Deleted\n")
			}
		}
	} else {
		client := github.ConnectEnterprise(token, "https://"+host)
		if action == "repositories" {
			var repositories []string
			repositories = github.ListUserRepos(client, user, 10)
			fmt.Println(repositories)
		} else if action == "create" {
			var repositories *meta.Repository
			repositories = github.CreateRepo(client, org)
			fmt.Println(repositories.GetFullName())
		} else if action == "delete" {
			var repositories *meta.Response
			repositories = github.DeleteRepo(client, org)
			if repositories.StatusCode == 204 {
				fmt.Printf("Deleted\n")
			}
		}
	}
}
