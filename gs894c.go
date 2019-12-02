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
}
