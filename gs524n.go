package plugins

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type GS524N struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	CreateAt string             `bson:"CreateAt"`
	SN       string             `bson:"SN"`
	UpdateAt string             `bson:"UpdateAt"`
	DeviceID string             `bson:"DeviceID"`
	SMSLimit int                `bson:"SMSLimit"`
	Metrics  struct {
		Battery struct {
			Alert       int  `bson:"Alert"`
			AlertSwitch bool `bson:"AlertSwitch"`
			SMSSwitch   bool `bson:"SMSSwitch"`
		} `bson:"Battery"`
		Smoke struct {
			AlertSwitch bool `bson:"AlertSwitch"`
			SMSSwitch   bool `bson:"SMSSwitch"`
		} `bson:"Smoke"`
		TearDown struct {
			AlertSwitch bool `bson:"AlertSwitch"`
			SMSSwitch   bool `bson:"SMSSwitch"`
		} `bson:"TearDown"`
	} `bson:"Metrics"`
}
