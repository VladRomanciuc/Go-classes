package views

import (
	
)

type Response struct {
	Code int 			`json:"code"`
	Body interface {} 	`json:"body"`
}

type Entry struct {
	Token  string `fauna:"token"`
	UserID string `fauna:"user_id"`
	ItemID string `fauna:"item_id"`
	Value  int    `fauna:"value"`
}

