package exhibition

import (
	"github.com/lessbutter/alloff-api/config/ioc"
	"github.com/lessbutter/alloff-api/internal/core/domain"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	"math"
)

func GetCheapestPrice(exhibition *domain.ExhibitionDAO) int {
	cheapestPrice := math.MaxInt64
	for _, pg := range exhibition.ProductGroups {
		for _, pd := range pg.Products {
			// specialPrice 가 있으면 비교대상은 specialPrice 로 한다.
			if pd.Product.SpecialPrice != 0 {
				if cheapestPrice > pd.Product.SpecialPrice {
					cheapestPrice = pd.Product.SpecialPrice
				}
				// specialPrice 가 없으면 비교대상은 discountedPrice 로 한다.
			} else {
				if cheapestPrice > pd.Product.DiscountedPrice {
					cheapestPrice = pd.Product.DiscountedPrice
				}
			}
		}
	}
	return cheapestPrice
}

func UpdateGroupdealInfo(groupDao *domain.GroupDAO, exhibitionDao *domain.ExhibitionDAO) {
	if len(groupDao.Users) == exhibitionDao.NumUsersRequired {
		go generateTicket(groupDao, exhibitionDao)
		go updateTotalParticipants(exhibitionDao)
	}
}

func generateTicket(groupDao *domain.GroupDAO, exhibitionDao *domain.ExhibitionDAO) {
	for _, user := range groupDao.Users {
		groupdealTicket := &domain.GroupdealTicketDAO{
			ID:           primitive.NewObjectID(),
			ExhibitionID: exhibitionDao.ID.Hex(),
			UserID:       user.ID.Hex(),
			Group:        groupDao,
		}
		_, err := ioc.Repo.GroupdealTickets.Insert(groupdealTicket)
		if err != nil {
			log.Println("error occurred on generate groupdeal ticket : ", err)
		}
	}
}

func updateTotalParticipants(exhibitionDao *domain.ExhibitionDAO) {
	exhibitionDao.TotalParticipants += exhibitionDao.NumUsersRequired
	_, err := ioc.Repo.Exhibitions.Upsert(exhibitionDao)
	if err != nil {
		log.Println("error occurred on update total participants : ", err)
	}
}

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
