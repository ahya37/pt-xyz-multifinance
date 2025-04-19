package repository

import (
	"ahya37/xyz_multifinance/helper"
	"ahya37/xyz_multifinance/model/domain"
	"context"
	"database/sql"
	"errors"
	"time"
)

type KonsumenRepositoryImpl struct {
}

func NewKonsumenRepository() KonsumenRepository {
	return &KonsumenRepositoryImpl{}
}

func (repository *KonsumenRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, konsumen domain.Konsumen) domain.Konsumen {
	SQL := "insert into konsumen(nik,full_name,legal_name,tempat_lahir,tanggal_lahir,gaji,foto_ktp,foto_selfie) value (?,?,?,?,?,?,?,?)"
	result, err := tx.ExecContext(ctx, SQL,
		konsumen.Nik,
		konsumen.FullName,
		konsumen.LegalName,
		konsumen.TempatLahir,
		konsumen.TanggalLahir,
		konsumen.Gaji,
		konsumen.FotoKTP,
		konsumen.FotoSelfie)
	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)

	konsumen.Id = int(id)
	return konsumen
}

func (repository *KonsumenRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, konsumen domain.Konsumen) domain.Konsumen {
	SQL := "update konsumen set nik = ?, full_name = ?, legal_name = ?, tempat_lahir = ?, tanggal_lahir = ?, gaji = ?, foto_ktp = ?, foto_selfie = ? where id = ?"
	_, err := tx.ExecContext(ctx, SQL,
		konsumen.Nik,
		konsumen.FullName,
		konsumen.LegalName,
		konsumen.TempatLahir,
		konsumen.TanggalLahir,
		konsumen.Gaji,
		konsumen.FotoKTP,
		konsumen.FotoSelfie,
		konsumen.Id)
	helper.PanicIfError(err)

	return konsumen
}

func (repository *KonsumenRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, konsumen domain.Konsumen) {
	SQL := "delete from konsumen where id = ?"
	_, err := tx.ExecContext(ctx, SQL, konsumen.Id)
	helper.PanicIfError(err)
}

func (repository *KonsumenRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, konsumenId int) (domain.Konsumen, error) {
	SQL := "SELECT id, nik, full_name, legal_name, tempat_lahir, tanggal_lahir, gaji, foto_ktp, foto_selfie from konsumen where id = ?"
	rows, err := tx.QueryContext(ctx, SQL, konsumenId)
	helper.PanicIfError(err)
	defer rows.Close()

	konsumen := domain.Konsumen{}
	if rows.Next() {
		var tanggalLahir string
		err := rows.Scan(&konsumen.Id, &konsumen.Nik, &konsumen.FullName, &konsumen.LegalName, &konsumen.TempatLahir, &tanggalLahir, &konsumen.Gaji, &konsumen.FotoKTP, &konsumen.FotoSelfie)
		helper.PanicIfError(err)

		tanggal, err := time.Parse("2006-01-02", tanggalLahir)
		helper.PanicIfError(err)

		konsumen.TanggalLahir = tanggal
		return konsumen, nil
	} else {
		return konsumen, errors.New("konsumen is not found")

	}
}
func (repository *KonsumenRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.Konsumen {
	SQL := "SELECT id, nik, full_name, legal_name, tempat_lahir, tanggal_lahir, gaji, foto_ktp, foto_selfie from konsumen"
	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(err)
	defer rows.Close()

	var konsumens []domain.Konsumen
	for rows.Next() {
		konsumen := domain.Konsumen{}
		var tanggalLahir string
		err := rows.Scan(
			&konsumen.Id,
			&konsumen.Nik,
			&konsumen.FullName,
			&konsumen.LegalName,
			&konsumen.TempatLahir,
			&tanggalLahir,
			&konsumen.Gaji,
			&konsumen.FotoKTP,
			&konsumen.FotoSelfie)
		helper.PanicIfError(err)
		tanggal, err := time.Parse("2006-01-02", tanggalLahir)
		helper.PanicIfError(err)
		konsumen.TanggalLahir = tanggal
		konsumens = append(konsumens, konsumen)
	}

	return konsumens

}

func (repository *KonsumenRepositoryImpl) GetLimitKonsumen(ctx context.Context, tx *sql.Tx) []domain.Konsumen {
	SQL := `
		SELECT 
            k.id,
            k.nik,
            k.full_name,
            k.legal_name,
            k.tempat_lahir,
            k.gaji,
            k.foto_ktp,
            k.foto_selfie,
            lk.id as limit_id,
            lk.tenor,
            lk.jumlah
        FROM konsumen k
        LEFT JOIN limit_konsumen lk ON k.nik = lk.nik
	`
	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(err)
	defer rows.Close()

	konsumenMap := make(map[string]domain.Konsumen)

	for rows.Next() {
		var konsumen domain.Konsumen
		var limitKonsumen domain.LimitKonsumen

		err := rows.Scan(
			&konsumen.Id,
			&konsumen.Nik,
			&konsumen.FullName,
			&konsumen.LegalName,
			&konsumen.TempatLahir,
			&konsumen.Gaji,
			&konsumen.FotoKTP,
			&konsumen.FotoSelfie,
			&limitKonsumen.Id,
			&limitKonsumen.Tenor,
			&limitKonsumen.Jumlah,
		)

		helper.PanicIfError(err)

		if _, ok := konsumenMap[konsumen.Nik]; !ok {
			konsumen.LimitKonsumen = []domain.LimitKonsumen{}
			konsumenMap[konsumen.Nik] = konsumen
		}
		if limitKonsumen.Id != 0 {
			cust := konsumenMap[konsumen.Nik]
			cust.LimitKonsumen = append(cust.LimitKonsumen, limitKonsumen)
			konsumenMap[konsumen.Nik] = cust
		}
	}

	result := make([]domain.Konsumen, 0, len(konsumenMap))
	for _, v := range konsumenMap {
		result = append(result, v)
	}
	return result
}
