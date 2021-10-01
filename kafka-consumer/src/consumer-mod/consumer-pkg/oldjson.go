
	type SensorFields struct {
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
	vehiclesJson[10] =`{"data":{"date-time":{"system":"2021-09-24T01:40:02+00:00"},"gps-info":{"Altitude":"576.0","Date":"240921","HDOP":"0.7","Latitude":"42.69020081","Longitude":"23.28133583","SatelliteUsed":9,"Speed":0,"Time":"014002.00","Validity":"A"},"modem-info":{"signal-quality":"31"},"stop-info":{}},"device-id":"00412524","device-type":"OBU","hostname":"obu","priority":1,"scheme-version":"v1_0_9","vehicle-id":"262322","id":"36a031e6-255c-4eaa-a72e-4fcd7e5973c7"}`



	
	vehiclesJson := `
	[{"data":{"date-time":{"system":"2021-09-24T01:40:01+00:00"},"gps-info":{"Altitude":"552.8","Date":"240921","HDOP":"0.7","Latitude":"42.70599365","Longitude":"23.31282425","SatelliteUsed":9,"Speed":52.782001495361328,"Time":"014001.00","Validity":"A"},"modem-info":{"signal-quality":"31"},"stop-info":{}},"device-id":"004101FB","device-type":"OBU","hostname":"obu","priority":1,"scheme-version":"v1_0_9","vehicle-id":"132801","id":"ddd21912-421c-4839-8669-153dfc4d6def"},{"data":{"date-time":{"system":"2021-09-24T01:40:01+00:00"},"gps-info":{"Altitude":"562.3","Date":"240921","HDOP":"0.7","Latitude":"42.64899063","Longitude":"23.41792297","SatelliteUsed":9,"Speed":0,"Time":"014001.00","Validity":"A"},"modem-info":{"signal-quality":"27"},"stop-info":{}},"device-id":"0040D702","device-type":"OBU","hostname":"obu","priority":1,"scheme-version":"v1_0_9","vehicle-id":"133665","id":"48111729-5685-484e-82cf-e0b217420649"}
,{"data":{"date-time":{"system":"2021-09-24T01:40:01+00:00"},"gps-info":{"Altitude":"530.7","Date":"240921","HDOP":"0.7","Latitude":"42.71719360","Longitude":"23.36151695","SatelliteUsed":9,"Speed":0,"Time":"014001.00","Validity":"A"},"modem-info":{"signal-quality":"23"},"stop-info":{}},"device-id":"004071AF","device-type":"OBU","hostname":"obu","priority":1,"scheme-version":"v1_0_9","vehicle-id":"142154","id":"ee9f5459-e17a-4dc9-a98b-de336abfc2b6"}
,{"data":{"date-time":{"system":"2021-09-24T01:40:02+00:00"},"gps-info":{"Altitude":"564.5","Date":"240921","HDOP":"0.7","Latitude":"42.73235321","Longitude":"23.25246811","SatelliteUsed":9,"Speed":21.668399810791016,"Time":"014001.00","Validity":"A"},"modem-info":{"signal-quality":"31"},"stop-info":{}},"device-id":"00415985","device-type":"OBU","hostname":"obu","priority":1,"scheme-version":"v1_0_9","vehicle-id":"141114","id":"db698e71-fc5b-4594-8c57-69425f4ea5b6"}
,{"data":{"date-time":{"system":"2021-09-24T01:40:02+00:00"},"gps-info":{"Altitude":"564.8","Date":"240921","HDOP":"0.7","Latitude":"42.65491104","Longitude":"23.41275978","SatelliteUsed":8,"Speed":0,"Time":"014001.00","Validity":"A"},"modem-info":{"signal-quality":"31"},"stop-info":{}},"device-id":"00412AC0","device-type":"OBU","hostname":"obu","priority":1,"scheme-version":"v1_0_9","vehicle-id":"341670","id":"f029a109-e153-460e-a83f-14ba87891d15"}
,{"data":{"date-time":{"system":"2021-09-24T01:40:02+00:00"},"gps-info":{"Altitude":"531.4","Date":"240921","HDOP":"0.7","Latitude":"42.71492767","Longitude":"23.35903358","SatelliteUsed":9,"Speed":0,"Time":"014001.00","Validity":"A"},"modem-info":{"signal-quality":"25"},"stop-info":{}},"device-id":"0040E225","device-type":"OBU","hostname":"obu","priority":1,"scheme-version":"v1_0_9","vehicle-id":"142306","id":"ada2f5f2-e04b-4178-9088-65b01a150201"}
,{"data":{"date-time":{"system":"2021-09-24T01:40:02+00:00"},"gps-info":{"Altitude":"562.4","Date":"240921","HDOP":"0.7","Latitude":"42.67761612","Longitude":"23.36746979","SatelliteUsed":9,"Speed":40.743999481201172,"Time":"014001.00","Validity":"A"},"modem-info":{"signal-quality":"24"},"stop-info":{}},"device-id":"0040A662","device-type":"OBU","hostname":"obu","priority":1,"scheme-version":"v1_0_9","vehicle-id":"234180","id":"c91e7890-3093-4ae8-b695-0f7a9d233b92"}
,{"data":{"date-time":{"system":"2021-09-24T01:40:02+00:00"},"gps-info":{"Altitude":"583.3","Date":"240921","HDOP":"0.8","Latitude":"42.69069290","Longitude":"23.28071404","SatelliteUsed":9,"Speed":0,"Time":"014002.00","Validity":"A"},"modem-info":{"signal-quality":"31"},"stop-info":{}},"device-id":"0040723D","device-type":"OBU","hostname":"obu","priority":1,"scheme-version":"v1_0_9","vehicle-id":"232043","id":"95bfbe82-23cb-44e7-9732-962f732ae138"}
,{"data":{"date-time":{"system":"2021-09-24T01:40:02+00:00"},"gps-info":{"Altitude":"555.8","Date":"240921","HDOP":"0.7","Latitude":"42.71362686","Longitude":"23.31405449","SatelliteUsed":9,"Speed":14.630800247192383,"Time":"014001.00","Validity":"A"},"modem-info":{"signal-quality":"19"},"stop-info":{}},"device-id":"00414CB9","device-type":"OBU","hostname":"obu","priority":1,"scheme-version":"v1_0_9","vehicle-id":"342820","id":"33084266-1fcc-48c8-8f29-39f6aa483ab0"}
,{"data":{"date-time":{"system":"2021-09-24T01:40:02+00:00"},"gps-info":{"Altitude":"554.0","Date":"240921","HDOP":"0.7","Latitude":"42.67057800","Longitude":"23.39697075","SatelliteUsed":9,"Speed":20.742399215698242,"Time":"014001.00","Validity":"A"},"modem-info":{"signal-quality":"22"},"stop-info":{}},"device-id":"00413A34","device-type":"OBU","hostname":"obu","priority":1,"scheme-version":"v1_0_9","vehicle-id":"234134","id":"89d72bce-1fdd-4cc3-926a-19d92cecb035"}
,{"data":{"date-time":{"system":"2021-09-24T01:40:02+00:00"},"gps-info":{"Altitude":"576.0","Date":"240921","HDOP":"0.7","Latitude":"42.69020081","Longitude":"23.28133583","SatelliteUsed":9,"Speed":0,"Time":"014002.00","Validity":"A"},"modem-info":{"signal-quality":"31"},"stop-info":{}},"device-id":"00412524","device-type":"OBU","hostname":"obu","priority":1,"scheme-version":"v1_0_9","vehicle-id":"262322","id":"36a031e6-255c-4eaa-a72e-4fcd7e5973c7"}]`
