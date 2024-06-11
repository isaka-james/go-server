package utils

import (
	"database/sql"
	"fmt"
	"os"
	"strconv"

	_ "github.com/lib/pq"
)

/* Remember this is postgres initialization function,
 * If you use another database like mysql/redis ,
 * Please replace this function with your database connection logic
 * InitializeDatabase initializes the database connection
 */


func InitializeDatabase() (*sql.DB, error) {
    // Database connection parameters
    host :=  os.Getenv("SERVERNAME_DB")
    port,_ := strconv. Atoi(os.Getenv("PORT_DB")) 
    user := os.Getenv("USERNAME_DB")
    password := os.Getenv("PASSWORD_DB")
    dbname := os.Getenv("DATABASE")

    // Construct the connection string
    connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
        host, port, user, password, dbname)

    // Open a connection to the database
    db, err := sql.Open("postgres", connStr)
    if err != nil {
        return nil, err
    }

    // Check if the connection is successful
    err = db.Ping()
    if err != nil {
        return nil, err
    }

    // Return the database connection
    return db, nil
}
