package app

import (
	"fmt"
	"go-game-poc/pkg/ws"
	"log"
	"net/http"

	gonanoid "github.com/matoous/go-nanoid/v2"
)

func (a *App) HandleWS(w http.ResponseWriter, r *http.Request) {
	log.Println("Connected")
	conn, err := ws.Upgrade(w, r)
	if err != nil {
		fmt.Fprintf(w, "%+v\n", err)
	}

	id, err := gonanoid.New()
	if err != nil {
		return
	}

	client := &ws.Client{
		ID:   id,
		Conn: conn,
		Pool: a.pool,
	}

	a.pool.Register <- client
	client.Read()
}
