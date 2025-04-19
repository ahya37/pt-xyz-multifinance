package service

import (
	"ahya37/xyz_multifinance/model/web"
	"context"
)

type LimitKonsumenService interface {
	Create(ctx context.Context, request web.LimitKonsumenCreateRequest) web.LimitKonsumenResponse
}
