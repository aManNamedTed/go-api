package main

// TODO: write documentation - 10/7/2019, 8:09PM PST
import (
	"encoding/json"
	"os"
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

func getProjects(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(projects)
}

func getProject(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	targetID, _ := strconv.Atoi(params["id"])

	if targetID < 1 || targetID > len(projects.Projects) {
		w.WriteHeader(204)
		return
	}

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

	r.HandleFunc("/projects", getProjects).Methods("GET")
	r.HandleFunc("/projects/", getProjects).Methods("GET")
	r.HandleFunc("/projects/{id}", getProject).Methods("GET")

	log.Fatal(http.ListenAndServe(":"+port, r))
}
