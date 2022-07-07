package views


type Config struct {
	Endpoint string
	Region  string
	Profile string
	KeyID   string
	Key  	string
}


//The structure of data to be handled + a json mapper for encoding/decoding
type Post struct{
	Id 		int64		`json:"id"`
	Title	string	`json:"title"`
	Text 	string	`json:"text"`
}