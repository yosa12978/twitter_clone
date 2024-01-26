package handlers

import (
	"context"
	"net/http"

	"github.com/gorilla/mux"
)

func NewRouter(ctx context.Context) http.Handler {
	mux := mux.NewRouter()
	mux.StrictSlash(true)

	mux.Use(recoveryMiddleware)
	mux.Use(loggingMiddleware)
	mux.Use(corsMiddleware)

	user := NewUser(ctx)

	mux.HandleFunc("/", user.GetUsers).Methods("GET")
	mux.HandleFunc("/{id}", user.GetUserById).Methods("GET")
	mux.HandleFunc("/", user.Signup).Methods("POST")

	return mux
}
