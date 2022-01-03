package mongo

import (
	"github.com/lessbutter/alloff-api/internal/core/domain"
	"github.com/lessbutter/alloff-api/internal/core/repository"
	"go.mongodb.org/mongo-driver/mongo"
)

type userRepo struct {
	col *mongo.Collection
}

type deviceRepo struct {
	col *mongo.Collection
}

func (repo *userRepo) Get(ID string) (*domain.UserDAO, error) {
	return nil, nil
}
func (repo *userRepo) GetByMobile(mobile string) (*domain.UserDAO, error) {
	return nil, nil
}
func (repo *userRepo) Insert(*domain.UserDAO) (*domain.UserDAO, error) {
	return nil, nil
}
func (repo *userRepo) Update(*domain.UserDAO) (*domain.UserDAO, error) {
	return nil, nil
}

func (repo *deviceRepo) GetByDeviceID(deviceID string) (*domain.DeviceDAO, error) {
	return nil, nil
}
func (repo *deviceRepo) ListAllowedByUser(userID string) ([]*domain.DeviceDAO, error) {
	return nil, nil
}
func (repo *deviceRepo) ListAllowed() ([]*domain.DeviceDAO, error) {
	return nil, nil
}
func (repo *deviceRepo) Upsert(*domain.DeviceDAO) (*domain.DeviceDAO, error) {
	return nil, nil
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
