/*
	Author: David Amante
	Date: October 7, 2019
	Goal: Create a REST API using Golang
*/

package main

import (
	"encoding/json"
	"os"
	"fmt"
	"io/ioutil"
	"strconv"
	"log"
	"net/http"
	"github.com/gorilla/mux"
)

type Projects struct {
	Projects []Project `json: "Projects"`
}

type Project struct {
	ID int `json:"id"`
	Title string `json:"title"`
	Link string `json:"link"`
	Date string `json:"date"`
	Description string `json:"description"`
	Technologies []string `json:"technologies"`
	Categories []string `json:"categories"`
}

var projects Projects

// Landing page to reassure visitors they are in the right place
func getLandingPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to David Amante's Projects API. Try adding /projects or /projects/{id} to the end of this site's URL.")
}

// Returns a response of all projects as a JSON object
func getProjects(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(projects)
}

// Returns a response of an id-specific project as a JSON object
func getProject(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	targetID, _ := strconv.Atoi(params["id"])

	// return 204: No Content error if project id is invalid
	if targetID < 1 || targetID > len(projects.Projects) {
		w.WriteHeader(204) 
		return
	}

	// search for id-specific project and encode it into a JSON object
	for _, project := range projects.Projects{
		if targetID == project.ID {
			json.NewEncoder(w).Encode(project) 
			return
		}
	}
	json.NewEncoder(w).Encode(&Project{})
}

func main() {
	r := mux.NewRouter()

	port := os.Getenv("PORT")
	if port == "" {
		port = "5000"
		log.Println("INFO: No PORT environment variable detected, defaulting to " + port)
	}

	jsonFile, _ :=  os.Open("Projects.json")

	byteValue, _ := ioutil.ReadAll(jsonFile)

	json.Unmarshal(byteValue, &projects)

	r.HandleFunc("/", getLandingPage)
	r.HandleFunc("/projects", getProjects).Methods("GET")
	r.HandleFunc("/projects/", getProjects).Methods("GET") 
	r.HandleFunc("/projects/{id}", getProject).Methods("GET")

	log.Fatal(http.ListenAndServe(":"+port, r))
}
