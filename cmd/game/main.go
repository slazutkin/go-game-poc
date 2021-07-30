package main

import (
	"go-game-poc/app"
	"log"
)

func main() {
	a := app.App{}

	if err := a.Start(); err != nil {
		log.Fatalf("Can't start application: %s", err.Error())
	}
}
