package repository

import (
	"ahya37/xyz_multifinance/model/domain"
	"context"
	"database/sql"
)

type KonsumenRepository interface {
	Save(ctx context.Context, tx *sql.Tx, konsumen domain.Konsumen) domain.Konsumen
	Update(ctx context.Context, tx *sql.Tx, konsumen domain.Konsumen) domain.Konsumen
	Delete(ctx context.Context, tx *sql.Tx, konsumen domain.Konsumen)
	FindById(ctx context.Context, tx *sql.Tx, konsumenId int) (domain.Konsumen, error)
	FindAll(ctx context.Context, tx *sql.Tx) []domain.Konsumen
	GetLimitKonsumen(ctx context.Context, tx *sql.Tx) []domain.Konsumen
}
