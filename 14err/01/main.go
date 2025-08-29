package main

import (
	"fmt"
	"log"
	"net/http"
)
func main() {
	mux := http.NewServeMux()
	mux.Handle("/", ErrorHandler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Simulate a panic
		panic("Something went wrong!")
	})))
	log.Fatal(http.ListenAndServe(":8080", mux))
}

func ErrorHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				log.Printf("Panic: %v", err)
				http.Error(w, fmt.Sprintf("Internal Server Error: %v", err), http.StatusInternalServerError)
			}
		}()
		next.ServeHTTP(w, r)
	})
}