package domain

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type GroupRequestStatus string

const (
	GroupRequestStatusPending GroupRequestStatus = "PENDING"
	GroupRequestStatusSuccess GroupRequestStatus = "SUCCESS"
	GroupRequestStatusFailed  GroupRequestStatus = "FAILED"
)

type GroupDAO struct {
	ID               primitive.ObjectID `bson:"_id, omitempty"`
	ExhibitionID     string
	NumUsersRequired int
	Users            []*UserDAO
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

type GroupdealTicketDAO struct {
	ID           primitive.ObjectID `bson:"_id, omitempty"`
	ExhibitionID string
	UserID       string
	Group        *GroupDAO
	CreatedAt    time.Time
}

type GroupRequestListParams struct {
	GroupID      *string
	UserID       *string
	ExhibitionID *string
}
