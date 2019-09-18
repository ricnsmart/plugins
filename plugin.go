package plugins

import "go.mongodb.org/mongo-driver/bson/primitive"

type Common struct {
	ID         primitive.ObjectID `bson:"_id,omitempty"`
	CreateAt   string             `bson:"CreateAt"`
	UpdateAt   string             `bson:"UpdateAt"`
	DeviceID   string             `bson:"DeviceID"`
	SN         string             `bson:"SN"`
	DeviceType int                `bson:"DeviceType"`
}

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
	Disable       bool    `bson:"Disable"`
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
