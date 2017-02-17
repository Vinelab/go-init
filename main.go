package main

import (
	"net/http"

	"github.com/mulkave/go-init/routes"
)

func main() {
	http.ListenAndServe(":80", routes.Register())
}
