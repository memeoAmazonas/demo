package model

import (
	"time"
)

type Task struct {
	CreatedAt time.Time `bson:"created_at"`
	UpdatedAt time.Time `bson:"updated_at"`
	Text      string    `bson:"text"`
	Completed bool      `bson:"completed"`
}
