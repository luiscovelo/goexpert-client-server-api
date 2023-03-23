package controller

import (
	"net/http"

	"github.com/luiscovelo/goexpert-client-server-api/server/rest"
	"github.com/luiscovelo/goexpert-client-server-api/server/service"
)

type quotationController struct {
	service *service.QuotationService
}

func NewQuotationController(service *service.QuotationService) *quotationController {
	return &quotationController{
		service: service,
	}
}

func (ctrl *quotationController) GetQuotation(w http.ResponseWriter, r *http.Request) {
	quotation, err := ctrl.service.GetQuotationAndSave()

	if err != nil {
		rest.NewBadRequest(w, err)
		return
	}

	rest.NewSuccessful(w, quotation)
}
