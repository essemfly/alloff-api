package mongo

import (
	"context"
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

	var device *domain.DeviceDAO

	// 기존에 없던경우
	if err := repo.col.FindOne(ctx, bson.M{"deviceid": deviceID}).Decode(&device); err != nil {
		newDevice := domain.DeviceDAO{
			ID:                primitive.NewObjectID(),
			DeviceId:          deviceID,
			AllowNotification: allowNotification,
			Created:           time.Now(),
			Updated:           time.Now(),
		}
		if userID != nil {
			newDevice.UserId = *userID
		}

		_, err := repo.col.InsertOne(ctx, newDevice)
		return err
	}

	if userID != nil {
		filter := bson.M{"$or": []bson.M{{"userid": userID}, {"deviceid": deviceID}}}
		_, err := repo.col.UpdateMany(ctx, filter, bson.M{"$set": bson.M{"allownotification": allowNotification, "userid": userID, "updated": time.Now()}})
		if err != nil {
			return err
		}
	} else {
		filter := bson.M{"deviceid": deviceID}
		_, err := repo.col.UpdateMany(ctx, filter, bson.M{"$set": bson.M{"allownotification": allowNotification, "updated": time.Now()}})
		if err != nil {
			return err
		}
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
