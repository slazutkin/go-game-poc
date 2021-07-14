package main

import "log"

func main() {
	a := app{}

	if err := a.Start(); err != nil {
		log.Fatal(err.Error())
	}
}
