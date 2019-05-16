package plugins

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type GS524N struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	CreateAt    string             `bson:"CreateAt"`
	SN          string             `bson:"SN"`
	UpdateAt    string             `bson:"UpdateAt"`
	DeviceID    string             `bson:"DeviceID"`
	AlertSwitch bool               `bson:"AlertSwitch"`
	SMSSwitch   bool               `bson:"SMSSwitch"`
	SMSLimit    int                `bson:"SMSLimit"`
	Metrics     struct {
		Battery struct {
			Alert float64 `bson:"Alert"`
		} `bson:"Battery"`
	} `bson:"Metrics"`
}

func (i *GS524N) Find(id string, coll *mongo.Collection) (err error) {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	filter := bson.M{ID: id}

	err = coll.FindOne(ctx, filter).Decode(i)

	return
}
