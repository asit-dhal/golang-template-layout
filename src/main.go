package main

import (
	"log"
	"net/http"
	"templmanager"
	"encoding/json"
    "os"
)


type UserData struct {
	Name        string
	City        string
	Nationality string
}

type SkillSet struct {
	Language string
	Level    string
}


type SkillSets []*SkillSet

type Configuration struct {
    LayoutPath    string
    IncludePath   string
}

func loadConfiguration(fileName string) {
	file, _ := os.Open(fileName)
	decoder := json.NewDecoder(file)
	configuration := Configuration{}
	err := decoder.Decode(&configuration)
	if err != nil {
  		log.Println("error:", err)	
	}
	log.Println("layout path: ", configuration.LayoutPath)
	log.Println("include path: ", configuration.IncludePath)
	templmanager.SetTemplateConfig(configuration.LayoutPath,  configuration.IncludePath)
}


func index(w http.ResponseWriter, r *http.Request) {
	templmanager.RenderTemplate(w, "index.tmpl", nil)
}

func aboutMe(w http.ResponseWriter, r *http.Request) {
	userData := &UserData{Name: "Asit Dhal", City: "Bhubaneswar", Nationality: "Indian"}
	templmanager.RenderTemplate(w, "aboutme.tmpl", userData)
}

func skillSet(w http.ResponseWriter, r *http.Request) {
	skillSets := SkillSets{&SkillSet{Language: "Golang", Level: "Beginner"},
		&SkillSet{Language: "C++", Level: "Advanced"},
		&SkillSet{Language: "Python", Level: "Advanced"}}
	templmanager.RenderTemplate(w, "skillset.tmpl", skillSets)
}

func main() {
	loadConfiguration("config.json")
	templmanager.LoadTemplates()

	server := http.Server{
		Addr: "127.0.0.1:8080",
	}

	http.HandleFunc("/", index)
	http.HandleFunc("/aboutme", aboutMe)
	http.HandleFunc("/skillset", skillSet)
	log.Println("Listening ...")
	server.ListenAndServe()
}
