package main

import (
	"log"
	"net/http"

	"github.com/ashiqYousuf/Gocache/internal/cache"
)

type application struct {
	cache cache.Cacher
}

func main() {
	app := &application{
		cache: cache.NewTTLCache[any, any](),
	}

	http.HandleFunc("/get", app.handleGetValue)
	http.HandleFunc("/set", app.handleSetValue)
	http.HandleFunc("/delete", app.handleDelete)
	http.HandleFunc("/pop", app.handleGetAndDelete)

	log.Println("Server listening on port 8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
