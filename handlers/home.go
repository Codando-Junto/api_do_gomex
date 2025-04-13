// filepath: /Users/rgomes/code/pessoal/mentoria/api_do_gomex/handlers/home.go
package handlers

import (
	"fmt"
	"net/http"
)

// HomeHandler handles requests to the homepage
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome to the homepage!")
}
