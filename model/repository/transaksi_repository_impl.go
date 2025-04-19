package repository

import (
	"ahya37/xyz_multifinance/helper"
	"ahya37/xyz_multifinance/model/domain"
	"context"
	"database/sql"
	"time"
)

type TransaksiRepositoryImpl struct {
}

func NewTransaksiRepositoryI() TransaksiRepository {
	return &TransaksiRepositoryImpl{}
}

func (repository *TransaksiRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, transaksi domain.Transaksi) domain.Transaksi {
	SQL := `INSERT INTO transaksi(no_kontrak, otr, admin_fee, jumlah_cicilan, jumlah_bunga, nama_aset, konsumen_id, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)`
	result, err := tx.ExecContext(ctx, SQL, transaksi.NoKontrak, transaksi.OTR, transaksi.AdminFee, transaksi.JumlahCicilan, transaksi.JumlahBunga, transaksi.NamaAset, transaksi.KonsumenId, time.Now(), time.Now())
	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)

	transaksi.Id = int(id)
	return transaksi
}

func (repository *TransaksiRepositoryImpl) TransactionKonsumen(ctx context.Context, tx *sql.Tx) []domain.Konsumen {
	SQL := `SELECT
			b.nik,
			a.id, 
			a.no_kontrak, 
			a.otr, 
			a.admin_fee, 
			a.jumlah_cicilan, 
			a.jumlah_bunga, 
			a.nama_aset,
			a.konsumen_id,
			a.created_at,
			a.updated_at
			from transaksi as a
			join konsumen as b on b.id = a.konsumen_id`

	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(err)
	defer rows.Close()

	transaksiMap := make(map[string]domain.Konsumen)

	for rows.Next() {
		var konsumen domain.Konsumen
		var transaksi domain.Transaksi
		var CreatedAt []byte
		var UpdatedAt []byte

		err := rows.Scan(
			&konsumen.Nik,
			&transaksi.Id,
			&transaksi.NoKontrak,
			&transaksi.OTR,
			&transaksi.AdminFee,
			&transaksi.JumlahCicilan,
			&transaksi.JumlahBunga,
			&transaksi.NamaAset,
			&transaksi.KonsumenId,
			&CreatedAt,
			&UpdatedAt,
		)

		helper.PanicIfError(err)

		if len(CreatedAt) > 0 {
			transaksi.CreatedAt, err = time.Parse("2006-01-02 15:04:05", string(CreatedAt))
			helper.PanicIfError(err)
		}

		if len(UpdatedAt) > 0 {
			transaksi.UpdatedAt, err = time.Parse("2006-01-02 15:04:05", string(UpdatedAt))
			helper.PanicIfError(err)
		}

		if _, ok := transaksiMap[konsumen.Nik]; !ok {
			konsumen.Transaksi = []domain.Transaksi{}
			transaksiMap[konsumen.Nik] = konsumen
		}
		cust := transaksiMap[konsumen.Nik]
		if transaksi.Id != 0 {
			cust.Transaksi = append(cust.Transaksi, transaksi)
			transaksiMap[konsumen.Nik] = cust
		}
	}

	result := make([]domain.Konsumen, 0, len(transaksiMap))
	for _, v := range transaksiMap {
		result = append(result, v)
	}
	return result
}
