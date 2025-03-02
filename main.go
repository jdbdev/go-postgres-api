package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"

	"github.com/gorilla/mux"
	"github.com/justinas/alice"
)

// STRUCTS

type routeResponse struct {
	Message string `json:"message"`
	ID string `json:"id,omitempty"`
}

func main() {
	log.Println("starting server...")
	// Middleware Chain:
	alice.New(loggingMiddleware())
	// http Router:
	log.Println("setting up routes...")
	router := mux.NewRouter()
	router.HandleFunc("/", indexPage).Methods("GET")
	router.HandleFunc("/users/register", register).Methods("POST")
	router.HandleFunc("/users/login", login).Methods("POST")
	router.HandleFunc("/projects", getProjects).Methods("GET")
	router.HandleFunc("/projects/{id}", getProject).Methods("GET")
	router.HandleFunc("/projects", createProject).Methods("POST")
	router.HandleFunc("/projects/{id}", updateProject).Methods("PUT")
	router.HandleFunc("/projects/{id}", deleteProject).Methods("DELETE")

	// http server:
	log.Println("listening on port 8000...")
	err := http.ListenAndServe(":8000", router)
	if err != nil {
		log.Fatal(err)
	}
}

// MIDDLEWARE (chaining using Alice package):

func loggingMiddleware(next http.Handler) http.Handler{
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
		// Parse values from Request: 
		log.Printf("%s %s %s\n", r.RemoteAddr, r.Method, r.URL)
		next.ServeHTTP(w,r)
	})
}

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

// Project - Get all projects (GET):
func getProjects(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	json.NewEncoder(w).Encode(routeResponse{Message: "Get all projects page"})
}

// Project - Get project instance (GET):
func getProject(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	json.NewEncoder(w).Encode(routeResponse{Message: "Get project page", ID: id})
}

// Project - Update (PUT):
func updateProject(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	json.NewEncoder(w).Encode(routeResponse{Message: "Update Project Page", ID:id})
}

// Project - Delete project (DELETE):
func deleteProject(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	json.NewEncoder(w).Encode(routeResponse{Message: "Delete project page", ID:id})
}


