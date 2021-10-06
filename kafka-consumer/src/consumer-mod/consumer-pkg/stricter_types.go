
package main
import (
	"time"
	"encoding/json"
	"strconv"
)

type StrictSensorFields struct {
Data struct {
	DateTime struct {
		System time.Time `json:"system" bson:"system"`
	} `json:"date-time" bson:"date-time"`
	GpsInfo struct {
		Altitude      float64 `json:"Altitude" bson:"Altitude"`
		Date          int     `json:"Date" bson:"Date"`
		Hdop          float64 `json:"HDOP" bson:"Hdop"`
		Latitude      float64 `json:"Latitude" bson:"Latitude"`
		Longitude     float64 `json:"Longitude" bson:"Longitude"`
		SatelliteUsed int     `json:"SatelliteUsed" bson:"SatelliteUsed"`
		Speed         float64 `json:"Speed" bson:"Speed"`
		Time          float64 `json:"Time" bson:"Time"`
		Validity      string  `json:"Validity" bson:"Validity"`
	} `json:"gps-info"`
	ModemInfo struct {
		SignalQuality int     `json:"signal-quality" bson:"signal-quality"`
	} `json:"modem-info" bson:"modem-info`
	StopInfo struct {
	} `json:"stop-info" bson:"stop-info"`
} `json:"data" bson:"data"`
DeviceID      string `json:"device-id" bson:"device-id"`
DeviceType    string `json:"device-type" bson:"device-type"`
Hostname      string `json:"hostname" bson:"hostname"`
Priority      int    `json:"priority" bson:"priority"`
SchemeVersion string `json:"scheme-version" bson:"scheme-version"`
VehicleID     int    `json:"vehicle-id" bson:"vehicle-id"`
ID            string `json:"id" bson:"id"`
IDDay	      time.Time `json:"IDDay" bson:"IDDay"`
IDHour        time.Time `json:"IDHour" bson:"IDHour"`
}
func cleanJSON(rawMessage string) []byte{
	messageMap := make(map[string]interface{})

        err := json.Unmarshal([]byte(rawMessage), &messageMap)
	if err != nil {
	    panic(err)
	}
	var s string
	var i int64
	var f float64

	s = messageMap["data"].(map[string]interface{})["gps-info"].(map[string]interface{})["Altitude"].(string)
	f, err = strconv.ParseFloat(s, 64)
	if err != nil {
		panic(err)
	}
	messageMap["data"].(map[string]interface{})["gps-info"].(map[string]interface{})["Altitude"] = f

	s = messageMap["data"].(map[string]interface{})["gps-info"].(map[string]interface{})["Date"].(string)
	i, err = strconv.ParseInt(s, 10, 32)
	if err != nil {
		panic(err)
	}
	messageMap["data"].(map[string]interface{})["gps-info"].(map[string]interface{})["Date"] = i

	s = messageMap["data"].(map[string]interface{})["gps-info"].(map[string]interface{})["HDOP"].(string)
	f, err = strconv.ParseFloat(s, 64)
	if err != nil {
		panic(err)
	}
	messageMap["data"].(map[string]interface{})["gps-info"].(map[string]interface{})["HDOP"] = f

	s = messageMap["data"].(map[string]interface{})["gps-info"].(map[string]interface{})["Latitude"].(string)
	f, err = strconv.ParseFloat(s, 64)
	if err != nil {
		panic(err)
	}
	messageMap["data"].(map[string]interface{})["gps-info"].(map[string]interface{})["Latitude"] = f

	s = messageMap["data"].(map[string]interface{})["gps-info"].(map[string]interface{})["Longitude"].(string)
	f, err = strconv.ParseFloat(s, 64)
	if err != nil {
		panic(err)
	}
	messageMap["data"].(map[string]interface{})["gps-info"].(map[string]interface{})["Longitude"] = f

	s = messageMap["data"].(map[string]interface{})["gps-info"].(map[string]interface{})["Time"].(string)
	f, err = strconv.ParseFloat(s, 64)
	if err != nil {
		panic(err)
	}
	messageMap["data"].(map[string]interface{})["gps-info"].(map[string]interface{})["Time"] = f

	s = messageMap["data"].(map[string]interface{})["modem-info"].(map[string]interface{})["signal-quality"].(string)
	i, err = strconv.ParseInt(s, 10, 32)
	if err != nil {
		panic(err)
	}
	messageMap["data"].(map[string]interface{})["modem-info"].(map[string]interface{})["signal-quality"] = i

	s = messageMap["vehicle-id"].(string)
	i, err = strconv.ParseInt(s, 10, 32)
	if err != nil {
		panic(err)
	}
	messageMap["vehicle-id"] = i

	newMessage, err := json.Marshal(messageMap)
	if err != nil {
		panic(err)
	}
	return newMessage

}
