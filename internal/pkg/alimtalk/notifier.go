package alimtalk

// func NotifyPaymentSuccessAlarm(payment *domain.PaymentDAO) {
// 	order, _ := ioc.Repo.Orders.GetByAlloffID(payment.MerchantUid)
// 	user, _ := ioc.Repo.Users.Get(order.User.ID.Hex())

// 	newAlimtalk := &domain.AlimtalkDAO{
// 		UserID:       user.ID.Hex(),
// 		Mobile:       user.Mobile,
// 		TemplateCode: domain.PAYMENT_OK,
// 		ReferenceID:  strconv.Itoa(payment.ID),
// 		SendDate:     nil,
// 	}

// 	newAlimtalk = newAlimtalk.FillTemplate()
// 	requestId, err := SendMessage(newAlimtalk)
// 	if err != nil {
// 		_, err = newAlimtalk.Add("FAILED")
// 	} else {
// 		newAlimtalk.ToastRequestID = requestId
// 		_, err = newAlimtalk.Add("COMPLETED")
// 	}
// 	if err != nil {
// 		log.Println(err)
// 	}
// }
