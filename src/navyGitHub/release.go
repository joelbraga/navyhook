package navyGitHub
import (
	"fmt"
	"os"
	"github.com/andrepinto/navyhook/src/generator"
	"github.com/andrepinto/navyhook/src/base"
	"github.com/andrepinto/navyhook/src/database"
)

const(
	ZIP = ".zip"
	RELEASE = "release"
	PRE_RELEASE = "prerelease"
)


func BuildRelease(releaseBuildInfo ReleaseBuildInfo) error{

	x := database.Action{
		Repository:releaseBuildInfo.RepositoryData.Name,
		Event: map[bool]string{true: PRE_RELEASE, false: RELEASE} [releaseBuildInfo.ReleaseData.Prerelease],
		UserName: releaseBuildInfo.UserData.Login,
		AvatarURL: releaseBuildInfo.UserData.AvatarUrl,
		Info:fmt.Sprintf("%s",releaseBuildInfo.ReleaseData.TagName),
	}

	x.Save()


	repoHookConfig := base.GetRepoHook(releaseBuildInfo.RepositoryData.Name, RELEASE)

	workspace := fmt.Sprintf("%s/%s", base.GetConfig().Workspace, releaseBuildInfo.RepositoryData.Name)
	os.Mkdir(workspace, 0700)

	path := fmt.Sprintf("%s/%d%s", workspace, releaseBuildInfo.ReleaseData.ID,ZIP)

	DownloadDoc(releaseBuildInfo.ReleaseData.ZipballURL, path, base.GetConfig().GitHubToken)

	_, err := os.OpenFile(path,os.O_RDWR, os.FileMode(0666))

	if err != nil {
		x.OnError(err.Error())
		return err
	}

	err, name :=Unzip(path, workspace)

	if err != nil {
		x.OnError(err.Error())
		return err
	}

	buildFolder := fmt.Sprintf("%s/%s", workspace, name)
	prjNavyFolder := fmt.Sprintf("%s%s", buildFolder, base.NAVY_HOOK_FOLDER)

	if _, err := os.Stat(prjNavyFolder); os.IsNotExist(err) {
		x.OnError(err.Error())
		return err
	}

	var tmpl, cmd string

	if releaseBuildInfo.ReleaseData.Prerelease{
		tmpl = fmt.Sprintf("%s/%s", prjNavyFolder, base.PRE_RELEASE_TPL_FILE)
		cmd = fmt.Sprintf("%s/%s", prjNavyFolder, base.PRE_RELEASE_FILE)
	}else{
		tmpl = fmt.Sprintf("%s/%s", prjNavyFolder, base.PRE_RELEASE_TPL_FILE)
		cmd = fmt.Sprintf("%s/%s", prjNavyFolder, base.RELEASE_FILE)
	}


	templateData := TemplateData{
		ReleaseBuildInfo: releaseBuildInfo,
		Workspace: buildFolder,
	}

	err = generator.WriteTemplate(templateData, tmpl, cmd)

	if err != nil {
		x.OnError(err.Error())
		return err
	}

	var result string

	if(repoHookConfig.Exec){
		result, _ = base.RunUnixCommand(cmd)
	}

	x.OnSuccess(result)

	if repoHookConfig.RemoveFolder {
		os.RemoveAll(buildFolder)
		os.RemoveAll(path)
	}

	return nil
}