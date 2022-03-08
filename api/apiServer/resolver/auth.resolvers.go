package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"

	"github.com/lessbutter/alloff-api/api/apiServer"
	"github.com/lessbutter/alloff-api/api/apiServer/mapper"
	"github.com/lessbutter/alloff-api/api/apiServer/middleware"
	"github.com/lessbutter/alloff-api/api/apiServer/model"
	"github.com/lessbutter/alloff-api/config/ioc"
	"github.com/lessbutter/alloff-api/internal/core/domain"
)

func (r *mutationResolver) RegisterNotification(ctx context.Context, deviceID string, allowNotification bool, userID *string) (*model.Device, error) {
	err := ioc.Repo.Devices.UpdateDevices(deviceID, allowNotification, userID)
	if err != nil {
		return nil, err
	}

	device, _ := ioc.Repo.Devices.GetByDeviceID(deviceID)
	return mapper.MapDeviceDaoToDevice(device), nil
}

func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewUser) (string, error) {
	_, err := ioc.Repo.Users.GetByMobile(input.Mobile)
	if err == nil {
		return "", errors.New("already registered mobile")
	}

	var user domain.UserDAO
	user.Uuid = input.UUID
	user.Mobile = input.Mobile

	if input.Email != nil {
		user.Email = *input.Email
	}

	if input.BaseAddress != nil {
		user.BaseAddress = *input.BaseAddress
	}

	if input.DetailAddress != nil {
		user.DetailAddress = *input.DetailAddress
	}

	if input.Postcode != nil {
		user.Postcode = *input.Postcode
	}

	_, err = ioc.Repo.Users.Insert(&user)
	if err != nil {
		return "", err
	}

	token, err := middleware.GenerateToken(user.Mobile, user.Uuid)
	if err != nil {
		return "", err
	}
	return token, nil
}

func (r *mutationResolver) UpdateUserInfo(ctx context.Context, input model.UserInfoInput) (*model.User, error) {
	user := middleware.ForContext(ctx)
	if user == nil {
		return nil, errors.New("invalid token")
	}

	if input.Mobile != nil {
		user.Mobile = *input.Mobile
	}

	if input.Email != nil {
		user.Email = *input.Email
	}

	if input.BaseAddress != nil {
		user.BaseAddress = *input.BaseAddress
	}

	if input.DetailAddress != nil {
		user.DetailAddress = *input.DetailAddress
	}

	if input.Postcode != nil {
		user.Postcode = *input.Postcode
	}

	if input.Name != nil {
		user.Name = *input.Name
	}

	if input.PersonalCustomsNumber != nil {
		user.PersonalCustomsNumber = *input.PersonalCustomsNumber
	}

	newUser, err := ioc.Repo.Users.Update(user)
	if err != nil {
		return mapper.MapUserDaoToUser(user), err
	}

	return mapper.MapUserDaoToUser(newUser), nil
}

func (r *mutationResolver) Login(ctx context.Context, input model.Login) (string, error) {
	uuid := input.UUID
	mobile := input.Mobile

	user, err := ioc.Repo.Users.GetByMobile(mobile)
	if err != nil {
		return "", err
	}

	if user.Uuid != uuid {
		user.Uuid = input.UUID
		_, err := ioc.Repo.Users.Update(user)
		if err != nil {
			return "", err
		}
	}

	token, err := middleware.GenerateToken(mobile, uuid)
	if err != nil {
		return "", err
	}
	return token, nil
}

func (r *mutationResolver) RefreshToken(ctx context.Context, input model.RefreshTokenInput) (string, error) {
	mobile, uuid, err := middleware.ParseToken(input.Token)
	if mobile == "" {
		return "", err
	}

	user, err := ioc.Repo.Users.GetByMobile(mobile)
	if err != nil {
		return "", err
	}

	if user.Uuid != uuid {
		return "", errors.New("devices changed")
	}

	token, err := middleware.GenerateToken(mobile, uuid)
	if err != nil {
		return "", err
	}
	return token, nil
}

func (r *queryResolver) User(ctx context.Context) (*model.User, error) {
	user := middleware.ForContext(ctx)
	if user == nil {
		return nil, errors.New("invalid token")
	}

	return mapper.MapUserDaoToUser(user), nil
}

// Mutation returns apiServer.MutationResolver implementation.
func (r *Resolver) Mutation() apiServer.MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }
