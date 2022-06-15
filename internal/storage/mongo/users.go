package mongo

import (
	"context"
	"log"
	"strings"
	"time"

	"github.com/lessbutter/alloff-api/internal/core/domain"
	"github.com/lessbutter/alloff-api/internal/core/repository"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type userRepo struct {
	col *mongo.Collection
}

type deviceRepo struct {
	col *mongo.Collection
}

func (repo *userRepo) Get(ID string) (*domain.UserDAO, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	productObjectId, _ := primitive.ObjectIDFromHex(ID)
	filter := bson.M{"_id": productObjectId}
	var user *domain.UserDAO
	if err := repo.col.FindOne(ctx, filter).Decode(&user); err != nil {
		return nil, err
	}
	return user, nil
}

func (repo *userRepo) GetByMobile(mobile string) (*domain.UserDAO, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	filter := bson.M{"mobile": mobile}
	var user *domain.UserDAO
	if err := repo.col.FindOne(ctx, filter).Decode(&user); err != nil {
		return nil, err
	}
	return user, nil
}

func (repo *userRepo) Insert(user *domain.UserDAO) (*domain.UserDAO, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	user.ID = primitive.NewObjectID()
	oid, err := repo.col.InsertOne(ctx, user)
	if err != nil {
		return nil, err
	}

	var newUser *domain.UserDAO
	filter := bson.M{"_id": oid.InsertedID}
	err = repo.col.FindOne(ctx, filter).Decode(&newUser)
	if err != nil {
		return nil, err
	}

	return newUser, nil
}

func (repo *userRepo) Update(user *domain.UserDAO) (*domain.UserDAO, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := repo.col.UpdateOne(ctx, bson.M{"_id": user.ID}, bson.M{"$set": &user})
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (repo *deviceRepo) GetByDeviceID(deviceID string) (*domain.DeviceDAO, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	var device *domain.DeviceDAO
	if err := repo.col.FindOne(ctx, bson.M{"deviceid": deviceID}).Decode(&device); err != nil {
		return nil, err
	}

	return device, nil
}

func (repo *deviceRepo) ListAllowedByUser(userID string) ([]*domain.DeviceDAO, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	cursor, err := repo.col.Find(ctx, bson.M{"userid": userID, "allownotification": true})
	if err != nil {
		return nil, err
	}

	var devices []*domain.DeviceDAO
	err = cursor.All(ctx, &devices)
	if err != nil {
		return nil, err
	}

	return devices, nil
}

func (repo *deviceRepo) ListAllowed() ([]*domain.DeviceDAO, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	cursor, err := repo.col.Find(ctx, bson.M{"allownotification": true})
	if err != nil {
		return nil, err
	}

	var devices []*domain.DeviceDAO
	err = cursor.All(ctx, &devices)
	if err != nil {
		return nil, err
	}

	return devices, nil
}

func (repo *deviceRepo) UpdateDevices(deviceID string, allowNotification bool, userID *string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	filter := bson.M{"deviceid": deviceID}
	var devices []*domain.DeviceDAO

	// 기존에 하나만 있던 경우
	cursor, err := repo.col.Find(ctx, filter)
	if err != nil {
		log.Println("err on update devices", err)
		return err
	}

	err = cursor.All(ctx, &devices)
	if err != nil {
		log.Println("decode error on devices", err)
		return err
	}

	if len(devices) == 0 {
		newDevice := domain.DeviceDAO{
			ID:                primitive.NewObjectID(),
			DeviceId:          deviceID,
			AllowNotification: allowNotification,
			IsRemoved:         false,
			Created:           time.Now(),
			Updated:           time.Now(),
		}
		if userID != nil {
			newDevice.UserId = *userID
		}

		_, err := repo.col.InsertOne(ctx, newDevice)
		return err
	}

	for idx, device := range devices {
		if idx > 0 {
			device.IsRemoved = true
			device.Updated = time.Now()
			_, err := repo.col.UpdateByID(ctx, device.ID, bson.M{"$set": &device})
			if err != nil {
				log.Println("err on update device", err)
			}
			continue
		}

		if userID != nil {
			filter := bson.M{"$or": []bson.M{{"userid": userID}, {"deviceid": deviceID}}}
			_, err := repo.col.UpdateMany(ctx, filter, bson.M{"$set": bson.M{"allownotification": allowNotification, "userid": userID, "updated": time.Now()}})
			if err != nil {
				log.Println("err on update user devices", err)
			}
		}

		device.IsRemoved = false
		device.AllowNotification = allowNotification
		device.Updated = time.Now()
		_, err := repo.col.UpdateByID(ctx, device.ID, bson.M{"$set": &device})
		if err != nil {
			log.Println("err on update device", err)
		}
	}

	return nil
}

func (repo *deviceRepo) MakeRemoved(deviceID string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	var device *domain.DeviceDAO

	trimmedDeviceId := strings.Trim(deviceID, "*")
	filter := bson.M{
		"deviceid": primitive.Regex{Pattern: trimmedDeviceId, Options: "i"},
	}
	err := repo.col.FindOne(ctx, filter).Decode(&device)
	if err != nil {
		return err
	}

	device.IsRemoved = true
	_, err = repo.col.UpdateOne(ctx, bson.M{"_id": device.ID}, bson.M{"$set": &device})
	if err != nil {
		return err
	}

	return nil
}

func MongoUsersRepo(conn *MongoDB) repository.UsersRepository {
	return &userRepo{
		col: conn.userCol,
	}
}

func MongoDevicesRepo(conn *MongoDB) repository.DevicesRepository {
	return &deviceRepo{
		col: conn.deviceCol,
	}
}
