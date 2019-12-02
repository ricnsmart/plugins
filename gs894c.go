package plugins

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type GS894C struct {
	ID         primitive.ObjectID `bson:"_id,omitempty"`
	CreateAt   string             `bson:"CreateAt"`
	SN         string             `bson:"SN"`
	UpdateAt   string             `bson:"UpdateAt"`
	DeviceID   string             `bson:"DeviceID"`
	DeviceType int                `bson:"DeviceType"`
	OneNetID   string             `bson:"OneNetID"`
	Metrics    struct {
		CH4 struct {
			AlertSwitch bool `bson:"AlertSwitch"`
			SMSSwitch   bool `bson:"SMSSwitch"`
		} `bson:"CH4"`
	} `bson:"Metrics"`
}
