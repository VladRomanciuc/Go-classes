package model

type Entry struct {
	Token  string `fauna:"token"`
	UserID string `fauna:"user_id"`
	ItemID string `fauna:"item_id"`
	Value  int    `fauna:"value"`
}