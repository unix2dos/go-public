package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	// 是否测试环境
	isTest bool = false

	// 测试环境
	TestMongoLink = "mongodb://poros_c:PKWyKN512EbNswf4XNBjXehDM2cNjxwR@s-2ze804aac801b914-pub.mongodb.rds.aliyuncs.com:3717/jinquan"
	TestDB        = "jinquan"

	// !!!正式环境!!!
	ProdMongoLink = "mongodb://user-core-service:0533e9234fa25a7f70d5a9cc29415488@s-2ze1de83b7e68e94.mongodb.rds.aliyuncs.com:3717,s-2ze0700f4cd183a4.mongodb.rds.aliyuncs.com:3717,s-2ze1a63a895afc84.mongodb.rds.aliyuncs.com:3717,s-2ze61d1c3f83a114.mongodb.rds.aliyuncs.com:3717,s-2zec3e106b3b6e94.mongodb.rds.aliyuncs.com:3717,s-2zedac4025f24254.mongodb.rds.aliyuncs.com:3717/poros"
	ProdDB        = "poros"
)

func main() {
	client := getClient()
	GetUser(client)
}

func GetUser(client *mongo.Client) {
	mongoDB := TestDB
	if !isTest {
		mongoDB = ProdDB
	}
	conn := client.Database(mongoDB).Collection("user")
	var result bson.M
	filter := bson.M{"_id": "284650506241"}
	err := conn.FindOne(context.Background(), filter).Decode(&result)
	a, _ := json.Marshal(result)
	fmt.Println("res:", string(a))
	fmt.Println("err:", err)
}

func getClient() *mongo.Client {
	mongoLink := TestMongoLink
	if !isTest {
		mongoLink = ProdMongoLink
	}
	// Set client options
	clientOptions := options.Client().ApplyURI(mongoLink)
	// Connect to MongoDB
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	return client
}
