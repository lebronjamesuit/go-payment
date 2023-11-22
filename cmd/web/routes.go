package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func (app *application) routes() http.Handler {

	// 	NewRouter is an implementation of Handler
	mux := chi.NewRouter()
	return mux

}
