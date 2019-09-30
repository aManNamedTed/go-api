package controllers

import (
	"net/http"
	"go-api/models"
	u "go-api/utils"
)

var GetProjectsFor = func(w http.ResponseWriter, r *http.Request) {
	data := models.GetProjects()
	resp := u.Message(true, "success")
	resp["data"] = data
	u.Respond(w, resp)
}