package main

import (
    "net/http"
    "log"
    "backend/server/api"
    "backend/server/api/config"
)

func main() {

    config.LoadConfig()

    // Initialize routes
    api.InitRoutes()

    

    corsHandler := corsMiddleware(http.DefaultServeMux)

    log.Println("Starting server on :80")
    if err := http.ListenAndServe(":80", corsHandler); err != nil {
        log.Fatal("ListenAndServe: ", err)
    }
}


// corsMiddleware adds CORS headers to the response
func corsMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Access-Control-Allow-Origin", "*")
        w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
        w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

        // Handle preflight request
        if r.Method == "OPTIONS" {
            return
        }

        next.ServeHTTP(w, r)
    })
}
