package websocket

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func Serve(w http.ResponseWriter, r *http.Request) {

	token := r.URL.Query().Get("Authorization")
	token = token[len("Bearer "):]

	section := chi.URLParam(r, "section")

	if err := Subscribe(section, token, w, r); err != nil {
		http.Error(w, "Could not open websocket connection", http.StatusBadRequest)
	}

}
