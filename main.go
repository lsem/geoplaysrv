package main

import (
	"fmt"
	"net/http"

	"github.com/lsem/geosrv/handlers"
)

func apiHandler(fn func(w http.ResponseWriter, r *http.Request)) func(w http.ResponseWriter, r *http.Request) {
	fmt.Println("decorating ...")
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		fn(w, r)
	}
}

func main() {
	fmt.Println("server starting ...")
	http.HandleFunc("/approxRect", apiHandler(handlers.ApproxRect))
	http.HandleFunc("/approxCircle", apiHandler(handlers.ApproxCircle))
	http.ListenAndServe(":8000", nil)
}
