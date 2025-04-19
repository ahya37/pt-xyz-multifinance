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

type LimitKonsumenServiceImpl struct {
	LimitKonsumenRepository repository.LimitKonsumenRepository
	DB                      *sql.DB
	Validate                *validator.Validate
}

func NewLimitKonsumenService(limitKonsumenRepository repository.LimitKonsumenRepository, DB *sql.DB, validate *validator.Validate) LimitKonsumenService {
	return &LimitKonsumenServiceImpl{
		LimitKonsumenRepository: limitKonsumenRepository,
		DB:                      DB,
		Validate:                validate,
	}
}

func (service *LimitKonsumenServiceImpl) Create(ctx context.Context, request web.LimitKonsumenCreateRequest) web.LimitKonsumenResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	limitKonsumen := domain.LimitKonsumen{
		Nik:    request.Nik,
		Tenor:  request.Tenor,
		Jumlah: request.Jumlah,
	}

	limitKonsumen = service.LimitKonsumenRepository.Save(ctx, tx, limitKonsumen)
	return helper.ToLimitKonsumenResponse(limitKonsumen)

}
