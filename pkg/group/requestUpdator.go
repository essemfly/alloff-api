package group

import (
	"log"

	"github.com/lessbutter/alloff-api/config/ioc"
	"github.com/lessbutter/alloff-api/internal/core/domain"
)

func CheckRequestPossible(request *domain.GroupRequestDAO) (bool, error) {
	statusFilter := []domain.GroupRequestStatus{domain.GroupRequestStatusPending, domain.GroupRequestStatusSuccess}
	groupRequests, err := ioc.Repo.GroupRequest.ListByGroupID(request.GroupID, statusFilter)
	if err != nil {
		log.Println("error on get group requests data : ", err)
	}
	group, err := ioc.Repo.Groups.Get(request.GroupID)
	if err != nil {
		log.Println("error on get group data : ", err)
	}

	if len(groupRequests) <= group.NumUsersRequired {
		request.Status = domain.GroupRequestStatusSuccess
		_, err := ioc.Repo.GroupRequest.Update(request)
		if err != nil {
			log.Println("error on update group request")
		}
		return true, nil
	} else {
		rank := getRankOfRequest(request, groupRequests)
		if rank <= group.NumUsersRequired {
			request.Status = domain.GroupRequestStatusSuccess
			_, err := ioc.Repo.GroupRequest.Update(request)
			if err != nil {
				log.Println("error on update group request")
			}
			return true, nil
		} else {
			request.Status = domain.GroupRequestStatusFailed
			_, err := ioc.Repo.GroupRequest.Update(request)
			if err != nil {
				log.Println("error on update group request")
			}
			return false, nil
		}
	}
}

func getRankOfRequest(request *domain.GroupRequestDAO, groupRequests []*domain.GroupRequestDAO) (rank int) {
	requestTime := request.CreatedAt
	rank = len(groupRequests) + 1 // 꼴지 부터 시작
	// 내 시간보다 느리게 요청한 시간을 만날때마다 순위를 1씩 올린다.
	for _, gr := range groupRequests {
		if gr.CreatedAt.After(requestTime) {
			rank -= 1
		}
	}
	return
}
