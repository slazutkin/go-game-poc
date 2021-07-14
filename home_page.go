package main

import "net/http"

func (a *app) HomePage(w http.ResponseWriter, r *http.Request) {
	a.tmpl.ExecuteTemplate(w, "index", nil)
}
