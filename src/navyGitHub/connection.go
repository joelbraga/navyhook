package navyGitHub

import (
	"github.com/andrepinto/navyhook/_vendor/src/github.com/google/go-github/github"
	"golang.org/x/oauth2"
	"github.com/andrepinto/navyhook/src/base"
)

func GetGitHubClient() *github.Client{
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: base.GetConfig().GitHubToken},
	)
	tc := oauth2.NewClient(oauth2.NoContext, ts)

	client := github.NewClient(tc)

	return client
}

