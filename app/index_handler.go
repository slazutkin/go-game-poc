package app

import "net/http"

func (a *App) HandleIndex(w http.ResponseWriter, r *http.Request) {
	a.tmpl.ExecuteTemplate(w, "index", nil)
}
