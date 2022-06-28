package views

import (
	"time"
)

type Response struct {
	Code int 			`json:"code"`
	Body interface {} 	`json:"body"`
}

type Task struct {
	Id           uint64 	`objectbox:"id"`
	Text         string
	DateCreated  time.Time 	`objectbox:"date"`
	DateFinished time.Time 	`objectbox:"date"`
}