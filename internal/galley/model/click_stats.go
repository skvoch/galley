package model

type ClickStats struct {
	Hash   string `json:"hash" binding:"required"`
	Count  int64  `json:"count" binding:"required"`
	Period string `json:"period" binding:"required"`
}
