package app

import (
	"fmt"
	"go-game-poc/pkg/event"
	"go-game-poc/pkg/templater"
	"go-game-poc/pkg/ws"
	"log"
	"net/http"
	"os"
	"text/template"

	"github.com/gorilla/mux"
)

type App struct {
	tmpl         *template.Template
	pool         *ws.Pool
	stateManager *StateManager
}

func (a *App) Start() error {

	t, err := templater.ParseTemplates(os.Getenv("TEMPLATES_PATH"))

	if err != nil {
		return err
	}

	a.tmpl = t

	inbound := make(chan *event.Event)
	outbound := make(chan *event.Event)

	a.pool = ws.NewPool(inbound, outbound)
	go a.pool.Start()

	a.stateManager = NewStateManager(inbound, outbound)
	go a.stateManager.Start()

	r := mux.NewRouter()
	r.HandleFunc("/", a.HandleIndex).Methods(http.MethodGet)
	r.HandleFunc("/ws", a.HandleWS)

	addr := fmt.Sprintf(":%s", os.Getenv("PORT"))

	log.Printf("Starting server at: %s", addr)

	return http.ListenAndServe(addr, r)
}
