package webhooks

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/lessbutter/alloff-api/config"
	"github.com/lessbutter/alloff-api/config/ioc"
	"github.com/lessbutter/alloff-api/internal/core/domain"
)

type Handler func(w http.ResponseWriter, r *http.Request) error

func (h Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if err := h(w, r); err != nil {
		// handle returned error here.
		http.Error(w, err.Error(), http.StatusBadRequest)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("error occured"))
	}
}

type IamportPaymentResponse struct {
	ImpUID      string `json:"imp_uid"`
	MerchantUID string `json:"merchant_uid"`
	Status      string `json:"status"`
}

func IamportHandler(w http.ResponseWriter, r *http.Request) error {
	if r.Method != "POST" {
		return errors.New("not allowed method")
	}

	var res *IamportPaymentResponse

	err := json.NewDecoder(r.Body).Decode(&res)
	if err != nil {
		return err
	}

	if res.Status != "paid" {
		return nil
	}

	orderDao, err := ioc.Repo.Orders.GetByAlloffID(res.MerchantUID)
	if err != nil {
		config.Logger.Error("ERR301:failed to find order order not found")
		return err
	}

	if orderDao.OrderStatus == domain.ORDER_PAYMENT_FINISHED {
		return nil
	}

	err = ioc.Service.OrderWithPaymentService.VerifyPayment(orderDao, res.ImpUID)
	if err != nil {
		config.Logger.Error("ERR405: failed to verify payment " + err.Error())
		return err
	}

	return nil
}
