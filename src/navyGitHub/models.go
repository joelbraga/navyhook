package navyGitHub
import "github.com/andrepinto/navyhook/_vendor/src/github.com/google/go-github/github"


type TemplateData struct {
	ReleaseBuildInfo ReleaseBuildInfo
	Workspace string
}

type ReleaseBuildInfo struct  {
	ReleaseData ReleaseData `json:"release_data,omitempty"`
	UserData UserData `json:"user_data,omitempty"`
	RepositoryData RepositoryData `json:"repository_data,omitempty"`
}


type ReleaseData struct {
	ID               int             `json:"id,omitempty"`
	TagName         string        	`json:"tag_name,omitempty"`
	Name            string        `json:"name,omitempty"`
	Draft           bool          `json:"draft,omitempty"`
	Prerelease      bool          `json:"prerelease,omitempty"`
	CreatedAt       github.Timestamp     `json:"created_at,omitempty"`
	PublishedAt     github.Timestamp     `json:"published_at,omitempty"`
	URL             string        `json:"url,omitempty"`
	AssetsURL       string        `json:"assets_url,omitempty"`
	UploadURL       string        `json:"upload_url,omitempty"`
	ZipballURL      string        `json:"zipball_url,omitempty"`
	TarballURL      string        `json:"tarball_url,omitempty"`
}

type UserData struct {
	Login		string  `json:"login,omitempty"`
	AvatarUrl	string 	`json:"avatar_url,omitempty"`
}

type RepositoryData struct {
	ID               int             `json:"id,omitempty"`
	Name             string          `json:"name,omitempty"`
	FullName         string          `json:"full_name,omitempty"`
	Description      string          `json:"description,omitempty"`
	CloneURL         string          `json:"clone_url,omitempty"`
	GitURL           string          `json:"git_url,omitempty"`
	Private      	 bool `json:"private"`
	URL              string `json:"url,omitempty"`
}

