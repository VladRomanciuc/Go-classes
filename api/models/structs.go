package models


type DynamoConfig struct {
	Endpoint string
	Region  string
	Profile string
	KeyID   string
	Key  	string
}


//The structure of data to be handled + a json mapper for encoding/decoding
type Post struct{
	Id 		int64		`json:"Id"`
	Title	string	`json:"Title"`
	Text 	string	`json:"Text"`
}

type ServiceError struct{
	Message string `json:"message"`
}