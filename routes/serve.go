package routes

import (
	"net/http"

	"github.com/mulkave/go-init/throw"
)

// Serve wraps the features to be served from the routes,
// implements centralised error handling.
func Serve(handler func(w http.ResponseWriter, r *http.Request)) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if e := recover(); e != nil {
				throw.Handle(e.(throw.Error), w, r)
			}
		}()

		handler(w, r)
	}
}
