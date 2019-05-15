package plugins

import "go.mongodb.org/mongo-driver/mongo"

const ID = "DeviceID"

type CurrentTemplate struct {
	WarnSwitch  bool    `bson:"WarnSwitch"`
	AlertSwitch bool    `bson:"AlertSwitch"`
	SMSSwitch   bool    `bson:"SMSSwitch"`
	Min         int     `bson:"Min"`
	Max         int     `bson:"Max"`
	Scale       int     `bson:"Scale"`
	Warn        float32 `bson:"Warn"`
	Alert       float32 `bson:"Alert"`
	SPL         int     `bson:"SPL"`
}

type TemperatureTemplate struct {
	WarnSwitch  bool    `bson:"WarnSwitch"`
	AlertSwitch bool    `bson:"AlertSwitch"`
	SMSSwitch   bool    `bson:"SMSSwitch"`
	Min         int     `bson:"Min"`
	Max         int     `bson:"Max"`
	Scale       int     `bson:"Scale"`
	Warn        float32 `bson:"Warn"`
	Alert       float32 `bson:"Alert"`
}

type VoltageTemplate struct {
	WarnSwitch    bool    `bson:"WarnSwitch"`
	AlertSwitch   bool    `bson:"AlertSwitch"`
	SMSSwitch     bool    `bson:"SMSSwitch"`
	Min           int     `bson:"Min"`
	Max           int     `bson:"Max"`
	Scale         int     `bson:"Scale"`
	Warn          float32 `bson:"Warn"`
	Alert         float32 `bson:"Alert"`
	WarnSMSSwitch bool    `bson:"WarnSMSSwitch"`
}

type IOTemplate struct {
	AlertSwitch bool `bson:"AlertSwitch"`
	SMSSwitch   bool `bson:"SMSSwitch"`
}

type IRTemplate struct {
	WarnSwitch  bool    `bson:"WarnSwitch"`
	AlertSwitch bool    `bson:"AlertSwitch"`
	SMSSwitch   bool    `bson:"SMSSwitch"`
	Warn        float32 `bson:"Warn"`
	Alert       float32 `bson:"Alert"`
	SPL         int     `bson:"SPL"`
}

type T4Template struct {
	WarnSwitch  bool    `bson:"WarnSwitch"`
	AlertSwitch bool    `bson:"AlertSwitch"`
	SMSSwitch   bool    `bson:"SMSSwitch"`
	Warn        float32 `bson:"Warn"`
	Alert       float32 `bson:"Alert"`
}

type Finder interface {
	Find(id string, coll *mongo.Collection) error
}
