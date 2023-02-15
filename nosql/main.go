package main

import (
	"fmt"
	"sync"
	"time"
	"log"
	"context"

	"go.mongodb.org/mongo-driver/mongo"
    // "go.mongodb.org/mongo-driver/mongo/options"
    "go.mongodb.org/mongo-driver/bson"
    . "nosql/src/pool"
    "strconv"
)

var (
	DBPool *MongoPool
	once sync.Once
)

type employee struct {
	Name 			string `bson:"name"`
	Gender			string `bson:"gender"`
	Nik 			string `bson:"nik"`
	PhoneNumber		string `bson:"phone_number"`
	Address 		string `bson:"address"`
}

type religion struct {
	id 	  int
	name  string
}

func getEmployees() {
	ctx := context.Background()
	conn, err := DBPool.GetConnection()
	if err != nil {
		log.Fatal(err.Error())
		return
	}
	defer DBPool.CloseConnection(conn)
	collection := GetCollection(conn, "hris", "employee")
	rows, err := collection.Find(ctx, bson.M{})
	if err != nil {
		fmt.Println("ERROR GET DATA MONGODB", err)
	}
	defer rows.Close(ctx)

	var result []employee
	for rows.Next(ctx) {
		var row employee
		err := rows.Decode(&row)
		if err != nil {
			fmt.Println("ERROR GET DATA MONGODB", err)
		}

		result = append(result, row)
	}

	for idx, data := range result {
		fmt.Println(strconv.Itoa(idx + 1) + ". Nama: " + data.Name + " Alamat: " + data.Address)
	}
}

func getEmployeeDetail(name string) {
	ctx := context.Background()
	conn, err := DBPool.GetConnection()
	if err != nil {
		log.Fatal(err.Error())
		return
	}
	defer DBPool.CloseConnection(conn)
	collection := GetCollection(conn, "hris", "employee")
	var result employee
	err = collection.FindOne(ctx, bson.M{"name": name}).Decode(&result)
	if err != nil {
		fmt.Println("ERROR GET DATA MONGODB", err)
	}

	fmt.Println("RESULT EMPLOYEE DETAIL", result)
}

func init() {
	once.Do(func() {
		DBPool = &MongoPool{
			Pool:        make(chan *mongo.Client, 10),
			Connections: 0,                           
			Timeout:     10 * time.Second,
			Uri:         "mongodb://localhost:27017",
			PoolSize:    10,
		}
	})
}

func main() {
	getEmployees()
	getEmployeeDetail("Ayu")
}