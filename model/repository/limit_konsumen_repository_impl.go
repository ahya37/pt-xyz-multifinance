package repository

import (
	"ahya37/xyz_multifinance/helper"
	"ahya37/xyz_multifinance/model/domain"
	"context"
	"database/sql"
)

type LimitKonsumenRepositoryImpl struct {
}

func NewLimitKonsumenRepository() LimitKonsumenRepository {
	return &LimitKonsumenRepositoryImpl{}
}

func (repository *LimitKonsumenRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, limitKonsumen domain.LimitKonsumen) domain.LimitKonsumen {
	SQL := "insert into limit_konsumen (nik,tenor,jumlah) values(?,?,?)"
	result, err := tx.ExecContext(ctx, SQL,
		limitKonsumen.Nik,
		limitKonsumen.Tenor,
		limitKonsumen.Jumlah)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)

	limitKonsumen.Id = int(id)
	return limitKonsumen
}
