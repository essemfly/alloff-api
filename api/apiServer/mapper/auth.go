package mapper

import (
	"github.com/lessbutter/alloff-api/api/apiServer/model"
	"github.com/lessbutter/alloff-api/internal/core/domain"
)

func MapUserDaoToUser(userDao *domain.UserDAO) *model.User {
	return &model.User{
		ID:                    userDao.ID.Hex(),
		UUID:                  userDao.Uuid,
		Mobile:                userDao.Mobile,
		Name:                  &userDao.Name,
		Email:                 &userDao.Email,
		BaseAddress:           &userDao.BaseAddress,
		DetailAddress:         &userDao.DetailAddress,
		Postcode:              &userDao.Postcode,
		PersonalCustomsNumber: &userDao.PersonalCustomsNumber,
	}
}

func MapDeviceDaoToDevice(deviceDao *domain.DeviceDAO) *model.Device {
	return &model.Device{
		ID:                deviceDao.ID.Hex(),
		DeviceID:          deviceDao.DeviceId,
		UserID:            &deviceDao.UserId,
		AllowNotification: deviceDao.AllowNotification,
	}
}
