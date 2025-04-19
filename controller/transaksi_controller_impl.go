package controller

import (
	"ahya37/xyz_multifinance/helper"
	"ahya37/xyz_multifinance/model/web"
	"ahya37/xyz_multifinance/service"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type TransaksiControllerImpl struct {
	TransaksiService service.TransaksiService
}

func NewTransaksiController(transaksiService service.TransaksiService) TransaksiController {
	return &TransaksiControllerImpl{
		TransaksiService: transaksiService,
	}
}

func (controller *TransaksiControllerImpl) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	transaksiRequest := web.TransaksiCreateRequest{}
	helper.RaedFromRequestBody(request, &transaksiRequest)

	transaksiResponse := controller.TransaksiService.Create(request.Context(), transaksiRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   transaksiResponse,
	}
	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *TransaksiControllerImpl) TransactionKonsumen(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	TransaksiResponse := controller.TransaksiService.TransactionKonsumen(request.Context())
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   TransaksiResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)

}
