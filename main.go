package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	if err := http.ListenAndServe(fmt.Sprintf(":%s", os.Getenv("PORT")), http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello"))
	})); err != nil {
		log.Fatal(err)
	}
}
