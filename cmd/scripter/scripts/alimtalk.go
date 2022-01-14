// package scripter

// import (
// 	"log"

// 	"github.com/lessbutter/alloff-api/config/ioc"
// 	"github.com/lessbutter/alloff-api/internal/core/domain"
// 	"github.com/lessbutter/alloff-api/internal/pkg/alimtalk"
// )

// func testAlimtalk() {

// 	groupID := "612368b287af4834a527aa08"
// 	pg, _ := ioc.Repo.ProductGroups.Get(groupID)

// 	user, _ := ioc.Repo.Users.GetByMobile("01097711882")
// 	newAlimtalk := &domain.AlimtalkDAO{
// 		UserID:       user.ID.Hex(),
// 		Mobile:       user.Mobile,
// 		TemplateCode: domain.TIMEDEAL_OPEN_NOTI,
// 		ReferenceID:  groupID,
// 		SendDate:     &pg.StartTime,
// 	}

// 	newAlimtalk = newAlimtalk.FillTemplate()
// 	requestID, err := alimtalk.SendMessage(newAlimtalk)
// 	if err != nil {
// 		_, err = newAlimtalk.Add("FAILED")
// 	} else {
// 		newAlimtalk.ToastRequestID = requestID
// 		_, err = newAlimtalk.Add("REQUESTED")
// 	}
// 	if err != nil {
// 		log.Println(err)
// 	}

// }
