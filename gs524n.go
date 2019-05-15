package plugin

import "go.mongodb.org/mongo-driver/bson/primitive"

type GS524N struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	CreateAt    string             `bson:"CreateAt"`
	SN          string             `bson:"SN"`
	UpdateAt    string             `bson:"UpdateAt"`
	DeviceID    string             `bson:"DeviceID"`
	AlertSwitch bool               `bson:"AlertSwitch"`
	SMSSwitch   bool               `bson:"SMSSwitch"`
}
