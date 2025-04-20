package service

import (
	"context"
)

type UploadService interface {
	UpdateKTPFilename(ctx context.Context, konsumenId int, filename string) error
	UpdateFotoSelfie(ctx context.Context, konsumenId int, filename string) error
}
