package plugins

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type RCNVJ struct {
	ID           primitive.ObjectID `bson:"_id,omitempty"`
	CreateAt     string             `bson:"CreateAt"`
	SN           string             `bson:"SN"`
	UpdateAt     string             `bson:"UpdateAt"`
	DeviceID     string             `bson:"DeviceID"`
	GPRSOperator int                `bson:"GPRSOperator"`
	DomainRecord string             `bson:"DomainRecord"`
	Interval     int                `bson:"Interval"`
	Version      float32            `bson:"Version"`
	GPRSRSSI     int                `bson:"GPRSRSSI"`
	ModulesNum   int                `bson:"ModulesNum"`
	SMSLimit     int                `bson:"SMSLimit"`
}

func (i *RCNVJ) Find(id string, coll *mongo.Collection) (err error) {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	filter := bson.M{ID: id}

	err = coll.FindOne(ctx, filter).Decode(i)

	return
}
