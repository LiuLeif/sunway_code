// 2022-05-20 10:37
package main

import (
	"context"
	"crypto/sha256"
	"encoding/gob"
	"errors"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type DeviceInfo struct {
	IMEI    string
	Version string
	Vendor  string
}

func (d DeviceInfo) ToCSV() []string {
	return []string{d.IMEI, d.Version}
}

var conn *mongo.Client
var db *mongo.Database
var device_info_collection *mongo.Collection
var account_collection *mongo.Collection

func InitModel() {
	gob.Register(&Account{})
	conn = getConnection()
	db = conn.Database("monitor")
	device_info_collection = db.Collection("device_info")
	account_collection = db.Collection("account")
}

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
	Vendor   string
}

func GetAccount(username string, password string) (Account, error) {
	fmt.Println(username)
	cursor, _ := account_collection.Find(context.TODO(), bson.M{"username": username})
	var accounts []Account
	cursor.All(context.TODO(), &accounts)
	if len(accounts) != 1 {
		return Account{}, errors.New("multiple account found")
	}
	h := sha256.New()
	h.Write([]byte(password))
	bs := h.Sum(nil)
	digest := fmt.Sprintf("%x", bs)
	if digest == accounts[0].Password {
		return accounts[0], nil
	}
	return Account{}, errors.New("wrong password")
}
