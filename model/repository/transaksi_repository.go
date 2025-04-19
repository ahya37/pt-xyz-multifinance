package repository

import (
	"ahya37/xyz_multifinance/model/domain"
	"context"
	"database/sql"
)

type TransaksiRepository interface {
	Save(ctx context.Context, tx *sql.Tx, transaksi domain.Transaksi) domain.Transaksi
	TransactionKonsumen(ctx context.Context, tx *sql.Tx) []domain.Konsumen
}
