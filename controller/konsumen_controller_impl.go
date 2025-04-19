package controller

import (
	"ahya37/xyz_multifinance/helper"
	"ahya37/xyz_multifinance/model/web"
	"ahya37/xyz_multifinance/service"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

type KonsumenControllerImpl struct {
	KonsumenService service.KonsumenService
}

func NewKonsumenController(KonsumenService service.KonsumenService) KonsumenController {
	return &KonsumenControllerImpl{
		KonsumenService: KonsumenService,
	}
}

func (controller *KonsumenControllerImpl) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	KonsumenCreateRequest := web.KonsumenCreateRequest{}
	helper.RaedFromRequestBody(request, &KonsumenCreateRequest)

	KonsumenResponse := controller.KonsumenService.Create(request.Context(), KonsumenCreateRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   KonsumenResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)

}

func (controller *KonsumenControllerImpl) Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	KonsumenUpdateRequest := web.KonsumenUpdateRequest{}
	helper.RaedFromRequestBody(request, &KonsumenUpdateRequest)

	konsumenId := params.ByName("konsumenId")
	id, err := strconv.Atoi(konsumenId)
	helper.PanicIfError(err)

	KonsumenUpdateRequest.Id = id

	KonsumenResponse := controller.KonsumenService.Update(request.Context(), KonsumenUpdateRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   KonsumenResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)

}

func (controller *KonsumenControllerImpl) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {

	konsumenId := params.ByName("konsumenId")
	id, err := strconv.Atoi(konsumenId)
	helper.PanicIfError(err)

	controller.KonsumenService.Delete(request.Context(), id)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
	}

	helper.WriteToResponseBody(writer, webResponse)

}

func (controller *KonsumenControllerImpl) FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	konsumenId := params.ByName("konsumenId")
	id, err := strconv.Atoi(konsumenId)
	helper.PanicIfError(err)

	KonsumenResponse := controller.KonsumenService.FindById(request.Context(), id)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   KonsumenResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)

}

func (controller *KonsumenControllerImpl) FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	KonsumenResponse := controller.KonsumenService.FindAll(request.Context())
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   KonsumenResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)

}

func (controller *KonsumenControllerImpl) GetLimitKonsumen(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	KonsumenResponse := controller.KonsumenService.GetLimitKonsumen(request.Context())
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   KonsumenResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)

}
