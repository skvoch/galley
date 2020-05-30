package model

type Task struct {
	ID          int    `json:"id,omitempty"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      int    `json:"status"`
	Urgency     int    `json:"urgency"`
}
