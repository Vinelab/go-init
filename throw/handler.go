package throw

import "net/http"

// Handle panics, it prefers to be deferred.
func Handle(e Error, w http.ResponseWriter, r *http.Request) {
	switch e.(type) {
	case InvalidInput:
		err := e.(InvalidInput) // cast to error type
		http.Error(w, err.Message, int(err.Status))
	}
}
