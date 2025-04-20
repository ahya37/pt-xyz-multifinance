package service

import (
	"ahya37/xyz_multifinance/helper"
	"ahya37/xyz_multifinance/model/repository"
	"context"
	"database/sql"
)

type UploadServiceImpl struct {
	KonsumenRepository repository.KonsumenRepository
	DB                 *sql.DB
}

func NewUploadService(konsumenRepository repository.KonsumenRepository, DB *sql.DB) *UploadServiceImpl {
	return &UploadServiceImpl{
		KonsumenRepository: konsumenRepository,
		DB:                 DB,
	}
}

func (service *UploadServiceImpl) UpdateKTPFilename(ctx context.Context, konsumenId int, filename string) error {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	konsumen, err := service.KonsumenRepository.FindById(ctx, tx, konsumenId)
	helper.PanicIfError(err)

	konsumen.FotoKTP = filename
	konsumen = service.KonsumenRepository.Update(ctx, tx, konsumen)
	return err

}

func (service *UploadServiceImpl) UpdateFotoSelfie(ctx context.Context, konsumenId int, filename string) error {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	konsumen, err := service.KonsumenRepository.FindById(ctx, tx, konsumenId)
	helper.PanicIfError(err)

	konsumen.FotoKTP = filename
	konsumen = service.KonsumenRepository.Update(ctx, tx, konsumen)
	return err

}
