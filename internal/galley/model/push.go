package model

type Push struct {
	Index   int64  `json:"hash" binding:"required"`
	Message string `json:"message" binding:"required"`
}
