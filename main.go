package main

import (
	"os"
	"fmt"
    "net/http"
    "github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	port := os.Getenv("HTTP_PORT")
	if port == "" {
		port = "8000"
	}

	router.HandleFunc("/hello", hello).Methods("GET")

	err := http.ListenAndServe(":" + port, router)
	if err != nil {
		fmt.Print(err)
	}
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!\n")	
}