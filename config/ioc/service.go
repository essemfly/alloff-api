package ioc

import (
	"github.com/lessbutter/alloff-api/internal/core/service"
)

type iocService struct {
	OrderWithPaymentService service.OrderWithPaymentService
}

var Service iocService
