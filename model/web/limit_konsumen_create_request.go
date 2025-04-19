package web

type LimitKonsumenCreateRequest struct {
	Nik    string `validate:"required,max=16,min=0" json:"nik"`
	Tenor  int    `validate:"required" json:"tenor"`
	Jumlah int    `validate:"required" json:"jumlah"`
}
