package main

import (
	"net/http"

	controllers "github.com/thampton/go-tutorial/Controllers"
)

func main() {
	controllers.RegisterControllers()
	http.ListenAndServe(":3000", nil)
}
