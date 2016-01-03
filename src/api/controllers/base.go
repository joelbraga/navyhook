package api


import (
	"net/http"
	"encoding/json"
	"github.com/andrepinto/navyhook/_vendor/src/github.com/spf13/viper"
	"github.com/andrepinto/navyhook/src/base"
	"github.com/andrepinto/navyhook/src/api/models"
)

func ShowApiVersion(w http.ResponseWriter, r *http.Request)  {

	api := models.Api{
		Name: viper.GetString("name"),
		Version: viper.GetString("version"),
	}

	json.NewEncoder(w).Encode(api)
}

func ShowApiRoutes(w http.ResponseWriter, r *http.Request){

}

func ShowConfigRepos(w http.ResponseWriter, r *http.Request){



	//fmt.Println(string(out[:]))

	json.NewEncoder(w).Encode(base.GetAllRepositoriesConfig())


}

