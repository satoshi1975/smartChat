package models

type Profile struct {
	ID        int    `json:"id"`
	UserID    int    `json:"user_id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Bio       string `json:"bio"`
	Friends   []int  `json:"friends,omitempty"`
	Blocked   []int  `json:"blocked,omitempty"`
}
