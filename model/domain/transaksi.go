package domain

import "time"

type Transaksi struct {
	Id            int
	NoKontrak     string
	OTR           int
	AdminFee      int
	JumlahCicilan int
	JumlahBunga   int
	NamaAset      string
	KonsumenId    int
	CreatedAt     time.Time
	UpdatedAt     time.Time
}
