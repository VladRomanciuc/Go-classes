package model

import (
	"log"
	"github.com/spf13/viper"
	f "github.com/fauna/faunadb-go/v4/faunadb"
	
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

