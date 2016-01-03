package api
import (
	"net/http"
	"encoding/json"
	"github.com/andrepinto/navyhook/_vendor/src/github.com/gorilla/mux"
	"github.com/andrepinto/navyhook/src/database"
)

func GetRepoActions(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	action := database.Action{
		Repository: vars["repo"],
	}

	json.NewEncoder(w).Encode(action.GetByRepository())

}

func GetAllRepoActions(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	action := database.Action{
		Repository: vars["repo"],
	}

	json.NewEncoder(w).Encode(action.GetAll())


}

