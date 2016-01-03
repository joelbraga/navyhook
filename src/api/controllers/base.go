package api


import (
	"net/http"
	"encoding/json"
	"github.com/andrepinto/navyhook/src/base"
	"github.com/andrepinto/navyhook/src/api/models"
)

func ShowApiVersion(w http.ResponseWriter, r *http.Request)  {

	api := models.Api{
		Name: base.GetConfig().Name,
		Version: base.GetConfig().Version,
	}

	json.NewEncoder(w).Encode(api)
}

func ShowApiRoutes(w http.ResponseWriter, r *http.Request){

}

func ShowConfigRepos(w http.ResponseWriter, r *http.Request){



	//fmt.Println(string(out[:]))

	json.NewEncoder(w).Encode(base.GetAllRepositoriesConfig())


}

