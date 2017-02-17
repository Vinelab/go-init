package routes

import v "github.com/husobee/vestigo"

// Register holds the routes to be registered
// when the server starts listening
func Register() *v.Router {
	r := v.NewRouter()

	return r
}
