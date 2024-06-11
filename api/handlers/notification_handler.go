package handlers

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	"backend/server/api/models"
	"backend/server/api/utils"
	_ "github.com/lib/pq"
)


func NotificationHandler(w http.ResponseWriter, r *http.Request) {

        username := r.URL.Query().Get("username")

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
    
        var rows *sql.Rows
        if username == "" {
            // If username is not provided, fetch all notifications
            //rows, err = db.Query("SELECT notification_type, message, time FROM notifications")
            return
        } else {
            // If username is provided, fetch notifications for that username
            rows, err = db.Query(`
            SELECT n.notification_type, n.notification_content AS message, n.timestamp AS time
            FROM notification n
            WHERE n.username = $1
            LIMIT 10
        `, username)
        }
        if err != nil {
            log.Fatal("Error querying the database:", err)
        }
        defer rows.Close()
    
        notifications := make([]models.Notification, 0)
        for rows.Next() {
            var notification models.Notification
            if err := rows.Scan(&notification.Type, &notification.Message, &notification.Time); err != nil {
                log.Fatal("Error scanning row:", err)
            }
            notifications = append(notifications, notification)
        }
    
        if err := rows.Err(); err != nil {
            log.Fatal("Error iterating over rows:", err)
        }
    
        w.Header().Set("Content-Type", "application/json")
        if err := json.NewEncoder(w).Encode(notifications); err != nil {
            log.Fatal("Error encoding JSON:", err)
        }
    
}