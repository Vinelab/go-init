package routes

import (
	"net/http"

	"github.com/vinelab/go-init/throw"
)

// Serve wraps the features to be served from the routes,
// implements centralised error handling.
func Serve(handler func(w http.ResponseWriter, r *http.Request)) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		defer throw.HandleHttpRequestErrors(w)

		handler(w, r)
	}
}
