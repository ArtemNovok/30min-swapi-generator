package main

import (
	"bytes"
	"embed"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"text/template"
)

type Char struct {
	Name   string `json:"name"`
	Height string `json:"height,omitempty"`
	Mass   string `json:"mass,omitempty"`
	Gender string `json:"gender,omitempty"`
}

//go:embed templates/*
var tempalteFS embed.FS

func (app *Config) HandleGet(w http.ResponseWriter, r *http.Request) {
	templ := template.Must(template.ParseFS(tempalteFS, "templates/main.html.gohtml"))
	w.WriteHeader(http.StatusOK)
	templ.ExecuteTemplate(w, "index", nil)
}

func (app *Config) HandlePost(w http.ResponseWriter, r *http.Request) {
	id := rand.Intn(50)
	url := fmt.Sprintf("https://swapi.dev/api/people/%v/", id)
	requset, err := http.NewRequest("GET", url, bytes.NewBuffer(nil))
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Unable to prepare request"))
	}

	client := http.Client{}

	response, err := client.Do(requset)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Unable to reach swapi.dev"))
	}
	defer response.Body.Close()

	var char Char

	err = json.NewDecoder(response.Body).Decode(&char)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Failed to decode response"))
	}
	templ := template.Must(template.ParseFS(tempalteFS, "templates/main.html.gohtml"))
	w.WriteHeader(http.StatusOK)
	templ.ExecuteTemplate(w, "char", char)

}
