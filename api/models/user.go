package models


type Notification struct {
    Type    string `json:"notification_type"`
    Message string `json:"message"`
    Time    string `json:"time"`
}


