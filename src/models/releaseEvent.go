package models
import "github.com/andrepinto/navyhook/_vendor/src/github.com/google/go-github/github"

type PullRequestEvent struct {
	Action      *string      `json:"action,omitempty"`
	Release *github.RepositoryRelease `json:"release,omitempty"`
	Repo   *github.Repository `json:"repository,omitempty"`
	Sender *github.User       `json:"sender,omitempty"`
}