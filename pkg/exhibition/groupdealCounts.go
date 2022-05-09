package exhibition

import (
	"github.com/lessbutter/alloff-api/internal/core/domain"
)

// GetUserPendingGroupCount : 유저의 그룹 참여 기록과 오픈 대기중인 그룹딜을 비교하여 대기중인 그룹딜의 갯수를 구한다.
func GetUserPendingGroupCount(exhibitions []*domain.ExhibitionDAO, userGroupRequests []*domain.GroupRequestDAO) (cnt int) {
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

// GetUserPurchasableGroupCount : 유저가 보유한 입장권 티켓과 라이브 중인 그룹딜을 비교하여 구매 가능한 그룹딜의 갯수를 구s다.
func GetUserPurchasableGroupCount(exhibitions []*domain.ExhibitionDAO, tickets []*domain.GroupdealTicketDAO) (cnt int) {
	cnt = 0
	exhibitionIds := []string{}
	ticketExhibitionIds := []string{}

	for _, exhibition := range exhibitions {
		exhibitionIds = append(exhibitionIds, exhibition.ID.Hex())
	}

	for _, ticket := range tickets {
		ticketExhibitionIds = append(ticketExhibitionIds, ticket.ExhibitionID)
	}

	for _, ticketExhibitionId := range ticketExhibitionIds {
		for _, exhibitionId := range exhibitionIds {
			if ticketExhibitionId == exhibitionId {
				cnt += 1
			}
		}
	}
	return
}
