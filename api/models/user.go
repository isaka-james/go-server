package models


type Notification struct {
    Type    string `json:"notification_type"`
    Message string `json:"message"`
    Time    string `json:"time"`
}

// MemberInfo represents the data structure for member information
type MemberInfo struct {
	Balance        float64 `json:"balance"`
	DaysAttended  int     `json:"days_attended"`
	MembershipPlan string  `json:"membership_plan"`
}

type Member struct {
	Name         string `json:"name"`
	Email        string `json:"email"`
	PhoneNo      string `json:"phone_no"`
	Balance      int    `json:"balance"`
	EmergencyNo  string `json:"emergency_no"`
	LastCheckIn  string `json:"last_check_in"`
    IsVIP string `json:"vip"`

}

