package main

import (
	"net/http"

	controllers "github.com/thampton/go-tutorial/controllers"
)

func main() {
	controllers.RegisterControllers()
	http.ListenAndServe(":3000", nil)
}
