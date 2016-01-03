package api

import(
	"net/http"
	"encoding/json"
	"github.com/andrepinto/navyhook/src/database"
	"io/ioutil"
	"github.com/andrepinto/navyhook/_vendor/src/github.com/gorilla/mux"
	"github.com/andrepinto/navyhook/src/api/models"
)

func CreateConfiguration(w http.ResponseWriter, r *http.Request){
	body, _ := ioutil.ReadAll(r.Body)

	var cfg = database.Configuration{}

	json.Unmarshal(body, &cfg)

	database.DB.Delete(&cfg)
	ok := cfg.SaveConfiguration()

	if ok{
		json.NewEncoder(w).Encode(cfg)
	}else{
		json.NewEncoder(w).Encode(models.ApiError{
			Code: "error",
			Message: "Error on execute",
		})
	}

}


func UpdateConfiguration(w http.ResponseWriter, r *http.Request){
	body, _ := ioutil.ReadAll(r.Body)
	vars := mux.Vars(r)

	var cfg = database.Configuration{}
	json.Unmarshal(body, &cfg)

	cfg.Name = vars["name"]

	oldConfig := database.GetConfigurationByName(cfg.Name)

	cfg.ID = oldConfig.ID

	ok := cfg.UpdateConfiguration()

	if ok{
		json.NewEncoder(w).Encode(cfg)
	}else{
		json.NewEncoder(w).Encode(models.ApiError{
			Code: "error",
			Message: "Error on execute",
		})
	}



}

func DeleteConfiguration(w http.ResponseWriter, r *http.Request){
	body, _ := ioutil.ReadAll(r.Body)
	vars := mux.Vars(r)

	var cfg = database.Configuration{}
	json.Unmarshal(body, &cfg)


	cfg.Name = vars["name"]

	oldConfig := database.GetConfigurationByName(cfg.Name)

	cfg.ID = oldConfig.ID

	ok := cfg.DeleteConfiguration()

	if ok{
		json.NewEncoder(w).Encode(cfg)
	}else{
		json.NewEncoder(w).Encode(models.ApiError{
			Code: "error",
			Message: "Error on execute",
		})
	}


}

