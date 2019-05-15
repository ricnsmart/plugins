package plugins

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type VJ struct {
	ID           primitive.ObjectID `bson:"_id,omitempty"`
	CreateAt     string             `bson:"CreateAt"`
	SN           string             `bson:"SN"`
	UpdateAt     string             `bson:"UpdateAt"`
	DeviceID     string             `bson:"DeviceID"`
	Category     string             `bson:"Category"`
	ControllerID string             `bson:"ControllerID"`
	Lan          string             `bson:"Lan"`
	Version      string             `bson:"Version"`
	Lines        Lines              `bson:"Lines"`
}

type Line struct {
	LineNo       int     `bson:"LineNo"`
	LineID       string  `bson:"LineID"`
	IsLeakage    float64 `bson:"IsLeakage"`
	Model        string  `bson:"Model"`
	Over         float64 `bson:"Over"`
	Max          float64 `bson:"Max"`
	LeakValue    float64 `bson:"LeakValue"`
	ErrLeakValue float64 `bson:"Err_LeakValue" json:"Err_LeakValue"`
	Under        float64 `bson:"Under"`
}

type Lines []Line

func (i *VJ) Find(id string, coll *mongo.Collection) (err error) {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	filter := bson.M{ID: id}

	err = coll.FindOne(ctx, filter).Decode(i)

	return
}
