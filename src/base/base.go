package base
import (
	"github.com/andrepinto/navyhook/src/database"
	"fmt"
)

func GetAllRepositoriesConfig() []database.Repository{
	/*var C database.Repositories
	viper.UnmarshalKey("repositories", &C)
	return C*/
	return database.GetAllRepositories()
}

func GetRepositoryConfig(name string) database.Repository{
	repos := GetAllRepositoriesConfig()
	var rp database.Repository
	for _, repo := range repos{
		if repo.Name == name{
			rp = repo
			break
		}
	}

	return rp
}

func GetRepoHook(name, hook string) database.Hook{

	var hk database.Hook

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
  	fmt.Println(rp)
	hks := rp.Hooks
	for _, hck := range hks{
		if hck.Name == hook{
			rst=true
		}
	}


	return rst
}

