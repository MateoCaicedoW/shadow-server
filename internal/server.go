package internal

import (
	"context"
	"fmt"
	"net/http"
	"sync"

	"github.com/go-chi/chi/v5"
)

// Instance of the server, it contains a router and its basic options,
// this instance is used to apply options to the server.
type Instance struct {
	*chi.Mux

	name string
	host string
	port string
}

func (r *Instance) StartServer() error {
	host := fmt.Sprintf("%v:%v", r.host, r.port)
	fmt.Printf("[info] Starting %v server on port %v\n", r.name, host)

	return http.ListenAndServe(host, r)
}

// New sets up and returns a new HTTP server with routes mounted
// for each of the different features in this application. It also
// sets up the default middleware for the server.
func NewServer(name string) *Instance {
	r := &Instance{
		Mux:  chi.NewRouter(),
		name: name,
		host: "0.0.0.0", //default host
		port: "3000",    //default port
	}

	if Environment == "development" {
		r.host = "127.0.0.1"
	}

	r.Use(setValuer)

	return r
}

type valuer struct {
	data map[string]any
	moot sync.Mutex
}

func setValuer(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		vlr := &valuer{
			data: map[string]any{
				// Adding base values that are useful for the handlers.
				"request":    r,
				"currentURL": r.URL.String(),
			},
		}

		r = r.WithContext(context.WithValue(r.Context(), "valuer", vlr))
		next.ServeHTTP(w, r)
	})
}
