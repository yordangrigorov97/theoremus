package main
import (
	"time"
)

type SensorFields struct {
Data struct {
	DateTime struct {
		System time.Time `json:"system" bson:"system"`
	} `json:"date-time" bson:"date-time"`
	GpsInfo struct {
		Altitude      string `json:"Altitude" bson:"Altitude"`
		Date          string     `json:"Date" bson:"Date"`
		Hdop          string `json:"HDOP" bson:"Hdop"`
		Latitude      string `json:"Latitude" bson:"Latitude"`
		Longitude     string `json:"Longitude" bson:"Longitude"`
		SatelliteUsed int64     `json:"SatelliteUsed" bson:"SatelliteUsed"`
		Speed         float64 `json:"Speed" bson:"Speed"`
		Time          string `json:"Time" bson:"Time"`
		Validity      string  `json:"Validity" bson:"Validity"`
	} `json:"gps-info"`
	ModemInfo struct {
		SignalQuality string     `json:"signal-quality" bson:"signal-quality"`
	} `json:"modem-info" bson:"modem-info`
	StopInfo struct {
	} `json:"stop-info" bson:"stop-info"`
} `json:"data" bson:"data"`
DeviceID      string `json:"device-id" bson:"device-id"`
DeviceType    string `json:"device-type" bson:"device-type"`
Hostname      string `json:"hostname" bson:"hostname"`
Priority      int64    `json:"priority" bson:"priority"`
SchemeVersion string `json:"scheme-version" bson:"scheme-version"`
VehicleID     string    `json:"vehicle-id" bson:"vehicle-id"`
ID            string `json:"id" bson:"id"`
IDDay	      time.Time `json:"IDDay" bson:"IDDay"`
IDHour        time.Time `json:"IDHour" bson:"IDHour"`
}
