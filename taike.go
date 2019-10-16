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
		Max        float64            `bson:"Max"`
		Min        float64            `bson:"Min"`
	}

	TK83 struct {
		ID         primitive.ObjectID `bson:"_id,omitempty"`
		CreateAt   string             `bson:"CreateAt"`
		UpdateAt   string             `bson:"UpdateAt"`
		DeviceID   string             `bson:"DeviceID"`
		SN         string             `bson:"SN"`
		DeviceType int                `bson:"DeviceType"`
		Length     float64            `bson:"Length"`
		Width      float64            `bson:"Width"`
		Height     float64            `bson:"Height"`
		Max        float64            `bson:"Max"`
		Min        float64            `bson:"Min"`
	}
)
