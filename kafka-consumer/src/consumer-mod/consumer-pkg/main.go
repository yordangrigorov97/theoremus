package main

import (
	"fmt"
	"context"
	"log"
	"encoding/json"
	"time"
	"os"
	"os/signal"
	"syscall"

	"github.com/segmentio/kafka-go"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	)

func main(){
	fmt.Println("Hello from kafka-consumer")
	kafka_uri, mongo_uri := getConfig()
	consumeKafka(kafka_uri, mongo_uri)

}

// consumeKafka listens for messages on the supplied kafka_uri,
// extends them with new fields, and finally sends them to mongo_uri.
//
// This function performs a blocking poll until a termination signal
// (e.g. CTRL-C) is detected.
//
// Implementation Note: We could use the Generator pattern for reading Kafka messages
// in order to split this function into smaller chunks.
//
// Example message received:
//
// message:= `{"data":{"date-time":{"system":"2021-09-24T01:40:01+00:00"},"gps-info":{"Altitude":"552.8","Date":"240921","HDOP":"0.7","Latitude":"42.70599365","Longitude":"23.31282425","SatelliteUsed":9,"Speed":52.782001495361328,"Time":"014001.00","Validity":"A"},"modem-info":{"signal-quality":"31"},"stop-info":{}},"device-id":"004101FB","device-type":"OBU","hostname":"obu","priority":1,"scheme-version":"v1_0_9","vehicle-id":"132801","id":"ddd21912-421c-4839-8669-153dfc4d6def"}`
func consumeKafka(kafka_uri string, mongo_uri string){
	fmt.Println("reading vehicles..")
	r := kafka.NewReader(kafka.ReaderConfig{
	    Brokers:   []string{kafka_uri},
	    Topic:     "vehicles",
	    Partition: 0,
	    MinBytes:  0,
	    MaxBytes:  600e6, // 600MB
	})
	r.SetOffset(0)
	cleanOnInt(r) // CTRL-C handler

	// Loop until CTRL-C
	for {
	    message, err := r.ReadMessage(context.Background())
	    if err != nil { // Probably Kafka still booting
		fmt.Println(err)
		time.Sleep(2 * time.Second)
	    }
	    fmt.Printf("Message at offset %d: %s = %s\n", message.Offset, string(message.Key), string(message.Value))

	    // Parse message
	    var message_obj SensorFields
	    err = json.Unmarshal(message.Value, &message_obj)
	    if err != nil { // If JSON invalid: skip message
	        fmt.Println(message.Value)
	        fmt.Println(err)
	        continue
	    }
	    // Extend message
	    addTimeKeys(&message_obj)

	    // Send message
	    writeMongo(&message_obj, mongo_uri)
	}

}
// cleanOnInt handles CTRL-C. This helps us exit
// cleanly out of the consumer's infinite loop
func cleanOnInt(r *kafka.Reader){
    c := make(chan os.Signal)
    signal.Notify(c, os.Interrupt, syscall.SIGTERM)
    go func() {
        <-c
	fmt.Println("done reading...")

	if err := r.Close(); err != nil {
	    log.Fatal("failed to close reader:", err)
	}

	fmt.Println("goodbye from kafka-consumer")
        os.Exit(1)
    }()
}

// addTimeKeys adds "IDDay" and "IDHour" to the message to
// facilitate the queries which will be executed on the DB
func addTimeKeys(message_obj *SensorFields){
	rfc3339time := message_obj.Data.DateTime.System

	d, h := trucateTime(rfc3339time)
	message_obj.IDDay = d
	message_obj.IDHour = h
}


// getMongoCollection gives us objects that help us interact
// with Mongo and also tests the connection to the DB
func getMongoCollection(coll_name string, mongo_uri string) (*mongo.Collection, *mongo.Client){
	clientOptions := options.Client().ApplyURI(mongo_uri)

	fmt.Println("Connecting to MongoDB..")
	client, err := mongo.Connect(context.Background(), clientOptions)

	if err != nil {
	    log.Fatal(err)
	}

	fmt.Println("Pinging MongoDB")
	err = client.Ping(context.Background(), nil)

	if err != nil {
	    log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")

	collection := client.Database("theoremus").Collection(coll_name)
	return collection, client
}

// writeMongo sends the object to mongo_uri
//
// Implementation note: if the structure of the rawmessage is highly variable
// (e.g. missing fields, extra fields), we could consider using a dynamic
// structure like bson's primitive.M:
//
// var mybson primitive.M
// err := bson.UnmarshalExtJSON(
// 	[]byte(JSONstr), true, &mybson)
// if err != nil {
// 	panic(err)
// }
// rfc3339str := mybson.Map()["data"].(primitive.D).Map()["date-time"].(primitive.D).Map()["system"]
func writeMongo(message_obj *SensorFields, mongo_uri string){
	fmt.Printf("%+v\n", message_obj)

	vehicles, client := getMongoCollection("vehicles", mongo_uri)
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	insertResult, err := vehicles.InsertOne(ctx, message_obj)

	// Check for any insertion errors
	if err != nil {
		fmt.Println("InsertOne ERROR:", err)
	} else {
		fmt.Println("InsertOne() API result:", insertResult)
	}


	err = client.Disconnect(context.Background())

	if err != nil {
	    log.Fatal(err)
	}
	fmt.Println("Connection to MongoDB closed.")

}

// trucateTime returns abbreviated versions of the input time.
// Example input: time("2021-09-24T01:40:02Z")
// Output 1: time("2021-09-24T01:00:00Z")
// Output 2: time("2021-09-24T00:00:00Z")
func trucateTime(t time.Time) (time.Time, time.Time) {
	h := t.Truncate(time.Hour)
	d := t.Truncate(24 * time.Hour)
	return d, h

}

// getConfig queries the os for the following environment variables:
//
// KAFKA_URI (default kafka:9092)
//
// MONGO_URI (default mongo:27017)
//
// if a variables is not found, the default value is returned
func getConfig() (string, string) {
	default_kafka_uri := "kafka:9092"
	default_mongo_uri := "mongodb://root:root@mongo:27017"
	kafka_uri:= os.Getenv("KAFKA_URI")
	mongo_uri:= os.Getenv("MONGO_URI")
	if kafka_uri == ""{
		fmt.Printf("KAFKA_URI not present, using default %v", default_kafka_uri)
		kafka_uri = default_kafka_uri
	} else {
		fmt.Printf("KAFKA_URI present: %v", kafka_uri)
	}
	if default_mongo_uri == ""{
		fmt.Printf("MONGO_URI not present, using default %v", default_mongo_uri)
		mongo_uri = default_mongo_uri
	} else {
		fmt.Printf("MONGO_URI present: %v", mongo_uri)
	}
	return kafka_uri, mongo_uri
}
