package main

import (
	"fmt"
	"http/router/routes"
	"net/http"
)

func main() {

	//make router
	Router := routes.HandleRoutes()

	fmt.Println("startuing server on localhost:8080")
	//set up server
	http.ListenAndServe(":8080", Router)
}
