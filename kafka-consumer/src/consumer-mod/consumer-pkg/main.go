package main

import (
	"fmt"
	"context"
	"log"
	"encoding/json"
	"reflect"
	"time"
	// "os"

	"github.com/segmentio/kafka-go"
	// "github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	)

func main(){
	// make a new reader that consumes from topic-A, partition 0, at offset 42
	fmt.Println("Hello from kafka-consumer")
	fmt.Println("reading vehicles..")
	r := kafka.NewReader(kafka.ReaderConfig{
	    Brokers:   []string{"kafka:9092"},
	    Topic:     "vehicles",
	    Partition: 0,
	    MinBytes:  0, // 10e3, // 10KB
	    MaxBytes:  600e6, // 600MB
	})
	r.SetOffset(0)

	// for {
	//     m, err := r.ReadMessage(context.Background())
	//     if err != nil {
	// 	break
	//     }
	//     fmt.Printf("message at offset %d: %s = %s\n", m.Offset, string(m.Key), string(m.Value))
	// }
	fmt.Println("done reading...")

	if err := r.Close(); err != nil {
	    log.Fatal("failed to close reader:", err)
	}

	writeMongo()
	fmt.Println("goodbye from kafka-consumer")

}
func writeMongo(){
	type MongoFields struct {
	Data struct {
		DateTime struct {
			System time.Time `json:"system"`
		} `json:"date-time"`
		GpsInfo struct {
			Altitude      string  `json:"Altitude"`
			Date          string  `json:"Date"`
			Hdop          string  `json:"HDOP"`
			Latitude      string  `json:"Latitude"`
			Longitude     string  `json:"Longitude"`
			SatelliteUsed int     `json:"SatelliteUsed"`
			Speed         float64 `json:"Speed"`
			Time          string  `json:"Time"`
			Validity      string  `json:"Validity"`
		} `json:"gps-info"`
		ModemInfo struct {
			SignalQuality string `json:"signal-quality"`
		} `json:"modem-info"`
		StopInfo struct {
		} `json:"stop-info"`
	} `json:"data"`
	DeviceID      string `json:"device-id"`
	DeviceType    string `json:"device-type"`
	Hostname      string `json:"hostname"`
	Priority      int    `json:"priority"`
	SchemeVersion string `json:"scheme-version"`
	VehicleID     string `json:"vehicle-id"`
	ID            string `json:"id"`
	}

	vehicleJson := `{"data":{"date-time":{"system":"2021-09-24T01:40:01+00:00"},"gps-info":{"Altitude":"552.8","Date":"240921","HDOP":"0.7","Latitude":"42.70599365","Longitude":"23.31282425","SatelliteUsed":9,"Speed":52.782001495361328,"Time":"014001.00","Validity":"A"},"modem-info":{"signal-quality":"31"},"stop-info":{}},"device-id":"004101FB","device-type":"OBU","hostname":"obu","priority":1,"scheme-version":"v1_0_9","vehicle-id":"132801","id":"ddd21912-421c-4839-8669-153dfc4d6def"}`

	var doc MongoFields
	err := json.Unmarshal([]byte(vehicleJson), &doc)

	// Print MongoDB docs object type
	fmt.Println("nMongoFields Docs:", reflect.TypeOf(doc))

	// Set client options
	uri := "mongodb://root:root@mongo:27017"
	clientOptions := options.Client().ApplyURI(uri)

	fmt.Println("Connecting to MongoDB..")
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
	    log.Fatal(err)
	}

	fmt.Println("Pinging MongoDB")
	err = client.Ping(context.TODO(), nil)

	if err != nil {
	    log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")

	ctx, _ := context.WithTimeout(context.Background(), 3*time.Second)

	collection := client.Database("theoremus").Collection("vehicles")

	fmt.Println("ndoc _id:", doc.ID)
	fmt.Println("doc Field Str:", doc.ID)

	// Call the InsertOne() method and pass the context and doc objects
	insertResult, insertErr := collection.InsertOne(ctx, doc)

	// Check for any insertion errors
	if insertErr != nil {
		fmt.Println("InsertOne ERROR:", insertErr)
	} else {
		fmt.Println("InsertOne() API result:", insertResult)
	}


	err = client.Disconnect(context.TODO())

	if err != nil {
	    log.Fatal(err)
	}
	fmt.Println("Connection to MongoDB closed.")


}

func readMongo(){
	uri := "mongodb://root:root@mongo:27017"
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()
	coll := client.Database("sample_mflix").Collection("movies")
	title := "Back to the Future"
	var result bson.M
	err = coll.FindOne(context.TODO(), bson.D{{"title", title}}).Decode(&result)
	if err == mongo.ErrNoDocuments {
		fmt.Printf("No document was found with the title %s\n", title)
		return
	}
	if err != nil {
		panic(err)
	}
	jsonData, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", jsonData)

}
