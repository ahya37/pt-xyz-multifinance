package app

import (
	"ahya37/xyz_multifinance/controller"
	"ahya37/xyz_multifinance/exception"

	"github.com/julienschmidt/httprouter"
)

func NewRouter(KonsumenController controller.KonsumenController, LimitKonsumenController controller.LimitKonsumenController, TransaksiController controller.TransaksiController, UploadController controller.UploadController) *httprouter.Router {
	router := httprouter.New()

	// CRUD data Konsumen
	router.GET("/api/konsumen", KonsumenController.FindAll)
	router.GET("/api/konsumen/:konsumenId", KonsumenController.FindById)
	router.POST("/api/konsumen", KonsumenController.Create)
	router.PUT("/api/konsumen/:konsumenId", KonsumenController.Update)
	router.DELETE("/api/konsumen/:konsumenId", KonsumenController.Delete)
	router.GET("/api/konsumenlimits", KonsumenController.GetLimitKonsumen)

	// Konsumen dan limitnya
	router.POST("/api/limit-konsumen", LimitKonsumenController.Create)

	// Transaksi
	router.POST("/api/transaksi", TransaksiController.Create)
	router.GET("/api/transaksikonsumen", TransaksiController.TransactionKonsumen)

	router.POST("/api/upload/ktp/:konsumenId", UploadController.UploadKTP)
	router.POST("/api/upload/fotoselfie/:konsumenId", UploadController.UploadFotoSelfie)

	router.PanicHandler = exception.ErrorHandler

	return router
}
