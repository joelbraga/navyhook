package navyGitHub

import (
	"github.com/andrepinto/navyhook/_vendor/src/github.com/google/go-github/github"
	"golang.org/x/oauth2"
	"github.com/andrepinto/navyhook/_vendor/src/github.com/spf13/viper"
)

func GetGitHubClient() *github.Client{
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: viper.GetString("gitHubToken")},
	)
	tc := oauth2.NewClient(oauth2.NoContext, ts)

	client := github.NewClient(tc)

	return client
}

