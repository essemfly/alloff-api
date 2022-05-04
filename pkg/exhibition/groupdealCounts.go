package exhibition

import (
	"github.com/lessbutter/alloff-api/internal/core/domain"
)

func GetUserParticipatesCount(exhibitions []*domain.ExhibitionDAO, userGroupRequests []*domain.GroupRequestDAO) (cnt int) {
	cnt = 0
	exhibitionIds := []string{}
	userExhibitionIds := []string{}

	for _, exhibition := range exhibitions {
		exhibitionIds = append(exhibitionIds, exhibition.ID.Hex())
	}

	for _, userGroup := range userGroupRequests {
		userExhibitionIds = append(userExhibitionIds, userGroup.ExhibitionID)
	}

	for _, userExhibitionId := range userExhibitionIds {
		for _, exhibitionId := range exhibitionIds {
			if userExhibitionId == exhibitionId {
				cnt += 1
			}
		}
	}
	return
}
