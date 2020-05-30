package model

type Push struct {
	Index   int64  `json:"hash"`
	Message string `json:"message"`
}
