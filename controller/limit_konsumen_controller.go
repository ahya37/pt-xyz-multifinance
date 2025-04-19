package controller

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type LimitKonsumenController interface {
	Create(write http.ResponseWriter, request *http.Request, params httprouter.Params)
}
