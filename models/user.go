package models

// User model
type User struct {
	ID             int64   `json:"id"`
	AttachableSgID string  `json:"attachable_sgid"`
	Name           string  `json:"name"`
	EmailAddress   string  `json:"email_address"`
	PersonableType string  `json:"personable_type"`
	Title          string  `json:"title"`
	Bio            string  `json:"bio"`
	CreatedAt      string  `json:"created_at"`
	UpdatedAt      string  `json:"updated_at"`
	Admin          bool    `json:"admin"`
	Owner          bool    `json:"owner"`
	TimeZone       string  `json:"time_zone"`
	AvatarURL      string  `json:"avatar_url"`
	Company        Company `json:"company"`
}
