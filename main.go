package main

import (
	"net/http"
	"os"

	"go.uber.org/atomic"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	var fail atomic.Bool
	http.ListenAndServe(":"+port, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/start-failing" {
			fail.Store(true)
		}

		if fail.Load() {
			w.WriteHeader(500)
		}
	}))
}
