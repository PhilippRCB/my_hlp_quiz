package main

import (
	"bytes"
	"encoding/json"
	"html/template"
	"log"
	"net/http"
)

type player struct {
	Name string `json:"name"`
}

var templates = template.Must(template.ParseFiles("buzzer.html"))

const adresse = "localhost:3333/buzzer/"

func buzzerHandler(w http.ResponseWriter, r *http.Request) {
	n := r.FormValue("name")
	p := &player{Name: n}
	body, _ := json.Marshal(p)
	request, _ := http.NewRequest("POST", adresse, bytes.NewBuffer(body))
	request.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	client.Do(request)
	viewBuzzerNameHandler(w, *p)
}

func viewBuzzerHandler(w http.ResponseWriter, r *http.Request) {
	p := &player{Name: "Teamname"}
	viewBuzzerNameHandler(w, *p)
}

func viewBuzzerNameHandler(w http.ResponseWriter, p player) {
	err := templates.ExecuteTemplate(w, "buzzer.html", p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func main() {
	http.HandleFunc("/buzzer/", viewBuzzerHandler)
	http.HandleFunc("/buzz/", buzzerHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
