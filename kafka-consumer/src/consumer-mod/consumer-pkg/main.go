package main

import (
	"fmt"
	"context"
	"log"
	"encoding/json"
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
	type Trainer struct {
	    Name string
	    Age  int
	    City string
	}
	vehiclesJson := make([]string, 11)
	vehiclesJson[0] = `{"data":{"date-time":{"system":"2021-09-24T01:40:01+00:00"},"gps-info":{"Altitude":"552.8","Date":"240921","HDOP":"0.7","Latitude":"42.70599365","Longitude":"23.31282425","SatelliteUsed":9,"Speed":52.782001495361328,"Time":"014001.00","Validity":"A"},"modem-info":{"signal-quality":"31"},"stop-info":{}},"device-id":"004101FB","device-type":"OBU","hostname":"obu","priority":1,"scheme-version":"v1_0_9","vehicle-id":"132801","id":"ddd21912-421c-4839-8669-153dfc4d6def"}`
vehiclesJson[1] =`{"data":{"date-time":{"system":"2021-09-24T01:40:01+00:00"},"gps-info":{"Altitude":"562.3","Date":"240921","HDOP":"0.7","Latitude":"42.64899063","Longitude":"23.41792297","SatelliteUsed":9,"Speed":0,"Time":"014001.00","Validity":"A"},"modem-info":{"signal-quality":"27"},"stop-info":{}},"device-id":"0040D702","device-type":"OBU","hostname":"obu","priority":1,"scheme-version":"v1_0_9","vehicle-id":"133665","id":"48111729-5685-484e-82cf-e0b217420649"}`
vehiclesJson[2] =`{"data":{"date-time":{"system":"2021-09-24T01:40:01+00:00"},"gps-info":{"Altitude":"530.7","Date":"240921","HDOP":"0.7","Latitude":"42.71719360","Longitude":"23.36151695","SatelliteUsed":9,"Speed":0,"Time":"014001.00","Validity":"A"},"modem-info":{"signal-quality":"23"},"stop-info":{}},"device-id":"004071AF","device-type":"OBU","hostname":"obu","priority":1,"scheme-version":"v1_0_9","vehicle-id":"142154","id":"ee9f5459-e17a-4dc9-a98b-de336abfc2b6"}`
vehiclesJson[3] =`{"data":{"date-time":{"system":"2021-09-24T01:40:02+00:00"},"gps-info":{"Altitude":"564.5","Date":"240921","HDOP":"0.7","Latitude":"42.73235321","Longitude":"23.25246811","SatelliteUsed":9,"Speed":21.668399810791016,"Time":"014001.00","Validity":"A"},"modem-info":{"signal-quality":"31"},"stop-info":{}},"device-id":"00415985","device-type":"OBU","hostname":"obu","priority":1,"scheme-version":"v1_0_9","vehicle-id":"141114","id":"db698e71-fc5b-4594-8c57-69425f4ea5b6"}`
vehiclesJson[4] =`{"data":{"date-time":{"system":"2021-09-24T01:40:02+00:00"},"gps-info":{"Altitude":"564.8","Date":"240921","HDOP":"0.7","Latitude":"42.65491104","Longitude":"23.41275978","SatelliteUsed":8,"Speed":0,"Time":"014001.00","Validity":"A"},"modem-info":{"signal-quality":"31"},"stop-info":{}},"device-id":"00412AC0","device-type":"OBU","hostname":"obu","priority":1,"scheme-version":"v1_0_9","vehicle-id":"341670","id":"f029a109-e153-460e-a83f-14ba87891d15"}`
vehiclesJson[5] =`{"data":{"date-time":{"system":"2021-09-24T01:40:02+00:00"},"gps-info":{"Altitude":"531.4","Date":"240921","HDOP":"0.7","Latitude":"42.71492767","Longitude":"23.35903358","SatelliteUsed":9,"Speed":0,"Time":"014001.00","Validity":"A"},"modem-info":{"signal-quality":"25"},"stop-info":{}},"device-id":"0040E225","device-type":"OBU","hostname":"obu","priority":1,"scheme-version":"v1_0_9","vehicle-id":"142306","id":"ada2f5f2-e04b-4178-9088-65b01a150201"}`
vehiclesJson[6] =`{"data":{"date-time":{"system":"2021-09-24T01:40:02+00:00"},"gps-info":{"Altitude":"562.4","Date":"240921","HDOP":"0.7","Latitude":"42.67761612","Longitude":"23.36746979","SatelliteUsed":9,"Speed":40.743999481201172,"Time":"014001.00","Validity":"A"},"modem-info":{"signal-quality":"24"},"stop-info":{}},"device-id":"0040A662","device-type":"OBU","hostname":"obu","priority":1,"scheme-version":"v1_0_9","vehicle-id":"234180","id":"c91e7890-3093-4ae8-b695-0f7a9d233b92"}`
vehiclesJson[7] =`{"data":{"date-time":{"system":"2021-09-24T01:40:02+00:00"},"gps-info":{"Altitude":"583.3","Date":"240921","HDOP":"0.8","Latitude":"42.69069290","Longitude":"23.28071404","SatelliteUsed":9,"Speed":0,"Time":"014002.00","Validity":"A"},"modem-info":{"signal-quality":"31"},"stop-info":{}},"device-id":"0040723D","device-type":"OBU","hostname":"obu","priority":1,"scheme-version":"v1_0_9","vehicle-id":"232043","id":"95bfbe82-23cb-44e7-9732-962f732ae138"}`
vehiclesJson[8] =`{"data":{"date-time":{"system":"2021-09-24T01:40:02+00:00"},"gps-info":{"Altitude":"555.8","Date":"240921","HDOP":"0.7","Latitude":"42.71362686","Longitude":"23.31405449","SatelliteUsed":9,"Speed":14.630800247192383,"Time":"014001.00","Validity":"A"},"modem-info":{"signal-quality":"19"},"stop-info":{}},"device-id":"00414CB9","device-type":"OBU","hostname":"obu","priority":1,"scheme-version":"v1_0_9","vehicle-id":"342820","id":"33084266-1fcc-48c8-8f29-39f6aa483ab0"}`
vehiclesJson[9] =`{"data":{"date-time":{"system":"2021-09-24T01:40:02+00:00"},"gps-info":{"Altitude":"554.0","Date":"240921","HDOP":"0.7","Latitude":"42.67057800","Longitude":"23.39697075","SatelliteUsed":9,"Speed":20.742399215698242,"Time":"014001.00","Validity":"A"},"modem-info":{"signal-quality":"22"},"stop-info":{}},"device-id":"00413A34","device-type":"OBU","hostname":"obu","priority":1,"scheme-version":"v1_0_9","vehicle-id":"234134","id":"89d72bce-1fdd-4cc3-926a-19d92cecb035"}`
vehiclesJson[11] =`{"data":{"date-time":{"system":"2021-09-24T01:40:02+00:00"},"gps-info":{"Altitude":"576.0","Date":"240921","HDOP":"0.7","Latitude":"42.69020081","Longitude":"23.28133583","SatelliteUsed":9,"Speed":0,"Time":"014002.00","Validity":"A"},"modem-info":{"signal-quality":"31"},"stop-info":{}},"device-id":"00412524","device-type":"OBU","hostname":"obu","priority":1,"scheme-version":"v1_0_9","vehicle-id":"262322","id":"36a031e6-255c-4eaa-a72e-4fcd7e5973c7"}`
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

	collection := client.Database("test").Collection("trainers")

	ash := Trainer{"Ash", 10, "Pallet Town"}
	// misty := Trainer{"Misty", 10, "Cerulean City"}
	// brock := Trainer{"Brock", 15, "Pewter City"}

	fmt.Println("Inserting Ash..")
	insertResult, err := collection.InsertOne(context.TODO(), ash)
	if err != nil {
	    log.Fatal(err)
	}

	fmt.Println("Inserted a single document: ", insertResult.InsertedID)

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
