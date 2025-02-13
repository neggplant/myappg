package utils

import (
	"context"
	"myappg/config"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"
)

var mongoClient *mongo.Client

func InitDB() {
	// 设置 MongoDB 客户端选项
	clientOptions := options.Client().ApplyURI(config.AppConfig.MongoDB.URI)

	// 连接到 MongoDB
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		Logger.Fatal("Failed to connect to MongoDB", zap.Error(err))
	}

	// 检查连接
	err = client.Ping(context.Background(), nil)
	if err != nil {
		Logger.Fatal("Failed to ping MongoDB", zap.Error(err))
	}

	mongoClient = client
	Logger.Info("Connected to MongoDB!")
}

func GetDB() *mongo.Database {
	return mongoClient.Database(config.AppConfig.MongoDB.Database)
}
