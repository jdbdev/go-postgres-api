package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// STRUCTS

type routeResponse struct {
	Message string `json:"message"`
}

func main() {
	fmt.Println("Simple API with Postgres")

	// http Router:
	router := http.NewServeMux()
	router.HandleFunc("GET /", indexPage)

	// http server:
	err := http.ListenAndServe(":8000", router)
	if err != nil {
		log.Fatal(err)
	}
}

// HANDLER FUNCTIONS:

// IndexPage
func indexPage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	json.NewEncoder(w).Encode(routeResponse{Message: "Hello World"})
	w.Write([]byte("index page"))
}
