package mongo

import (
	"context"
	"strings"
	"time"

	"github.com/lessbutter/alloff-api/internal/core/domain"
	"github.com/lessbutter/alloff-api/internal/core/repository"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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

func (repo *deviceRepo) List(offset, limit int) ([]*domain.DeviceDAO, int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	options := options.Find()
	options.SetSkip(int64(offset))
	options.SetLimit(int64(limit))

	filter := bson.M{}
	totalCount, _ := repo.col.CountDocuments(ctx, filter)
	cursor, err := repo.col.Find(ctx, filter, options)
	if err != nil {
		return nil, 0, err
	}

	var devices []*domain.DeviceDAO
	err = cursor.All(ctx, &devices)
	if err != nil {
		return nil, 0, err
	}

	return devices, int(totalCount), nil
}

func (repo *deviceRepo) ListAllowedByUser(userID string) ([]*domain.DeviceDAO, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	cursor, err := repo.col.Find(ctx, bson.M{"userid": userID, "allownotification": true, "isremoved": false})
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

	cursor, err := repo.col.Find(ctx, bson.M{"allownotification": true, "isremoved": false})
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

func (repo *deviceRepo) Upsert(device *domain.DeviceDAO) (*domain.DeviceDAO, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 120*time.Second)
	defer cancel()

	opts := options.Update().SetUpsert(true)
	filter := bson.M{"_id": device.ID}
	if _, err := repo.col.UpdateOne(ctx, filter, bson.M{"$set": &device}, opts); err != nil {
		return nil, err
	}

	var updatedDevice *domain.DeviceDAO
	if err := repo.col.FindOne(ctx, filter).Decode(&updatedDevice); err != nil {
		return nil, err
	}

	return updatedDevice, nil
}

func (repo *deviceRepo) Delete(ID string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	deviceObjectId, _ := primitive.ObjectIDFromHex(ID)
	filter := bson.M{"_id": deviceObjectId}

	if _, err := repo.col.DeleteOne(ctx, filter); err != nil {
		return err
	}

	return nil
}

func (repo *deviceRepo) RemoveByToken(deviceID string) error {
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
