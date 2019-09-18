package plugins

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type PMC350 struct {
	ID           primitive.ObjectID `bson:"_id,omitempty"`
	CreateAt     string             `bson:"CreateAt"`
	SN           string             `bson:"SN"`
	UpdateAt     string             `bson:"UpdateAt"`
	DeviceID     string             `bson:"DeviceID"`
	DeviceType   int                `bson:"DeviceType"`
	GPRSOperator int                `bson:"GPRSOperator"`
	DomainRecord string             `bson:"DomainRecord"`
	CT           int                `bson:"CT"`
	Interval     int                `bson:"Interval"`
	SMSLimit     int                `bson:"SMSLimit"`
	Metrics      struct {
		Ia CurrentTemplate     `bson:"Ia"`
		Ib CurrentTemplate     `bson:"Ib"`
		Ic CurrentTemplate     `bson:"Ic"`
		IR IRTemplate          `bson:"IR"`
		T1 TemperatureTemplate `bson:"T1"`
		T2 TemperatureTemplate `bson:"T2"`
		T3 TemperatureTemplate `bson:"T3"`
		T4 T4Template          `bson:"T4"`
		Ua VoltageTemplate     `bson:"Ua"`
		Ub VoltageTemplate     `bson:"Ub"`
		Uc VoltageTemplate     `bson:"Uc"`
	} `bson:"Metrics"`
}
