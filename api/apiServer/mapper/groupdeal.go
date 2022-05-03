package mapper

import (
	"github.com/lessbutter/alloff-api/api/apiServer/model"
	"github.com/lessbutter/alloff-api/internal/core/domain"
)

func MapGroupDaoToGroup(groupDao *domain.GroupDAO) *model.Group {
	users := []*model.User{}
	for _, userDao := range groupDao.Users {
		users = append(users, MapUserDaoToUser(userDao))
	}

	return &model.Group{
		ID:               groupDao.ID.Hex(),
		ExhibitionID:     groupDao.ExhibitionID,
		NumUsersRequired: groupDao.NumUsersRequired,
		Users:            users,
	}
}
