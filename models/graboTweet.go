package models

import "time"

//GraboTweet struct de tweets
type GraboTweet struct {
	UserID  string    `bson:"userID" json:"userID,omitempty"`
	Mensaje string    `bson:"mensaje" json:"mensaje,omitempty"`
	Fecha   time.Time `bson:"fecha" json:"fecha,omitempty"`
}
