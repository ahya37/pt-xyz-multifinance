package repository

import (
	"ahya37/xyz_multifinance/model/domain"
	"context"
	"database/sql"
)

type LimitKonsumenRepository interface {
	Save(ctx context.Context, tx *sql.Tx, limitKonsumen domain.LimitKonsumen) domain.LimitKonsumen
}
