package domain

import (
	"time"

	"github.com/lessbutter/alloff-api/api/apiServer/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserDAO struct {
	ID            primitive.ObjectID `json:"id" bson:"_id"`
	Name          string             `json:"name"`
	Uuid          string             `json:"uuid" bson:"uuid"`
	Mobile        string             `json:"mobile"`
	Email         string             `json:"email"`
	BaseAddress   string             `json:"baseaddress"`
	DetailAddress string             `json:"detailaddress"`
	Postcode      string             `json:"postcode"`
	CreatedAt     time.Time          `json:"created_at"`
}

func (userDao *UserDAO) GetUserAddress() string {
	addr := ""
	if userDao.BaseAddress != "" {
		addr = userDao.BaseAddress + " " + userDao.DetailAddress
	}
	return addr
}

func (userDao *UserDAO) ToDTO() *model.User {
	return &model.User{
		ID:            userDao.ID.Hex(),
		UUID:          userDao.Uuid,
		Mobile:        userDao.Mobile,
		Name:          &userDao.Name,
		Email:         &userDao.Email,
		BaseAddress:   &userDao.BaseAddress,
		DetailAddress: &userDao.DetailAddress,
		Postcode:      &userDao.Postcode,
	}
}

type DeviceDAO struct {
	ID                primitive.ObjectID `bson:"_id, omitempty"`
	UserId            string
	DeviceId          string
	AllowNotification bool
	Created           time.Time
	Updated           time.Time
}

func (deviceDao *DeviceDAO) ToDTO() *model.Device {
	return &model.Device{
		ID:                deviceDao.ID.Hex(),
		DeviceID:          deviceDao.DeviceId,
		UserID:            &deviceDao.UserId,
		AllowNotification: deviceDao.AllowNotification,
	}
}
