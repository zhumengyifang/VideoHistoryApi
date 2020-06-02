package MongoDbUtil

import (
	"context"
	"gindemo/internal/Config"
	"gindemo/internal/Model/MongodbModel"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

var pool *mongo.Client

func init() {
	pool, _ = newPool()
}

func newPool() (*mongo.Client, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	opts := options.Client().ApplyURI(Config.GetMongo().Host).SetMaxPoolSize(Config.GetMongo().MaxPoolSize).SetMinPoolSize(Config.GetMongo().MinPoolSize)
	client, err := mongo.Connect(ctx, opts)
	if err != nil {
		return nil, err
	}
	return client, nil
}

func InsertApiLog(log MongodbModel.MongoApiLog) (interface{}, error) {
	collection := pool.Database("History").Collection("ApiLog")
	insertResult, err := collection.InsertOne(context.TODO(), log)
	if err != nil {
		return nil, err
	}
	return insertResult.InsertedID, nil
}

func InsertTaskLog(log MongodbModel.MongoTaskLog) (interface{}, error) {
	collection := pool.Database("History").Collection(log.TaskType)
	insertResult, err := collection.InsertOne(context.TODO(), log)
	if err != nil {
		return nil, err
	}
	return insertResult.InsertedID, nil
}
