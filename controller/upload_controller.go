package controller

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type UploadController interface {
	UploadKTP(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	UploadFotoSelfie(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
}
