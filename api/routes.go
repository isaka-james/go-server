package api

import (
    "net/http"
    "github.com/gorilla/mux"
    "backend/server/api/handlers"
)

// InitRoutes initializes all routes
func InitRoutes() {
    r := mux.NewRouter()

    // Login route
    r.HandleFunc("/login", handlers.LoginHandler).Methods("POST")

    // User route
    r.HandleFunc("/notifications", handlers.NotificationHandler).Methods("GET")


	
    // Set the router for the application
    http.Handle("/", r)
}
