package alimtalk

import (
	"log"
	"strconv"
	"time"

	"github.com/lessbutter/alloff-api/config"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"

	"github.com/lessbutter/alloff-api/config/ioc"
	"github.com/lessbutter/alloff-api/internal/core/domain"
	"github.com/lessbutter/alloff-api/internal/utils"
)

func NotifyPaymentSuccessAlarm(payment *domain.PaymentDAO) {
	order, _ := ioc.Repo.Orders.GetByAlloffID(payment.MerchantUid)
	user, _ := ioc.Repo.Users.Get(order.User.ID.Hex())

	templateCode := domain.PAYMENT_OK
	for _, item := range order.OrderItems {
		if item.DeliveryDescription.DeliveryType == domain.Foreign {
			templateCode = domain.PAYMENT_OK_OVERSEAS
			break
		}
	}

	newAlimtalk := &domain.AlimtalkDAO{
		UserID:       user.ID.Hex(),
		Mobile:       user.Mobile,
		TemplateCode: templateCode,
		ReferenceID:  strconv.Itoa(payment.ID),
		SendDate:     nil,
		TemplateParams: map[string]string{
			"orderedAt":   utils.TimeFormatterForKorea(order.UpdatedAt),
			"productName": payment.Name,
			"amount":      utils.PriceFormatter(payment.Amount),
		},
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	requestId, err := SendMessage(newAlimtalk)
	if err != nil {
		newAlimtalk.Status = domain.ALIMTALK_STATUS_FAILED
	} else {
		newAlimtalk.ToastRequestID = requestId
		newAlimtalk.Status = domain.ALIMTALK_STATUS_COMPLETED
	}

	_, err = ioc.Repo.Alimtalks.Insert(newAlimtalk)
	if err != nil {
		config.Logger.Error("error on insert alimtalks", zap.Error(err))
	}
}

func NotifyOrderCancelAlarm(orderItem *domain.OrderItemDAO) {
	user, _ := ioc.Repo.Users.Get(orderItem.User.ID.Hex())

	newAlimtalk := &domain.AlimtalkDAO{
		UserID:       user.ID.Hex(),
		Mobile:       user.Mobile,
		TemplateCode: domain.PAYMENT_CANCEL,
		ReferenceID:  strconv.Itoa(orderItem.OrderID),
		SendDate:     nil,
		TemplateParams: map[string]string{
			"취소시간": utils.TimeFormatterForKorea(orderItem.CancelFinishedAt),
			"상품명":  orderItem.ProductName,
			"결제금액": utils.PriceFormatter(orderItem.RefundInfo.RefundAmount),
		},
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	requestId, err := SendMessage(newAlimtalk)
	if err != nil {
		newAlimtalk.Status = domain.ALIMTALK_STATUS_FAILED
	} else {
		newAlimtalk.ToastRequestID = requestId
		newAlimtalk.Status = domain.ALIMTALK_STATUS_COMPLETED
	}

	_, err = ioc.Repo.Alimtalks.Insert(newAlimtalk)
	if err != nil {
		config.Logger.Error("error on insert alimtalks", zap.Error(err))
	}
}

func ChangeExhibitionNotifyStatus(userDao *domain.UserDAO, exhibitionDao *domain.ExhibitionDAO) (*domain.AlimtalkDAO, error) {
	uid := userDao.ID.Hex()
	exId := exhibitionDao.ID.Hex()
	alreadyRegistered, _ := ioc.Repo.Alimtalks.GetByDetail(uid, domain.EXHIBITION_ALARM, exId)
	// TODO: ALIMTALK_STATUS_READY 인 친구들은 메시지 보내진지 어케 알지 ?

	// 이미 등록된 알림톡이 있으며, 그 알림톡이 취소되지않고 전송 대기중인 경우
	// 메시지 발송을 취소하고, 알림톡 상태를 STATUS_CANCELED로 바꾼다음 persist
	if alreadyRegistered != nil && alreadyRegistered.Status == domain.ALIMTALK_STATUS_READY {
		// TODO: Check for delete message works
		err := DeleteMessage(alreadyRegistered)
		if err != nil {
			log.Println("delete message error", err)
		}
		alreadyRegistered.Status = domain.ALIMTALK_STATUS_CANCLED
		_, err = ioc.Repo.Alimtalks.Update(alreadyRegistered)
		if err != nil {
			return nil, err
		}
		return nil, nil

		// 이미 등록된 알림톡이 있으며, 그 알림톡이 취소되거나 발송에 실패했던 경우
		// 메시지 발송을 다시 등록하고, 알림톡 상태를 STATUS_READY로 바꾼다음 persist
	} else if alreadyRegistered != nil && (alreadyRegistered.Status == domain.ALIMTALK_STATUS_CANCLED || alreadyRegistered.Status == domain.ALIMTALK_STATUS_FAILED) {
		requestId, err := SendMessage(alreadyRegistered)
		if err != nil {
			alreadyRegistered.Status = domain.ALIMTALK_STATUS_FAILED
		} else {
			alreadyRegistered.ToastRequestID = requestId
			alreadyRegistered.Status = domain.ALIMTALK_STATUS_READY
		}

		alreadyRegistered.TemplateParams = map[string]string{
			"createdAt":      utils.TimeFormatterForKorea(time.Now().Add(time.Hour * 9)),
			"exhibitionName": exhibitionDao.Title,
		}
		alreadyRegistered.UpdatedAt = time.Now()

		_, err = ioc.Repo.Alimtalks.Update(alreadyRegistered)
		if err != nil {
			config.Logger.Error("error on update alimtalks", zap.Error(err))
			return nil, err
		}
		return alreadyRegistered, nil
	}

	// 이미 등록된 알림톡이 있어도, 그게 취소된 상태이거나
	// 이미 등록된 알림톡이 없으면 새로운 알림톡을 만든다.
	newAlimtalk := &domain.AlimtalkDAO{
		ID:           primitive.NewObjectID(),
		UserID:       uid,
		Mobile:       userDao.Mobile,
		TemplateCode: domain.DEAL_OPEN,
		ReferenceID:  exId,
		TemplateParams: map[string]string{
			"createdAt":      utils.TimeFormatterForKorea(time.Now().Add(time.Hour * 9)),
			"exhibitionName": exhibitionDao.Title,
		},
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	requestId, err := SendMessage(newAlimtalk)
	if err != nil {
		newAlimtalk.Status = domain.ALIMTALK_STATUS_FAILED
	} else {
		newAlimtalk.ToastRequestID = requestId
		newAlimtalk.Status = domain.ALIMTALK_STATUS_READY
	}

	_, err = ioc.Repo.Alimtalks.Insert(newAlimtalk)
	if err != nil {
		config.Logger.Error("error on insert alimtalks", zap.Error(err))
		return nil, err
	}

	return newAlimtalk, nil
}
