package model

import (
	"log"
	"encoding/json"
	"net/http"
	"strings"
	"github.com/spf13/viper"
	"github.com/gorilla/mux"
	f "github.com/fauna/faunadb-go/v4/faunadb"
	"github.com/VladRomanciuc/Go-classes/api/views"
	
)

func viperEnv(key string) string {
	
	viper.SetConfigFile(".env")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Error while reading config file %s", err)
	  }
	value, ok := viper.Get(key).(string)
	if !ok {
		log.Fatalf("Invalid type assertion")
	  }
	
	return value
}

var (
	key = viperEnv("FAUNA_KEY")
	endPoint = f.Endpoint("https://db.eu.fauna.com:443")
	adminClient = f.NewFaunaClient(key, endPoint)
	dbName = "api-go"
)

func Connect() {
	res, err := adminClient.Query(
		f.If(
			f.Exists(f.Database(dbName)),
			true,
			f.CreateDatabase(f.Obj{"name": dbName})))

	if err != nil {
		panic(err)
	}

	if res != f.BooleanV(true) {
		log.Printf("Created Database: %s\n %s", dbName, res)
	} else {
		log.Printf("Database: %s, Already Exists\n %s", dbName, res)
	}

}


func DbClient() (dbClient *f.FaunaClient) {
	var res f.Value
	var err error

	var secret string

	res, err = adminClient.Query(
		f.CreateKey(f.Obj{
			"database": f.Database(dbName),
			"role":     "server"}))

	if err != nil {
		panic(err)
	}

	err = res.At(f.ObjKey("secret")).Get(&secret)

	if err != nil {
		panic(err)
	}

	log.Printf("Database: %s, specifc key: %s\n%s", dbName, secret, res)

	dbClient = adminClient.NewSessionClient(secret)

	return dbClient
}

func ExtractToken(headers http.Header) string {
	header := headers.Get("Authorization")
	return strings.Split(header, " ")[1]
}

// SendJSON - send json with 200 status
func SendJSON(w http.ResponseWriter, data interface{}) {
	js, err := json.Marshal(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

// SendSuccess - send 200 message with empyt json body
func SendSuccess(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte("{}"))
}

// SendInternalServerError - send 500 error with error in message key
func SendInternalServerError(w http.ResponseWriter, err error) {
	sendError(w, err, http.StatusInternalServerError)
}

// SendNotFound - send 404 error with error in message key
func SendNotFound(w http.ResponseWriter, err error) {
	sendError(w, err, http.StatusNotFound)
}

func sendError(w http.ResponseWriter, err error, code int) {
	data := map[string]string{
		"message": err.Error(),
	}

	if js, _ := json.Marshal(data); js != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(code)
		w.Write(js)
		return
	}
	http.Error(w, err.Error(), http.StatusInternalServerError)
}



func CreateEntry(r *http.Request, value int) views.Entry {
	token := ExtractToken(r.Header)
	vars := mux.Vars(r)
	userID := vars["user_id"]
	itemID := vars["item_id"]

	return views.Entry {
		Token:  token,
		UserID: userID,
		ItemID: itemID,
		Value:  value,
	}
}

func Mapper(vs []f.RefV, f func(f.RefV) interface{}) []interface{} {
	vsm := make([]interface{}, len(vs))
	for i, v := range vs {
		vsm[i] = f(v)
	}
	return vsm
}