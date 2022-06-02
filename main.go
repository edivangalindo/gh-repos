package main

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"time"

	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

func main() {

	token := os.Getenv("GH_AUTH_TOKEN")

	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)

	tc := oauth2.NewClient(ctx, ts)
	client := github.NewClient(tc)

	stat, _ := os.Stdin.Stat()

	if (stat.Mode() & os.ModeCharDevice) != 0 {
		fmt.Fprintln(os.Stderr, "No users detected. Hint: cat users.txt | gh-repos")
		os.Exit(1)
	}

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {

		user := scanner.Text()

		// Get repos for each user
		opt := &github.RepositoryListOptions{
			ListOptions: github.ListOptions{PerPage: 100},
		}
		for {

			repos, resp, err := client.Repositories.List(ctx, user, opt)

			if resp.StatusCode == 404 {
				fmt.Println("User not found")
				break
			}

			if err != nil {
				fmt.Println(err)
				if resp.Remaining == 0 {
					fmt.Printf("Rate limit exceeded, waitting for %+v minutes\n", resp.Rate.Reset.Sub(time.Now()).Minutes())
					time.Sleep(resp.Rate.Reset.Sub(time.Now()))
				}
				continue
			}

			// Print repos
			for _, repo := range repos {
				if *repo.Fork {
					continue
				}

				fmt.Println(user + "/" + repo.GetName())
			}

			opt.Page = resp.NextPage

			if resp.NextPage == 0 {
				break
			}
		}
	}
}
