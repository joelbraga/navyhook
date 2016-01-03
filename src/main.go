package main

import(
	"fmt"
	"net/http"
	"log"
	"github.com/andrepinto/navyhook/src/api"
	"github.com/andrepinto/navyhook/_vendor/src/github.com/spf13/viper"
)

func main()  {

	//LOAD CONFIG
	viper.SetConfigName("config")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	port := viper.GetString("port")

	apiRouter := api.ApiRouter()

	http.Handle("/", apiRouter)

	log.Println("start listening port "+port)

	go http.ListenAndServe(":"+port, nil)

	fmt.Scanln()
}
