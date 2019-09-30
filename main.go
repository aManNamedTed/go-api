package main

import (
	"github.com/gorilla/mux"
	//"go-api/app"
	"go-api/controllers"
	"os"
	"fmt"
	"net/http"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/projects", controllers.GetProjectsFor).Methods("GET")

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	fmt.Println(port)
	err := http.ListenAndServe(":" + port, router)
	if err != nil {
		fmt.Print(err)
	}
}