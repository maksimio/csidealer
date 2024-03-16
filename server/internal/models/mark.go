package models

type Mark struct {
	Id        string `json:"id"`
	Text      string `json:"text"`
	IsActive  bool   `json:"is_active"`
	Timestamp int64  `json:"timestamp"`
}
