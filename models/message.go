package models

// MessageRequest model
type MessageRequest struct {
	Command     string `json:"command"`
	Creator     User   `json:"creator"`
	CallbackURL string `json:"callback_url"`
}

// MessageResponse model
type MessageResponse struct {
	Content string `json:"content"`
}
