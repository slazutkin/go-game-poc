package app

import (
	"fmt"
	"go-game-poc/pkg/ws"
	"log"
	"net/http"
)

func (a *App) HandleWS(w http.ResponseWriter, r *http.Request) {
	log.Println("Connected")
	conn, err := ws.Upgrade(w, r)
	if err != nil {
		fmt.Fprintf(w, "%+v\n", err)
	}

	client := &ws.Client{
		Conn: conn,
		Pool: a.pool,
	}

	a.pool.Register <- client
	client.Read()
}
