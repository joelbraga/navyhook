package api
import (
	"net/http"
	"encoding/json"
	"io/ioutil"
	"github.com/andrepinto/navyhook/src/database"
	"github.com/andrepinto/navyhook/_vendor/src/github.com/gorilla/mux"
	"github.com/andrepinto/navyhook/src/api/models"
	"strconv"
	"fmt"
)



func AddRepo(w http.ResponseWriter, r *http.Request){
	body, _ := ioutil.ReadAll(r.Body)
	repo := database.Repository{}
	json.Unmarshal(body, &repo)

	repo.SaveRepository()

	json.NewEncoder(w).Encode(repo)

}

func GetRepo(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	repo := database.GetRepositoryByName(vars["name"])
	repo.Hooks = database.GetHooksByRepo(repo.ID)
	json.NewEncoder(w).Encode(repo)

}

func GetAllRepos(w http.ResponseWriter, r *http.Request){

	json.NewEncoder(w).Encode(database.GetAllRepositories())

}

func DeleteRepo(w http.ResponseWriter, r *http.Request){
	body, _ := ioutil.ReadAll(r.Body)
	vars := mux.Vars(r)

	var cfg = database.Repository{}
	json.Unmarshal(body, &cfg)


	cfg.Name = vars["name"]

	oldConfig := database.GetRepositoryByName(cfg.Name)

	cfg.ID = oldConfig.ID

	ok := cfg.DeleteRepository()

	if ok{
		json.NewEncoder(w).Encode(cfg)
	}else{
		json.NewEncoder(w).Encode(models.ApiError{
			Code: "error",
			Message: "Error on execute",
		})
	}
}

func CreateRepoHook(w http.ResponseWriter, r *http.Request){
	body, _ := ioutil.ReadAll(r.Body)
	vars := mux.Vars(r)

	repo := database.GetRepositoryByName(vars["name"])

	hook := database.Hook{}
	json.Unmarshal(body, &hook)

	hook.RepositoryId = repo.ID

	hook.SaveHook()

	json.NewEncoder(w).Encode(hook)

}


func UpdateRepoHook(w http.ResponseWriter, r *http.Request){
	body, _ := ioutil.ReadAll(r.Body)
	vars := mux.Vars(r)
	fmt.Println(vars)
	hook := database.Hook{}
	json.Unmarshal(body, &hook)


	hook.ID, _ = strconv.Atoi(vars["id"])


	hook.UpdateHook()

	json.NewEncoder(w).Encode(hook)

}


func DeleteRepoHook(w http.ResponseWriter, r *http.Request){
	body, _ := ioutil.ReadAll(r.Body)
	vars := mux.Vars(r)

	repo := database.GetRepositoryByName(vars["name"])

	hook := database.Hook{}
	json.Unmarshal(body, &hook)

	hook.RepositoryId = repo.ID
	hook.ID, _ = strconv.Atoi(vars["id"])

	hook.DeleteHook()

	json.NewEncoder(w).Encode(hook)

}