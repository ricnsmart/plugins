package plugins

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type RCN350F struct {
	ID           primitive.ObjectID `bson:"_id,omitempty"`
	CreateAt     string             `bson:"CreateAt"`
	SN           string             `bson:"SN"`
	UpdateAt     string             `bson:"UpdateAt"`
	DeviceID     string             `bson:"DeviceID"`
	GPRSOperator int                `bson:"GPRSOperator"`
	DomainRecord string             `bson:"DomainRecord"`
	CT           int                `bson:"CT"`
	Interval     int                `bson:"Interval"`
	SMSLimit     int                `bson:"SMSLimit"`

	//  rcn350f特有
	Buzzer     byte    `bson:"Buzzer"` // 蜂鸣器开关
	BreakShort byte    `bson:"BreakShort"`
	ICCID      string  `bson:"ICCID"`
	Version    float32 `bson:"Version"` // 固件版本

	Metrics struct {
		DO1 IOTemplate          `bson:"DO1"`
		DO2 IOTemplate          `bson:"DO2"`
		DI1 IOTemplate          `bson:"DI1"`
		DI2 IOTemplate          `bson:"DI2"`
		Ia  CurrentTemplate     `bson:"Ia"`
		Ib  CurrentTemplate     `bson:"Ib"`
		Ic  CurrentTemplate     `bson:"Ic"`
		IR  IRTemplate          `bson:"IR"`
		T1  TemperatureTemplate `bson:"T1"`
		T2  TemperatureTemplate `bson:"T2"`
		T3  TemperatureTemplate `bson:"T3"`
		T4  T4Template          `bson:"T4"`
		Ua  VoltageTemplate     `bson:"Ua"`
		Ub  VoltageTemplate     `bson:"Ub"`
		Uc  VoltageTemplate     `bson:"Uc"`
	} `bson:"Metrics"`
}

func (i *RCN350F) Find(id string, coll *mongo.Collection) (err error) {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	filter := bson.M{ID: id}

	err = coll.FindOne(ctx, filter).Decode(i)

	return
}
