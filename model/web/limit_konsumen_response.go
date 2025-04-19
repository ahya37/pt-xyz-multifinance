package web

type LimitKonsumenResponse struct {
	Id     int    `json:"id"`
	Nik    string `json:"nik"`
	Tenor  int    `json:"tenor"`
	Jumlah int    `json:"jumlah"`
}
