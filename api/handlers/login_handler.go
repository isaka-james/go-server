package handlers

import (
	"database/sql"
    "encoding/json"
    "log"
    "net/http"
	"backend/server/api/models"
    "backend/server/api/utils"

)


// LoginHandler handles user login
func LoginHandler(w http.ResponseWriter, r *http.Request) {

    // Parse the request body to get the login credentials
    var creds models.Credentials

    // convert the post names to our model so we can interact with backend
    err := json.NewDecoder(r.Body).Decode(&creds)

    // If the user didn't send the valid post names
    if err != nil {
        response := models.ResponseAuth{Status: "error", Message: "Invalid request payload"}
        json.NewEncoder(w).Encode(response)
        w.WriteHeader(http.StatusBadRequest)
        return
    }

    // Initialize the database connection
    db, err := utils.InitializeDatabase()
    if err != nil {
        log.Println("Failed to connect to the database:", err)
        response := models.ResponseAuth{Status: "error", Message: "Internal server error"}
        json.NewEncoder(w).Encode(response)
        w.WriteHeader(http.StatusInternalServerError)
        return
    }
    defer db.Close()

    // Here we interact with the database
    // Query the database for the username
    var storedPassword string
    query := "SELECT password FROM users WHERE username=$1"
    err = db.QueryRow(query, creds.Username).Scan(&storedPassword)
    if err != nil {
        if err == sql.ErrNoRows {
            response := models.ResponseAuth{Status: "error", Message: "Invalid username or password"}
            json.NewEncoder(w).Encode(response)
            w.WriteHeader(http.StatusUnauthorized)
        } else {
            log.Println("Database query error:", err)
            response := models.ResponseAuth{Status: "error", Message: "Internal server error"}
            json.NewEncoder(w).Encode(response)
            w.WriteHeader(http.StatusInternalServerError)
        }
        return
    }


    // Check if the provided password matches the stored password
    if creds.Password != storedPassword {
        response := models.ResponseAuth{Status: "error", Message: "Invalid username or password"}
        json.NewEncoder(w).Encode(response)
        w.WriteHeader(http.StatusUnauthorized)
        return
    }

    // If authentication is successful, respond with a success message
    response := models.ResponseAuth{Status: "success"}
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(response)
}
