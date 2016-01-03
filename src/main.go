package main

import(
	"fmt"
	"net/http"
	"log"
	"github.com/andrepinto/navyhook/src/api"
	//"github.com/andrepinto/navyhook/src/database"
	"github.com/andrepinto/navyhook/src/base"
)


func init(){

}

func main()  {

	port := base.GetConfig().Port

	apiRouter := api.ApiRouter()

	http.Handle("/", apiRouter)

	log.Println("start listening port "+port)

	go http.ListenAndServe(":"+port, nil)

	fmt.Scanln()
}
