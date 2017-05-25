package handlers

import (
	"fmt"
	"net/http"
)

// ApproxCircle  is a handler for /approxCircle request.
// Note, this handler assumes it is already decorated to set proper
// HTTP response headers.
func ApproxCircle(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Not supported ..")
	http.Error(w, "Not implemented", 501)
}
