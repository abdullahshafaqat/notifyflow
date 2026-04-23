package models

type Notification struct {
	ID      string `json:"id" bson:"id"`
	To      string `json:"to" bson:"to"`
	Message string `json:"message" bson:"message"`
	Status  string `json:"status" bson:"status"`
	Retry   int    `json:"retry" bson:"retry"`
}
