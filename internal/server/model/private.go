package model

import "time"

type PrivateData struct {
	ID      uint32    `bson:"_id"`
	Title   string    `bson:"title"`
	Type    uint32    `bson:"type"`
	UserID  uint32    `bson:"user_id"`
	Content []byte    `bson:"content"`
	Updated time.Time `bson:"updated"`
	Deleted bool      `bson:"deleted"`
}

type PrivateDataType struct {
	ID    uint
	Title string
}
