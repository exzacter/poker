package main

import (
	"net/http"
	"html/template"
	"log"
)

func home(w http.ResponseWriter, r *http.Request) {

	files := []string{
		"./frontend/base/base.tmpl",
		"./frontend/base/nav.tmpl",
		"./frontend/home/main.tmpl",
	}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	err = ts.ExecuteTemplate(w, "base", nil)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}

}

func login(w http.ResponseWriter, r *http.Request) {

	files := []string {
		"./frontend/base/base.tmpl",
		"./frontend/base/nav.tmpl",
		"./frontend/login/main.tmpl",
	}
	
	ts, err := template.ParseFiles(files...)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	err = ts.ExecuteTemplate(w, "base", nil)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

func register(w http.ResponseWriter, r *http.Request) {

	files := []string{
		"./frontend/base/base.tmpl",
		"./frontend/base/nav.tmpl",
		"./frontend/register/main.tmpl",
	}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	err = ts.ExecuteTemplate(w, "base", nil)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}

}

