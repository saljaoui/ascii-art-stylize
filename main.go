package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	ascii "ascii-art/funcs"
)

var (
	tmpl *template.Template
	Port = "7970"
)

type FormData struct {
	ErrorMsg string
}



func main() {
	fmt.Println("Go app...")
	fmt.Println("http://localhost:" + Port)

	var err error
	tmpl, err = template.ParseGlob("templates/*.html")
	if err != nil {
		fmt.Printf("Error parsing templates: %v \n", err)
		return
	}

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.HandleFunc("/", index)
	http.HandleFunc("/ascii-art", processor)
	log.Fatal(http.ListenAndServe(":"+Port, nil))
}

func index(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Bad Request: Only GET method is allowed", http.StatusMethodNotAllowed)
		return
	}

	if r.URL.Path != "/" && r.URL.Path != "ascii-art" {
		w.WriteHeader(http.StatusNotFound)
		tmpl.ExecuteTemplate(w, "error.html", nil)

		return
	}

	tmpl.ExecuteTemplate(w, "index.html", nil)
}

func processor(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Bad Request: Only POST method is allowed", http.StatusMethodNotAllowed)
		return
	}
	var data FormData
	input := r.PostFormValue("inputTextName")

	if !ascii.ScanInput(input) {
		data.ErrorMsg = "Input contains invalid characters"
		w.WriteHeader(http.StatusBadRequest)
		tmpl.ExecuteTemplate(w, "index.html", data)
		return
	}

	if input == "" {
		data.ErrorMsg = "Input Is Empty"
		w.WriteHeader(http.StatusBadRequest)
		tmpl.ExecuteTemplate(w, "index.html", data)
		return
	}

	banners := r.PostFormValue("Banners")
	if !ascii.IsBanners(banners) {
		http.Error(w, "Bad Request: 400 Something went wrong (╯°□°)╯!", http.StatusBadRequest)
		return
	}

	output, err := ascii.GenerateAsciiArt(input, banners)
	if err != nil {
		http.Error(w, "Internal Server Error: 500 Something went wrong (╯°□°)╯!", http.StatusInternalServerError)
		return
	}

	err = tmpl.ExecuteTemplate(w, "ascii-art.html", output)
	if err != nil {
		http.Error(w, "Internal Server Error: 500 Something went wrong (╯°□°)╯! ", http.StatusInternalServerError)
	}
}
