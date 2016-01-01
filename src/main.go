package main

import(
	"fmt"
	"net/http"
	"log"
	"github.com/andrepinto/navyhook/src/rest"
)

func main()  {


	apiRouter := rest.ApiRouter()

	http.Handle("/", apiRouter)

	log.Println("start listening port 4000")

	go http.ListenAndServe(":4000", nil)

	fmt.Scanln()
}
