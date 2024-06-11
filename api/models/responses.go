package models

// Struct for the response JSON
type ResponseUpgrade struct {
    Status  string `json:"status"`
    Message string `json:"message"`
}

type ResponseSign struct {
    Status  string `json:"status"`
    Message string `json:"message"`
}
