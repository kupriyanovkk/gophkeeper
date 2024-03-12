package model

import "github.com/google/uuid"

type User struct {
	Login    string
	Password string
	ID       *uuid.UUID
}
