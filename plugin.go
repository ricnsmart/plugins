package plugins

import "go.mongodb.org/mongo-driver/bson/primitive"

type (
	Common struct {
		ID         primitive.ObjectID `bson:"_id,omitempty"`
		CreateAt   string             `bson:"CreateAt"`
		UpdateAt   string             `bson:"UpdateAt"`
		DeviceID   string             `bson:"DeviceID"`
		SN         string             `bson:"SN"`
		DeviceType int                `bson:"DeviceType"`
	}

	CurrentTemplate struct {
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

	TemperatureTemplate struct {
		WarnSwitch  bool    `bson:"WarnSwitch"`
		AlertSwitch bool    `bson:"AlertSwitch"`
		SMSSwitch   bool    `bson:"SMSSwitch"`
		Min         int     `bson:"Min"`
		Max         int     `bson:"Max"`
		Scale       int     `bson:"Scale"`
		Warn        float32 `bson:"Warn"`
		Alert       float32 `bson:"Alert"`
	}

	VoltageTemplate struct {
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

	IOTemplate struct {
		AlertSwitch bool `bson:"AlertSwitch"`
		SMSSwitch   bool `bson:"SMSSwitch"`
	}

	IRTemplate struct {
		WarnSwitch  bool    `bson:"WarnSwitch"`
		AlertSwitch bool    `bson:"AlertSwitch"`
		SMSSwitch   bool    `bson:"SMSSwitch"`
		Warn        float32 `bson:"Warn"`
		Alert       float32 `bson:"Alert"`
		SPL         int     `bson:"SPL"`
	}

	T4Template struct {
		WarnSwitch  bool    `bson:"WarnSwitch"`
		AlertSwitch bool    `bson:"AlertSwitch"`
		SMSSwitch   bool    `bson:"SMSSwitch"`
		Warn        float32 `bson:"Warn"`
		Alert       float32 `bson:"Alert"`
	}
)
