package domain

import (
	"time"
)

type Konsumen struct {
	Id            int
	Nik           string
	FullName      string
	LegalName     string
	TempatLahir   string
	TanggalLahir  time.Time
	Gaji          int
	FotoKTP       string
	FotoSelfie    string
	LimitKonsumen []LimitKonsumen
	Transaksi     []Transaksi
}
