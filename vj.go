package plugins

import "go.mongodb.org/mongo-driver/bson/primitive"

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
