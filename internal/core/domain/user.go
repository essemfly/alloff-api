package domain

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserDAO struct {
	ID            primitive.ObjectID `json:"id" bson:"_id"`
	Name          string
	Uuid          string `json:"uuid" bson:"uuid"`
	Mobile        string `json:"mobile"`
	Email         string `json:"email"`
	BaseAddress   string `json:"baseaddress"`
	DetailAddress string `json:"detailaddress"`
	Postcode      string `json:"postcode"`
	Created       time.Time
}

type DeviceDAO struct {
	ID                primitive.ObjectID `bson:"_id, omitempty"`
	UserId            string
	DeviceId          string
	AllowNotification bool
	Created           time.Time
	Updated           time.Time
}
