package logging

// Message structure or Model
type LogMessage struct {
	CreatedDate   string `json:"createddate" bson:"createddate"`
	CreatedTime   string `json:"createdtime" bson:"createdtime"`
	Level         string `json:"level"   bson:"level"`
	ServiceName   string `json:"servicename"   bson:"servicename"`
	CallingMethod string `json:"callingmethod"   bson:"callingmethod"`
	Host          string `json:"host"   bson:"host"`
	Body          string `json:"body"   bson:"body"`
	Latency       string `json:"latency"   bson:"latency"`
}
