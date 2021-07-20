package main

import (
	"go-game-poc/app"
	"log"
)

func main() {
	a := app.App{}

	if err := a.Start(); err != nil {
		log.Fatal(err.Error())
	}
}
