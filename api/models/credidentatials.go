package models

type Credentials struct {
    Username string `json:"username"`
    Password string `json:"password"`
}

type ResponseAuth struct {
    Status  string `json:"status"`
    Message string `json:"message,omitempty"`
}