package helper

import (
	"ahya37/xyz_multifinance/model/domain"
	"ahya37/xyz_multifinance/model/web"
)

func TokonsumenResponse(konsumen domain.Konsumen) web.KonsumenResponse {
	return web.KonsumenResponse{
		Id:           konsumen.Id,
		Nik:          konsumen.Nik,
		FullName:     konsumen.FullName,
		LegalName:    konsumen.LegalName,
		TempatLahir:  konsumen.TempatLahir,
		TanggalLahir: konsumen.TanggalLahir.Format("2006-01-02"),
		Gaji:         konsumen.Gaji,
		FotoKTP:      konsumen.FotoKTP,
		FotoSelfie:   konsumen.FotoSelfie,
	}
}

func ToKonsumenResponses(konsumens []domain.Konsumen) []web.KonsumenResponse {
	var konsumenResponses []web.KonsumenResponse
	for _, konsumens := range konsumens {
		konsumenResponses = append(konsumenResponses, TokonsumenResponse(konsumens))
	}

	return konsumenResponses
}

func ToKonsumenWithLimitResponse(konsumen domain.Konsumen) web.KonsumenWithLimitResponse {
	var limitResponses []web.LimitKonsumenResponse
	for _, limit := range konsumen.LimitKonsumen {
		limitResponses = append(limitResponses, web.LimitKonsumenResponse{
			Id:     limit.Id,
			Tenor:  limit.Tenor,
			Jumlah: limit.Jumlah,
		})
	}

	return web.KonsumenWithLimitResponse{
		Nik:           konsumen.Nik,
		FullName:      konsumen.FullName,
		LimitKonsumen: limitResponses,
	}
}

func ToKonsumenResponsesWithLimit(konsumens []domain.Konsumen) []web.KonsumenWithLimitResponse {
	var konsumenResponses []web.KonsumenWithLimitResponse
	for _, konsumen := range konsumens {
		limitResponses := make([]web.LimitKonsumenResponse, len(konsumen.LimitKonsumen))
		for i, limit := range konsumen.LimitKonsumen {
			limitResponses[i] = web.LimitKonsumenResponse{
				Id:     limit.Id,
				Tenor:  limit.Tenor,
				Jumlah: limit.Jumlah,
			}
		}
		konsumenResponses = append(konsumenResponses, web.KonsumenWithLimitResponse{
			Nik:           konsumen.Nik,
			FullName:      konsumen.FullName,
			LimitKonsumen: limitResponses,
		})
	}
	return konsumenResponses
}

func ToTransaksiWithKonsumenResponses(konsumens []domain.Konsumen) []web.KonsumenWithTransactionResponse {
	var transaksiResponses []web.KonsumenWithTransactionResponse
	for _, konsumen := range konsumens {
		transaksiKonsumenResponses := make([]web.TransaksiResponse, len(konsumen.Transaksi))
		for i, trx := range konsumen.Transaksi {
			transaksiKonsumenResponses[i] = web.TransaksiResponse{
				Id:            trx.Id,
				NoKontrak:     trx.NoKontrak,
				OTR:           trx.OTR,
				AdminFee:      trx.AdminFee,
				JumlahCicilan: trx.JumlahCicilan,
				JumlahBunga:   trx.JumlahBunga,
				NamaAset:      trx.NamaAset,
				KonsumenId:    trx.KonsumenId,
				CreatedAt:     trx.CreatedAt,
				UpdatedAt:     trx.UpdatedAt,
			}
		}
		transaksiResponses = append(transaksiResponses, web.KonsumenWithTransactionResponse{
			Nik:       konsumen.Nik,
			Transaksi: transaksiKonsumenResponses,
		})
	}

	return transaksiResponses
}

func ToLimitKonsumenResponse(limitKonsumen domain.LimitKonsumen) web.LimitKonsumenResponse {
	return web.LimitKonsumenResponse{
		Id:     limitKonsumen.Id,
		Nik:    limitKonsumen.Nik,
		Jumlah: limitKonsumen.Jumlah,
	}
}

func ToTransaksiResponse(transaksi domain.Transaksi) web.TransaksiResponse {
	return web.TransaksiResponse{
		Id:            transaksi.Id,
		NoKontrak:     transaksi.NoKontrak,
		OTR:           transaksi.OTR,
		AdminFee:      transaksi.AdminFee,
		JumlahCicilan: transaksi.JumlahCicilan,
		JumlahBunga:   transaksi.JumlahBunga,
		NamaAset:      transaksi.NamaAset,
		KonsumenId:    transaksi.KonsumenId,
		CreatedAt:     transaksi.CreatedAt,
		UpdatedAt:     transaksi.UpdatedAt,
	}
}
