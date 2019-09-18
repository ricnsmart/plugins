package plugins

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type RCNVJ struct {
	ID           primitive.ObjectID `bson:"_id,omitempty"`
	CreateAt     string             `bson:"CreateAt"`
	SN           string             `bson:"SN"`
	UpdateAt     string             `bson:"UpdateAt"`
	DeviceID     string             `bson:"DeviceID"`
	DeviceType   int                `bson:"DeviceType"`
	GPRSOperator int                `bson:"GPRSOperator"`
	DomainRecord string             `bson:"DomainRecord"`
	Interval     int                `bson:"Interval"`
	Version      float32            `bson:"Version"`
	GPRSRSSI     int                `bson:"GPRSRSSI"`
	ModulesNum   int                `bson:"ModulesNum"`
	SMSLimit     int                `bson:"SMSLimit"`
}
