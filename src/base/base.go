package base
import (
	"github.com/andrepinto/navyhook/src/base/models"
	"github.com/andrepinto/navyhook/_vendor/src/github.com/spf13/viper"

)

func GetAllRepositoriesConfig() basemodels.Repositories{
	var C basemodels.Repositories
	viper.UnmarshalKey("repositories", &C)
	return C
}

func GetRepositoryConfig(name string) basemodels.CRepo{
	repos := GetAllRepositoriesConfig()
	var rp basemodels.CRepo
	for _, repo := range repos{
		if repo.Name == name{
			rp = repo
			break
		}
	}

	return rp
}

func GetRepoHook(name, hook string) basemodels.Hook{

	var hk basemodels.Hook

	rp := GetRepositoryConfig(name)

		hks := rp.Hooks
		for _, hck := range hks{
			if hck.Name == hook{
				hk = hck
				break
			}
		}


	return hk
}

func CheckRepoHook(name, hook string) bool{

	rst := false

	rp := GetRepositoryConfig(name)

	hks := rp.Hooks
	for _, hck := range hks{
		if hck.Name == hook{
			rst=true
		}
	}


	return rst
}

