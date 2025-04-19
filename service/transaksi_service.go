package service

import (
	"ahya37/xyz_multifinance/model/web"
	"context"
)

type TransaksiService interface {
	Create(ctx context.Context, request web.TransaksiCreateRequest) web.TransaksiResponse
	TransactionKonsumen(ctx context.Context) []web.KonsumenWithTransactionResponse
}
