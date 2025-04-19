package web

import "time"

type TransaksiResponse struct {
	Id            int       `json:"id"`
	NoKontrak     string    `json:"no_kontrak"`
	OTR           int       `json:"otr"`
	AdminFee      int       `json:"admin_fee"`
	JumlahCicilan int       `json:"jumlah_cicilan"`
	JumlahBunga   int       `json:"jumlah_bunga"`
	NamaAset      string    `json:"nama_aset"`
	KonsumenId    int       `json:"konsumen_id"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}
