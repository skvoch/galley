package model

type ClickStats struct {
	Hash   string `json:"hash"`
	Count  int64  `json:"count"`
	Period string `json:"period"`
}
