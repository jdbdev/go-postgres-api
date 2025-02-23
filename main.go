package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
)

// STRUCTS

type routeResponse struct {
	Message string `json:"message"`
}

func main() {
	log.Println("starting server...")

	// http Router:
	log.Println("setting up routes...")
	router := http.NewServeMux()
	router.HandleFunc("GET /", indexPage)
	router.HandleFunc("POST /users/register", register)
	router.HandleFunc("POST /users/login", login)
	router.HandleFunc("POST /projects/create", createProject)
	router.HandleFunc("POST /projects/update", updateProject)

	// http server:
	log.Println("listening on port 8000...")
	err := http.ListenAndServe(":8000", router)
	if err != nil {
		log.Fatal(err)
	}
}

// MIDDLEWARE:

// HANDLER FUNCTIONS:

// IndexPage
func indexPage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	json.NewEncoder(w).Encode(routeResponse{Message: "Hello World"})
	w.Write([]byte("index page"))

	// Output request Header to console (testing only)
	fmt.Println(r.Header)

	// Better formatting using DumpRequest()
	requestDump, err := httputil.DumpRequest(r, true)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("REQUEST: %s\n", string(requestDump))
	_, err = w.Write(requestDump)
	if err != nil {
		return
	}
}

// User - Register (POST):
func register(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	json.NewEncoder(w).Encode(routeResponse{Message: "User Registration Page"})
}

// User - Login  (POST):
func login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	json.NewEncoder(w).Encode(routeResponse{Message: "Login Page"})
}

// Project - Create (POST):
func createProject(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	json.NewEncoder(w).Encode(routeResponse{Message: "Create Project Page"})
}

// Project - Update (PUT):
func updateProject(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	json.NewEncoder(w).Encode(routeResponse{Message: "Update Project Page"})
}

// Project - Get all projects:
// Project - Get project instance:
// Project - Delete project:
