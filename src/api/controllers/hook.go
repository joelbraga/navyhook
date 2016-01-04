package api
import (
	"net/http"
	"io/ioutil"
	"encoding/json"
	"github.com/andrepinto/navyhook/src/navyGitHub"
	"github.com/andrepinto/navyhook/src/base"
	"github.com/andrepinto/navyhook/src/base/models"
	"fmt"
)

const(
	RELEASE = "release"
	GIT_HEADER_EVENT = "X-Github-Event"
	SUCCESS = "ok"
)

func HookReceive(w http.ResponseWriter, r *http.Request){

	evName := r.Header.Get(GIT_HEADER_EVENT)

	body, _ := ioutil.ReadAll(r.Body)


	if(evName == RELEASE){
		PrepareRelease(body)
	}

	w.Write([]byte(SUCCESS))
}

func PrepareRelease(body []byte){
	rls := models.PullRequestEvent{}
	json.Unmarshal(body, &rls)

	fmt.Println(rls)

	println(base.CheckRepoHook(*rls.Repo.Name, RELEASE))

	if base.CheckRepoHook(*rls.Repo.Name, RELEASE) {
		releaseData := navyGitHub.ReleaseData{
			ID		   : *rls.Release.ID,
			TagName    : *rls.Release.TagName,
			Name       : *rls.Release.Name,
			Draft      : *rls.Release.Draft,
			Prerelease : *rls.Release.Prerelease,
			CreatedAt  : *rls.Release.CreatedAt,
			PublishedAt: *rls.Release.PublishedAt,
			URL        : *rls.Release.URL,
			AssetsURL  : *rls.Release.AssetsURL,
			UploadURL  : *rls.Release.UploadURL,
			ZipballURL : *rls.Release.ZipballURL,
			TarballURL : *rls.Release.TarballURL,
		}

		userData := navyGitHub.UserData{
			Login: *rls.Sender.Login,
			AvatarUrl: *rls.Sender.AvatarURL,
		}

		repositoryData := navyGitHub.RepositoryData{
			ID          : *rls.Repo.ID,
			Name        : *rls.Repo.Name,
			FullName    : *rls.Repo.FullName   ,
			Description : *rls.Repo.Description,
			CloneURL    : *rls.Repo.CloneURL   ,
			GitURL      : *rls.Repo.GitURL     ,
			Private     : *rls.Repo.Private    ,
			URL         : *rls.Repo.URL        ,
		}

		releaseBuildInfo := navyGitHub.ReleaseBuildInfo{
			ReleaseData: releaseData,
			UserData: userData,
			RepositoryData: repositoryData,
		}

		go navyGitHub.BuildRelease(releaseBuildInfo)
	}


}