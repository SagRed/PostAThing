package main

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/go-redis/redis/v8"
	"github.com/gorilla/mux"
)

var templates *template.Template
var client *redis.Client

func main() {
	client = redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
	templates = template.Must(template.ParseGlob("templates/*.html"))

	r := mux.NewRouter()
	r.HandleFunc("/", home).Methods("GET")
	http.Handle("/", r)
	http.ListenAndServe(":8000", nil)
}

func home(w http.ResponseWriter, r *http.Request) {
	data, err := client.Get(client.Context(), "data").Result()
	if err != nil {
		return
	}
	fmt.Println("Getting data...")
	templates.ExecuteTemplate(w, "index.html", data)
}
