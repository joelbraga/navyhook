package navyGitHub
import (
	"fmt"
	"github.com/andrepinto/navyhook/_vendor/src/github.com/spf13/viper"
	"os"
	"github.com/andrepinto/navyhook/src/generator"
	"github.com/andrepinto/navyhook/src/base"
	"errors"
)

const(
	ZIP = ".zip"
	RELEASE = "release"
)


func BuildRelease(releaseBuildInfo ReleaseBuildInfo) error{

	repoHookConfig := base.GetRepoHook(releaseBuildInfo.RepositoryData.Name, RELEASE)

	workspace := fmt.Sprintf("%s/%s", viper.GetString("workspace"), releaseBuildInfo.RepositoryData.Name)
	os.Mkdir(workspace, 0700)

	path := fmt.Sprintf("%s/%d%s", workspace, releaseBuildInfo.ReleaseData.ID,ZIP)

	DownloadDoc(releaseBuildInfo.ReleaseData.ZipballURL, path)

	_, err := os.OpenFile(path,os.O_RDWR, os.FileMode(0666))

	if err != nil {
		panic(err)
	}

	_, name :=Unzip(path, workspace)


	prjNavyFolder := fmt.Sprintf("%s/%s%s", workspace, name, base.NAVY_HOOK_FOLDER)

	if _, err := os.Stat(prjNavyFolder); os.IsNotExist(err) {
		return errors.New("navyhook: navyhook folder not exist")
	}

	var tmpl, cmd string

	if releaseBuildInfo.ReleaseData.Prerelease{
		tmpl = fmt.Sprintf("%s/%s", prjNavyFolder, base.PRE_RELEASE_TPL_FILE)
		cmd = fmt.Sprintf("%s/%s", prjNavyFolder, base.PRE_RELEASE_FILE)
	}else{
		tmpl = fmt.Sprintf("%s/%s", prjNavyFolder, base.PRE_RELEASE_TPL_FILE)
		cmd = fmt.Sprintf("%s/%s", prjNavyFolder, base.RELEASE_FILE)
	}


	generator.WriteTemplate(releaseBuildInfo, tmpl, cmd)

	result, _ := base.RunUnixCommand(cmd)

	fmt.Println(result)

	if repoHookConfig.RemoveFolder {
		os.RemoveAll(fmt.Sprintf("%s/%s", workspace, name))
		os.RemoveAll(path)
	}

	return nil
}