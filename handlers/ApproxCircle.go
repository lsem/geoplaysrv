package handlers

import (
	"fmt"
	"net/http"
)

// ApproxCircle  is a handler for /approxCircle request.
func ApproxCircle(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Handler is working ..")
}
