package domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type GroupRequestStatus string

const (
	GroupRequestStatusPending GroupRequestStatus = "PENDING"
	GroupRequestStatusSuccess GroupRequestStatus = "SUCCESS"
	GroupRequestStatusFailed  GroupRequestStatus = "FAILED"
)

type GroupDAO struct {
	ID              primitive.ObjectID `bson:"_id, omitempty"`
	ExhibitionID    string
	NumUserRequired int
	Users           []*UserDAO
}

type GroupRequestDAO struct {
	ID           primitive.ObjectID `bson:"_id, omitempty"`
	UserID       string
	ExhibitionID string
	GroupID      string
	RequestLink  string
	Status       GroupRequestStatus
	CreatedAt    time.Time
}
