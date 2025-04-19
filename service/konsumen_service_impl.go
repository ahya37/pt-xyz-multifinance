package service

import (
	"ahya37/xyz_multifinance/exception"
	"ahya37/xyz_multifinance/helper"
	"ahya37/xyz_multifinance/model/domain"
	"ahya37/xyz_multifinance/model/repository"
	"ahya37/xyz_multifinance/model/web"
	"context"
	"database/sql"
	"time"

	"github.com/go-playground/validator/v10"
)

type KonsumenServiceImpl struct {
	KonsumenRepository repository.KonsumenRepository
	DB                 *sql.DB
	Validate           *validator.Validate
}

func NewKonsumenService(KonsumenRepository repository.KonsumenRepository, DB *sql.DB, validate *validator.Validate) KonsumenService {
	return &KonsumenServiceImpl{
		KonsumenRepository: KonsumenRepository,
		DB:                 DB,
		Validate:           validate,
	}
}

func (service *KonsumenServiceImpl) Create(ctx context.Context, request web.KonsumenCreateRequest) web.KonsumenResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	tanggalLahir, err := time.Parse("2006-01-02", request.TanggalLahir)
	helper.PanicIfError(err)

	customer := domain.Konsumen{
		Nik:          request.Nik,
		FullName:     request.FullName,
		LegalName:    request.LegalName,
		TempatLahir:  request.TempatLahir,
		TanggalLahir: tanggalLahir,
		Gaji:         request.Gaji,
		FotoKTP:      request.FotoKTP,
		FotoSelfie:   request.FotoSelfie,
	}

	customer = service.KonsumenRepository.Save(ctx, tx, customer)

	return helper.TokonsumenResponse(customer)

}

func (service *KonsumenServiceImpl) Update(ctx context.Context, request web.KonsumenUpdateRequest) web.KonsumenResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	tanggalLahir, err := time.Parse("2006-01-02", request.TanggalLahir)
	helper.PanicIfError(err)

	customer, err := service.KonsumenRepository.FindById(ctx, tx, request.Id)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}
	customer.Nik = request.Nik
	customer.FullName = request.FullName
	customer.LegalName = request.LegalName
	customer.TempatLahir = request.TempatLahir
	customer.TanggalLahir = tanggalLahir
	customer.Gaji = request.Gaji
	customer.FotoKTP = request.FotoKTP
	customer.FotoSelfie = request.FotoSelfie

	customer = service.KonsumenRepository.Update(ctx, tx, customer)

	return helper.TokonsumenResponse(customer)
}

func (service *KonsumenServiceImpl) Delete(ctx context.Context, customerId int) {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	customer, err := service.KonsumenRepository.FindById(ctx, tx, customerId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	service.KonsumenRepository.Delete(ctx, tx, customer)
}

func (service *KonsumenServiceImpl) FindById(ctx context.Context, customerId int) web.KonsumenResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	customer, err := service.KonsumenRepository.FindById(ctx, tx, customerId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return helper.TokonsumenResponse(customer)
}

func (service *KonsumenServiceImpl) FindAll(ctx context.Context) []web.KonsumenResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	konsumen := service.KonsumenRepository.FindAll(ctx, tx)

	return helper.ToKonsumenResponses(konsumen)
}

func (service *KonsumenServiceImpl) GetLimitKonsumen(ctx context.Context) []web.KonsumenWithLimitResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	konsumen := service.KonsumenRepository.GetLimitKonsumen(ctx, tx)

	return helper.ToKonsumenResponsesWithLimit(konsumen)
}
