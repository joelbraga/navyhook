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


	apiRouter := api.ApiRouter()

	http.Handle("/", apiRouter)

	port := map[bool]string{true: base.GetConfig().Port, false:"4000"} [len(base.GetConfig().Port)>0]
	log.Println("start listening port "+port)

	go http.ListenAndServe(":"+port, nil)

	fmt.Scanln()
}
