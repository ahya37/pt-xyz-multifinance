package main

import (
	"ahya37/xyz_multifinance/app"
	"ahya37/xyz_multifinance/controller"
	"ahya37/xyz_multifinance/helper"
	"ahya37/xyz_multifinance/middleware"
	"ahya37/xyz_multifinance/model/repository"
	"ahya37/xyz_multifinance/service"
	"net/http"

	_ "github.com/go-sql-driver/mysql"

	"github.com/go-playground/validator/v10"
)

func main() {

	db := app.NewDB()
	validate := validator.New()

	KonsumenRepository := repository.NewKonsumenRepository()
	KonsumenService := service.NewKonsumenService(KonsumenRepository, db, validate)
	KonsumenController := controller.NewKonsumenController(KonsumenService)

	limitKonsumenRepository := repository.NewLimitKonsumenRepository()
	limitKonsumeService := service.NewLimitKonsumenService(limitKonsumenRepository, db, validate)
	limitKonsumenController := controller.NewLimitKonsumenController(limitKonsumeService)

	transaksiRepository := repository.NewTransaksiRepositoryI()
	transaksiService := service.NewTransaksiService(transaksiRepository, db, validate)
	transaksiController := controller.NewTransaksiController(transaksiService)

	router := app.NewRouter(KonsumenController, limitKonsumenController, transaksiController)

	server := http.Server{
		Addr:    "localhost:3000",
		Handler: middleware.NewAuthMiddleware(router),
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
