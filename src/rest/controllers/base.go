package controllers


import (
	"fmt"
	"net/http"
)

func ShowApiVersion(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprint(w, "1.0.0\n")
}