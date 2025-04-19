package web

type KonsumenResponse struct {
	Id           int    `json:"id"`
	Nik          string `json:"nik"`
	FullName     string `json:"full_name"`
	LegalName    string `json:"legal_name"`
	TempatLahir  string `json:"tempat_lahir"`
	TanggalLahir string `json:"tanggal_lahir"`
	Gaji         int    `json:"gaji"`
	FotoKTP      string `json:"foto_ktp"`
	FotoSelfie   string `json:"foto_selfie"`
}

type KonsumenWithLimitResponse struct {
	Nik           string                  `json:"nik"`
	FullName      string                  `json:"full_name"`
	LimitKonsumen []LimitKonsumenResponse `json:"limit_konsumen"`
}

type KonsumenWithTransactionResponse struct {
	Nik       string              `json:"nik"`
	Transaksi []TransaksiResponse `json:"transaksi"`
}
