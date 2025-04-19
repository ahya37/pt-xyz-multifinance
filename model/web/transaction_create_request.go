package web

import "time"

type TransaksiCreateRequest struct {
	NoKontrak     string `validate:"required" json:"no_kontrak"`
	OTR           int    `validate:"required" json:"otr"`
	AdminFee      int    `validate:"required" json:"admin_fee"`
	JumlahCicilan int    `validate:"required" json:"jumlah_cicilan"`
	JumlahBunga   int    `validate:"required" json:"jumlah_bunga"`
	NamaAset      string `validate:"required" json:"nama_aset"`
	KonsumenId    int    `validate:"required" json:"konsumen_id"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
}
