package controller

import (
	"ahya37/xyz_multifinance/helper"
	"ahya37/xyz_multifinance/model/web"
	"ahya37/xyz_multifinance/service"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type LimitKonsumenControllerImpl struct {
	LimitKonsumenService service.LimitKonsumenService
}

func NewLimitKonsumenController(limitKonsumenService service.LimitKonsumenService) LimitKonsumenController {
	return &LimitKonsumenControllerImpl{
		LimitKonsumenService: limitKonsumenService,
	}
}

func (controller *LimitKonsumenControllerImpl) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	limitKonsumenRequest := web.LimitKonsumenCreateRequest{}
	helper.RaedFromRequestBody(request, &limitKonsumenRequest)

	limitKonsumenResponse := controller.LimitKonsumenService.Create(request.Context(), limitKonsumenRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   limitKonsumenResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}
