package helper

import "log"

func PanicIfError(err error) {
	if err != nil {
		log.Printf("Error memulai transaksi: %v", err)
		panic(err)
	}
}
