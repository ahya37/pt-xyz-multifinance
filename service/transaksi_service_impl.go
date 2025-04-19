package service

import (
	"ahya37/xyz_multifinance/helper"
	"ahya37/xyz_multifinance/model/domain"
	"ahya37/xyz_multifinance/model/repository"
	"ahya37/xyz_multifinance/model/web"
	"context"
	"database/sql"

	"github.com/go-playground/validator/v10"
)

type TransaksiServiceImpl struct {
	TransaksiRepository repository.TransaksiRepository
	DB                  *sql.DB
	Validate            *validator.Validate
}

func NewTransaksiService(transaksiRepository repository.TransaksiRepository, DB *sql.DB, validate *validator.Validate) TransaksiService {
	return &TransaksiServiceImpl{
		TransaksiRepository: transaksiRepository,
		DB:                  DB,
		Validate:            validate,
	}
}

func (service *TransaksiServiceImpl) Create(ctx context.Context, request web.TransaksiCreateRequest) web.TransaksiResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	transaksi := domain.Transaksi{
		NoKontrak:     request.NoKontrak,
		OTR:           request.OTR,
		AdminFee:      request.AdminFee,
		JumlahCicilan: request.JumlahCicilan,
		JumlahBunga:   request.JumlahBunga,
		NamaAset:      request.NamaAset,
		KonsumenId:    request.KonsumenId,
	}

	transaksi = service.TransaksiRepository.Save(ctx, tx, transaksi)
	return helper.ToTransaksiResponse(transaksi)

}

func (service *TransaksiServiceImpl) TransactionKonsumen(ctx context.Context) []web.KonsumenWithTransactionResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	transaksi := service.TransaksiRepository.TransactionKonsumen(ctx, tx)

	return helper.ToTransaksiWithKonsumenResponses(transaksi)

}
