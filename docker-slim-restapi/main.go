package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Docker-Slim REST API"))
	})

	r.Route("/commands", func(r chi.Router) {
		r.Route("/{commandName}", func(r chi.Router) {
			r.Get("/executions", CommandHandler)
			r.Post("/executions", CommandHandler)
		})
	})
	http.ListenAndServe(":3333", r)
}

func CommandHandler(w http.ResponseWriter, r *http.Request) {
	commandName := chi.URLParam(r, "commandName")

	command := Command{}
	command.CreatedAt = time.Now().Local()
	command.Name = commandName
	command.Result = "CMD " + commandName + " command result...."

	commandJson, err := json.Marshal(command)
	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	//Write json response back to response
	w.Write(commandJson)
}

type Command struct {
	Name      string
	Result    string
	CreatedAt time.Time
}
