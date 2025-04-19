package web

type KonsumenCreateRequest struct {
	Nik          string `validate:"required,max=16,min=0" json:"nik"`
	FullName     string `validate:"required,max=255,min=0" json:"full_name"`
	LegalName    string `validate:"required,max=255,min=0" json:"legal_name"`
	TempatLahir  string `validate:"required" json:"tempat_lahir"`
	TanggalLahir string `validate:"required" json:"tanggal_lahir" `
	Gaji         int    `validate:"required" json:"gaji"`
	FotoKTP      string `validate:"required" json:"foto_ktp"`
	FotoSelfie   string `validate:"required" json:"foto_selfie"`
}
