package datamdl

import (
	"context"
	"errors"

	"corelab.mkcl.org/MKCLOS/coredevelopmentplatform/corepkgv2/loggermdl"
	"github.com/labstack/gommon/log"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//MongoDAO - CRUD struct
type MongoDAO struct {
	collectionName string
}

// GetMongoDAO return mongo DAO instance
func GetMongoDAO(collection string) *MongoDAO {
	return &MongoDAO{
		collectionName: collection,
	}
}

//Insert : Insert into DB
func (mg *MongoDAO) Insert(data interface{}) error {
	_, insertError := db.Collection(mg.collectionName).InsertOne(context.TODO(), data)
	if insertError != nil {
		log.Error("Insert : ERROR_WHILE_INSERTION_IN_COLLECTION :", mg.collectionName)
		return errors.New("Insert : ERROR_WHILE_INSERTION_IN_COLLECTION : " + mg.collectionName)
	}
	return nil
}

// Update will update single entry
func (mg *MongoDAO) Update(selector map[string]interface{}, data interface{}) error {

	collection := db.Collection(mg.collectionName)
	_, updateError := collection.UpdateOne(context.Background(), selector, bson.M{"$set": data})
	if updateError != nil {
		log.Error("Update : ERROR_WHILE_UPDATING_IN_COLLECTION :", mg.collectionName)
		return errors.New("Update : ERROR_WHILE_UPDATING_IN_COLLECTION : " + mg.collectionName)
	}
	return nil
}

// FindData will return query for selector
func (mg *MongoDAO) FindData(selector map[string]interface{}) ([]interface{}, error) {

	collection := db.Collection(mg.collectionName)

	cur, err := collection.Find(context.Background(), selector)
	if err != nil {
		log.Error("FindData : ERROR_WHILE_FINDING_IN_COLLECTION :", mg.collectionName, err)
		return nil, errors.New("FindData : ERROR_WHILE_FINDING_IN_COLLECTION : " + mg.collectionName)
	}
	defer cur.Close(context.Background())
	var results []interface{}
	for cur.Next(context.Background()) {
		var result bson.M
		err := cur.Decode(&result)
		if err != nil {
			log.Error("FindData : ERROR_WHILE_DECODING_IN_COLLECTION :", mg.collectionName)
			return nil, errors.New("FindData : ERROR_WHILE_DECODING_IN_COLLECTION : " + mg.collectionName)
		}
		results = append(results, result)
	}

	return results, nil
}

// GetProjectedData will return query for selector and projector
func (mg *MongoDAO) GetProjectedData(selector map[string]interface{}, projector map[string]interface{}) ([]interface{}, error) {

	collection := db.Collection(mg.collectionName)
	ops := &options.FindOptions{}
	ops.Projection = projector
	cur, err := collection.Find(context.Background(), selector, ops)
	if err != nil {
		log.Error("GetProjectedData : ERROR_WHILE_FINDING_IN_COLLECTION :", mg.collectionName)
		return nil, errors.New("GetProjectedData : ERROR_WHILE_FINDING_IN_COLLECTION : " + mg.collectionName)
	}
	defer cur.Close(context.Background())
	var results []interface{}
	for cur.Next(context.Background()) {
		var result bson.M
		err := cur.Decode(&result)
		if err != nil {
			log.Error("GetProjectedData : ERROR_WHILE_DECODING_IN_COLLECTION :", mg.collectionName)
			return nil, errors.New("GetProjectedData : ERROR_WHILE_DECODING_IN_COLLECTION : " + mg.collectionName)
		}
		results = append(results, result)
	}

	return results, nil
}

// Upsert will update single entry
func (mg *MongoDAO) Upsert(selector map[string]interface{}, data interface{}) error {

	collection := db.Collection(mg.collectionName)
	ops := options.UpdateOptions{}
	ops.SetUpsert(true)
	_, updateError := collection.UpdateOne(context.Background(), selector, bson.M{"$set": data}, &ops)
	if updateError != nil {
		log.Error("Upsert : ERROR_WHILE_UPDATING_IN_COLLECTION :", mg.collectionName)
		return errors.New("Upsert : ERROR_WHILE_UPDATING_IN_COLLECTION : " + mg.collectionName)
	}
	return nil
}

// PushData - append in array
func (mg *MongoDAO) PushData(selector map[string]interface{}, data interface{}) error {

	collection := db.Collection(mg.collectionName)
	_, updateError := collection.UpdateOne(context.Background(), selector, bson.M{"$push": data})
	if updateError != nil {
		log.Error("PushData : ERROR_WHILE_UPDATING_IN_COLLECTION :", mg.collectionName)
		return errors.New("PushData : ERROR_WHILE_UPDATING_IN_COLLECTION : " + mg.collectionName)
	}
	return nil
}

// GetAggregateData - return result using aggregation query
func (mg *MongoDAO) GetAggregateData(selector interface{}) ([]interface{}, error) {

	collection := db.Collection(mg.collectionName)
	cur, err := collection.Aggregate(context.Background(), selector)
	if err != nil {
		loggermdl.LogError(err)
		return nil, err
	}
	defer cur.Close(context.Background())
	var results []interface{}
	for cur.Next(context.Background()) {
		var result bson.M
		err := cur.Decode(&result)
		if err != nil {
			loggermdl.LogError(err)
			return nil, err
		}
		results = append(results, result)
	}

	return results, nil
}
