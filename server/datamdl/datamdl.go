package datamdl

import (
	"akshu-21/finance-manager/model"
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/BurntSushi/toml"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//Package DB Variable
var db *mongo.Database

// InitConfig initConfig
func InitConfig(fpath string, config interface{}) (toml.MetaData, error) {
	return toml.DecodeFile(fpath, config)
}

//InitMongo :- to initialize connection with mongoDB atlas
func InitMongo() error {
	connectionURI := getConnectionURI()
	client, err := mongo.NewClient(options.Client().ApplyURI(connectionURI))
	if err != nil {
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	// defer client.Disconnect(ctx)
	fmt.Println("Database cluster connected successfully")
	db = client.Database(model.Config.MongoDBName)
	InitDAO()
	return nil
}

//getConnectionURI
func getConnectionURI() string {

	connectionURI := model.Config.MongoURL

	connectionURI = strings.ReplaceAll(connectionURI, "<username>", model.Config.MongoUsername)
	connectionURI = strings.ReplaceAll(connectionURI, "<password>", model.Config.MongoPass)
	fmt.Println(connectionURI)
	return connectionURI

}
