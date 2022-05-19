package scripts

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/lessbutter/alloff-api/config/ioc"
	"github.com/lessbutter/alloff-api/internal/core/domain"
	"github.com/lessbutter/alloff-api/internal/utils"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	"time"
)

func AddMockGroup(exhibition *domain.ExhibitionDAO, completed bool) (*domain.GroupDAO, error) {
	newGroup := &domain.GroupDAO{}

	switch completed {
	case true:
		users := []*domain.UserDAO{}
		for i := 1; i <= exhibition.NumUsersRequired; i++ {
			newUser, err := generateMockUser()
			if err != nil {
				return nil, err
			}
			users = append(users, newUser)
		}

		newGroup = &domain.GroupDAO{
			ID:               primitive.NewObjectID(),
			ExhibitionID:     exhibition.ID.Hex(),
			NumUsersRequired: exhibition.NumUsersRequired,
			Users:            users,
		}

		for _, user := range users {
			newGroupRequest := &domain.GroupRequestDAO{
				ID:           primitive.NewObjectID(),
				UserID:       user.ID.Hex(),
				ExhibitionID: exhibition.ID.Hex(),
				GroupID:      newGroup.ID.Hex(),
				Status:       domain.GroupRequestStatusSuccess,
				CreatedAt:    time.Now(),
			}

			_, err := ioc.Repo.GroupRequest.Insert(newGroupRequest)
			if err != nil {
				log.Println("err on insert mock group request : ", err)
				return nil, err
			}
		}

		_, err := ioc.Repo.Groups.Insert(newGroup)
		if err != nil {
			log.Println("error on insert mock group : ", err)
			return nil, err
		}

	case false:
		newGroup = &domain.GroupDAO{
			ID:               primitive.NewObjectID(),
			ExhibitionID:     exhibition.ID.Hex(),
			NumUsersRequired: exhibition.NumUsersRequired,
		}

		_, err := ioc.Repo.Groups.Insert(newGroup)
		if err != nil {
			log.Println("error on insert mock group : ", err)
			return nil, err
		}
	}
	return newGroup, nil
}

func PushMockUserToGroup(groupDao *domain.GroupDAO) (*domain.GroupDAO, error) {
	newUser, err := generateMockUser()
	if err != nil {
		return nil, err
	}

	exhibition, err := ioc.Repo.Exhibitions.Get(groupDao.ExhibitionID)
	if err != nil {
		return nil, err
	}

	if time.Now().After(exhibition.StartTime) {
		return nil, fmt.Errorf("ERR504:expired join group time")
	}

	if exhibition.RecruitStartTime.After(newUser.Created) && !exhibition.AllowOldUser {
		return nil, fmt.Errorf("ERR503:not new user")
	}

	if len(groupDao.Users) < exhibition.NumUsersRequired {
		groupDao.Users = append(groupDao.Users, newUser)
		groupDao, err = ioc.Repo.Groups.Update(groupDao)
		if err != nil {
			return nil, err
		}
		newRequest := &domain.GroupRequestDAO{
			ID:           primitive.NewObjectID(),
			UserID:       newUser.ID.Hex(),
			ExhibitionID: groupDao.ExhibitionID,
			GroupID:      groupDao.ID.Hex(),
			Status:       domain.GroupRequestStatusSuccess,
		}
		_, err = ioc.Repo.GroupRequest.Insert(newRequest)
		if err != nil {
			return nil, err
		}
	} else {
		return nil, fmt.Errorf("ERR501:group already completed")
	}
	return groupDao, nil
}

func generateMockUser() (*domain.UserDAO, error) {
	sampleUUID, _ := uuid.NewUUID()
	sampleMobile := utils.CreateMockMobile()

	duplicated := true
	for duplicated {
		user, _ := ioc.Repo.Users.GetByMobile(sampleMobile)
		if user == nil {
			duplicated = false
			break
		}
		sampleMobile = utils.CreateMockMobile()
	}

	newUser := &domain.UserDAO{
		ID:      primitive.NewObjectID(),
		Uuid:    sampleUUID.String(),
		Mobile:  sampleMobile,
		Created: time.Now(),
		Updated: time.Now(),
	}
	_, err := ioc.Repo.Users.Insert(newUser)
	if err != nil {
		log.Println("err on insert mock user : ", err)
		return nil, err
	}
	return newUser, nil
}
