package models

import (
	"github.com/jinzhu/gorm"
	"fmt"
)

type Project struct {
	gorm.Model
	ProjectId uint `json:"id"`
	Title string `json:"title"`
	Link string `json:"link"`
	Date string `json:"date"`
	Description string `json:"description"`
	Technologies []string `json:"technologies"`
	Categories []string `json:"categories"`
}

func GetProject(id uint) (*Project) {
	project := &Project{}
	err := GetDB().Table("projects").Where("id = ?", id).First(project).Error
	if err != nil {
		return nil
	}

	return project
}

func GetProjects() ([]*Project) {
	projects := make([]*Project, 0)
	err := GetDB().Table("projects").Find(&projects).Error
	if err != nil {
		fmt.Println(err)
		return nil
	}

	return projects
}