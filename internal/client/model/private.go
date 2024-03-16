package model

import "time"

type PrivateLoginPass struct {
	Id       int
	Title    string
	Type     int
	Login    string
	Password string
	Updated  time.Time
}

type PrivateText struct {
	Id      int
	Title   string
	Type    int
	Text    string
	Updated time.Time
}

type PrivateCard struct {
	Id         int
	Title      string
	Type       int
	CardNumber string
	CVV        string
	Due        string
	Updated    time.Time
}

type PrivateFile struct {
	Id      int
	Title   string
	Type    int
	Path    string
	Binary  []byte
	Updated time.Time
}

type PrivateDataList struct {
	Id    int
	Title string
}
