// 2022-05-20 10:37
package main

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type DeviceInfo struct {
	IMEI   string
	Vendor string
}

var conn = getConnection()
var db = conn.Database("monitor")
var device_info_collection = db.Collection("device_info")

func getConnection() *mongo.Client {
	clientOptions := options.Client().ApplyURI("mongodb://myuser:mypassword@localhost:27017")
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MongoDB!")
	return client
}

func GetDeviceInfo(config map[string]interface{}) []DeviceInfo {
	cursor, _ := device_info_collection.Find(context.TODO(), config)
	var ret []DeviceInfo
	cursor.All(context.TODO(), &ret)
	return ret
}

func InsertDeviceInfo(info DeviceInfo) {
	device_info_collection.InsertOne(context.TODO(), info)
}

func IsUserValid(accout string, password string) bool {
	return true
}
