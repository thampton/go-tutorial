package main

import (
	"net/http"

	controllers "github.com/thampton/go-tutorial/controllers"
)

/*
 HTTPRequest - Struct for HTTP Request:w
*/
type HTTPRequest struct {
	Method string
}

func main() {
	controllers.RegisterControllers()
	http.ListenAndServe(":3000", nil)
}
