package main

import (
	"html/template"
	"log"
	"math/rand"
	"net/http"
)
//127.0.0.1:7887

type pageData struct {
	Title string
	Number int
}

func indexHandler(writer http.ResponseWriter, request *http.Request){
	t, _ := template.ParseFiles("index.html")
	indexData := pageData{Title:"Start page", Number:0}
	_ = t.Execute(writer, indexData)
}

func aboutHandler(writer http.ResponseWriter, request *http.Request){
	t, _ := template.ParseFiles("about.html")
	aboutData := pageData{Title:"About page", Number:0}
	_ = t.Execute(writer, aboutData)
}

func luckyHandler(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "text/html")
	number := rand.Intn(100)
	t, _ := template.ParseFiles("lucky.html")
	luckyData := pageData{Title:"Lucky Number", Number:number}
	_ = t.Execute(writer, luckyData)
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", indexHandler)
	mux.HandleFunc("/about/", aboutHandler)
	mux.HandleFunc("/lucky/", luckyHandler)
	err := http.ListenAndServe(":7887", mux)
	if err != nil {
		log.Fatal("Server unable to start: %V", err)
	}
}