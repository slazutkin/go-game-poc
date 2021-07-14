package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"text/template"

	"github.com/gorilla/mux"
)

type app struct {
	tmpl *template.Template
}

func (a *app) Start() error {

	t, err := parseTemplates(os.Getenv("TEMPLATES_PATH"))

	if err != nil {
		return err
	}

	a.tmpl = t

	r := mux.NewRouter()
	r.HandleFunc("/", a.HomePage).Methods(http.MethodGet)

	addr := fmt.Sprintf(":%s", os.Getenv("PORT"))

	log.Printf("Starting server at: %s", addr)

	return http.ListenAndServe(addr, r)
}
