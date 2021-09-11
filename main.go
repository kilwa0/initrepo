package main

import (
	"context"
	"flag"
	"fmt"
	meta "github.com/google/go-github/v34/github"
	"github.com/kilwa0/initrepo/github"
	"log"
)

func main() {
	var user, token, action, org, host string
	flag.StringVar(&token, "token", "", "Github secret token")
	flag.StringVar(&user, "user", "", "Github User")
	flag.StringVar(&org, "organization", "", "Gihub Organization, if empty defaults to user")
	flag.StringVar(&action, "action", "repositories", `usage:
[create, delete, repositories]`)
	flag.StringVar(&host, "host", "", "Github Host")
	flag.Parse()
	ctx := context.Background()

	if host == "" {
		client := github.Connect(token)
		usr, _, err := client.Users.Get(ctx, "")
		if err != nil {
			log.Printf("[ERROR] %s", err)
		}
		if action == "repositories" {
			var repositories []string
			repositories = github.ListUserRepos(client, user, 10)
			for index, reponame := range repositories {
				fmt.Println(index+1, reponame)
			}
		} else if action == "create" {
			var repositories *meta.Repository
			repositories = github.CreateRepo(client, org)
			fmt.Printf(`
git init
git add -A
git commit -m "first commit"
git branch -M main
git remote add origin git@github.com:%s/%s.git
git push -u origin main
`, *usr.Login, repositories.GetName())
		} else if action == "delete" {
			var repositories *meta.Response
			repositories = github.DeleteRepo(client, org)
			if repositories.StatusCode == 204 {
				fmt.Printf("Deleted\n")
			}
		}
	} else {
		client := github.ConnectEnterprise(token, "https://"+host)
		usr, _, err := client.Users.Get(ctx, "")
		if err != nil {
			log.Printf("[ERROR] %s", err)
		}
		if action == "repositories" {
			var repositories []string
			repositories = github.ListUserRepos(client, user, 10)
			fmt.Println(repositories)
		} else if action == "create" {
			var repositories *meta.Repository
			repositories = github.CreateRepo(client, org)
			fmt.Printf(`
git init
git add -A
git commit -m "first commit"
git branch -M main
git remote add origin git@%s:%s/%s.git
git push -u origin main
`, host, *usr.Login, repositories.GetName())
		} else if action == "delete" {
			var repositories *meta.Response
			repositories = github.DeleteRepo(client, org)
			if repositories.StatusCode == 204 {
				fmt.Printf("Deleted\n")
			}
		}
	}
}
