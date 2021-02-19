package server

import (
	"fmt"
	"net/http"
)

// PlayerServer function
func PlayerServer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "20")
}
