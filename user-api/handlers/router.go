package handlers

import (
	"net/http"

	"github.com/gorilla/mux"
)

func NewRouter() http.Handler {
	mux := mux.NewRouter()
	mux.StrictSlash(true)
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Wish you were here"))
	})
	return mux
}
