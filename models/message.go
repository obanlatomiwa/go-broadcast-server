package models

import "time"

type Message struct {
	ClientId string    `json:"client_id"`
	Text     string    `json:"text"`
	Date     time.Time `json:"date"`
}
