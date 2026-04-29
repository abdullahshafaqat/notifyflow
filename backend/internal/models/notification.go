package models

import "time"

type Notification struct {
	ID        string    `json:"id" bson:"id"`
	To        string    `json:"to" bson:"to"`
	Message   string    `json:"message" bson:"message"`
	Status    string    `json:"status" bson:"status"`
	Retry     int       `json:"retry" bson:"retry"`
	SendAt    time.Time `json:"send_at" bson:"send_at"`
	LastError string    `json:"last_error,omitempty" bson:"last_error,omitempty"`
}
