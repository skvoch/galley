package model

type Task struct {
	ID          int    `json:"id,omitempty" binding:"required"`
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
	Status      int    `json:"status" binding:"required"`
	Urgency     int    `json:"urgency" binding:"required"`
}
