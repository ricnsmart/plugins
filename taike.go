package plugins

import "go.mongodb.org/mongo-driver/bson/primitive"

type (
	TK82 struct {
		ID         primitive.ObjectID `bson:"_id,omitempty"`
		CreateAt   string             `bson:"CreateAt"`
		UpdateAt   string             `bson:"UpdateAt"`
		DeviceID   string             `bson:"DeviceID"`
		SN         string             `bson:"SN"`
		DeviceType int                `bson:"DeviceType"`
		Max        int                `bson:"Max"`
		Min        int                `bson:"Min"`
	}

	TK83 struct {
		ID         primitive.ObjectID `bson:"_id,omitempty"`
		CreateAt   string             `bson:"CreateAt"`
		UpdateAt   string             `bson:"UpdateAt"`
		DeviceID   string             `bson:"DeviceID"`
		SN         string             `bson:"SN"`
		DeviceType int                `bson:"DeviceType"`
		Length     int                `bson:"Length"`
		Width      int                `bson:"Width"`
		Height     int                `bson:"Height"`
		Max        int                `bson:"Max"`
		Min        int                `bson:"Min"`
	}
)
