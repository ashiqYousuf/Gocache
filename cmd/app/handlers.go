package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/ashiqYousuf/Gocache/internal/cache"
)

func (app *application) handleGetValue(w http.ResponseWriter, r *http.Request) {
	key := r.URL.Query().Get("key")
	value, exists := app.cache.Get(key)
	if !exists {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(cache.RenderKeyValue{Key: key, Value: value})
}

func (app *application) handleSetValue(w http.ResponseWriter, r *http.Request) {
	key := r.URL.Query().Get("key")
	value, err := strconv.Atoi(r.URL.Query().Get("value"))
	if err != nil {
		log.Fatal(err)
	}

	app.cache.Set(key, value)
	w.Write([]byte("Value set\n"))
}

func (app *application) handleDelete(w http.ResponseWriter, r *http.Request) {
	key := r.URL.Query().Get("key")
	app.cache.Remove(key)

	w.Write([]byte("Cache key deleted"))
}

func (app *application) handleGetAndDelete(w http.ResponseWriter, r *http.Request) {
	key := r.URL.Query().Get("key")
	value, exists := app.cache.Pop(key)

	if !exists {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(cache.RenderKeyValue{Key: key, Value: value})
}
