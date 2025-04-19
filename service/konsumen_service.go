package service

import (
	"ahya37/xyz_multifinance/model/web"
	"context"
)

type KonsumenService interface {
	Create(ctx context.Context, request web.KonsumenCreateRequest) web.KonsumenResponse
	Update(ctx context.Context, request web.KonsumenUpdateRequest) web.KonsumenResponse
	Delete(ctx context.Context, konsumenId int)
	FindById(ctx context.Context, konsumenId int) web.KonsumenResponse
	FindAll(ctx context.Context) []web.KonsumenResponse
	GetLimitKonsumen(ctx context.Context) []web.KonsumenWithLimitResponse
}
