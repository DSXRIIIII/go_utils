package main

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Person struct {
	Name    string `bson:"name"`
	Age     int    `bson:"age"`
	Address string `bson:"address"`
}

func main() {
	// 创建一个 MongoDB 客户端连接
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal(err)
	}

	// 创建一个上下文对象
	ctx := context.Background()

	// 连接到 MongoDB 服务
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	// 获取要操作的集合
	collection := client.Database("testdb").Collection("persons")

	// 插入数据
	p1 := Person{Name: "Tom", Age: 20, Address: "Shanghai"}
	_, err = collection.InsertOne(ctx, p1)
	if err != nil {
		log.Fatal(err)
	}

	// 查询数据
	var p2 Person
	err = collection.FindOne(ctx, bson.M{"name": "Tom"}).Decode(&p2)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Name: %s, Age: %d, Address: %s\n", p2.Name, p2.Age, p2.Address)

	// 更新数据
	filter := bson.M{"name": "Tom"}
	update := bson.M{"$set": bson.M{"age": 25}}
	_, err = collection.UpdateOne(ctx, filter, update)
	if err != nil {
		log.Fatal(err)
	}

	// 删除数据
	_, err = collection.DeleteOne(ctx, bson.M{"name": "Tom"})
	if err != nil {
		log.Fatal(err)
	}

	// 断开连接
	err = client.Disconnect(ctx)
	if err != nil {
		log.Fatal(err)
	}
}
