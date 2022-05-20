// 2022-05-20 10:37
package main

import (
	"context"
	"crypto/sha256"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
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
var account_collection = db.Collection("account")

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

type Account struct {
	Username string
	Password string
}

func IsUserValid(username string, password string) bool {
	cursor, _ := account_collection.Find(context.TODO(), bson.M{"username": username})
	var accounts []Account
	cursor.All(context.TODO(), &accounts)
	if len(accounts) != 1 {
		return false
	}
	h := sha256.New()
	h.Write([]byte(password))
	bs := h.Sum(nil)
	digest := fmt.Sprintf("%x", bs)
	fmt.Println(digest)
	fmt.Println(accounts[0].Password)    
	return digest == accounts[0].Password
}
