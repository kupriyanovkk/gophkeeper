package model

import "time"

type PrivateData struct {
	ID      uint32    `json:"id"`
	Title   string    `json:"title"`
	Type    uint32    `json:"type"`
	UserID  uint32    `json:"user_id"`
	Content []byte    `json:"content"`
	Updated time.Time `json:"updated"`
	Deleted bool      `json:"deleted"`
}

type PrivateDataType struct {
	ID    uint   `json:"id"`
	Title string `json:"title"`
}
