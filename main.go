package main

import (
	"fmt"
	"net/http"

	"github.com/lsem/geosrv/handlers"
)

func apiHandler(fn func(w http.ResponseWriter, r *http.Request)) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		fn(w, r)
	}
}

func main() {
	fmt.Println("Server is starting at localhost:8000 ..")
	http.HandleFunc("/approxRect", apiHandler(handlers.ApproxRect))
	http.HandleFunc("/approxCircle", apiHandler(handlers.ApproxCircle))
	http.ListenAndServe(":8000", nil)
}
