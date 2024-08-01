package ppdb

import (
	"context"
	"database/sql"
	"ppdb-be/pkg/errors"

	// "golang.org/x/crypto/bcrypt"

	// "encoding/json"
	// "log"
	// "strconv"
	// "strings"
	// "time"

	ppdbEntity "ppdb-be/internal/entity/ppdb"
)

// func (d Data) GetKaryawan(ctx context.Context) ([]joEntity.GetKaryawan, error) {
// 	var (
// 		karyawan      joEntity.GetKaryawan
// 		karyawanArray []joEntity.GetKaryawan
// 		err           error
// 	)

// 	rows, err := (*d.stmt)[getKaryawan].QueryxContext(ctx)
// 	if err != nil {
// 		return karyawanArray, errors.Wrap(err, "[DATA] [GetKaryawan]")
// 	}

// 	defer rows.Close()

// 	for rows.Next() {
// 		if err = rows.StructScan(&karyawan); err != nil {
// 			return karyawanArray, errors.Wrap(err, "[DATA] [GetKaryawan]")
// 		}
// 		karyawanArray = append(karyawanArray, karyawan)
// 	}
// 	return karyawanArray, err
// }

// func (d Data) GetCountKaryawan(ctx context.Context) (int, error) {
// 	var (
// 		err   error
// 		total int
// 	)

// 	rows, err := (*d.stmt)[getCountKaryawan].QueryxContext(ctx)
// 	if err != nil {
// 		return total, errors.Wrap(err, "[DATA] [GetCountKaryawan]")
// 	}

// 	defer rows.Close()

// 	for rows.Next() {
// 		if err = rows.Scan(&total); err != nil {
// 			return total, errors.Wrap(err, "[DATA] [GetCountKaryawan]")
// 		}

// 	}
// 	return total, err
// }

// func (d Data) InsertKaryawan(ctx context.Context, karyawan joEntity.GetKaryawan) (string, error) {
// 	var (
// 		err    error
// 		result string
// 	)

// 	_, err = (*d.stmt)[insertKaryawan].ExecContext(ctx,
// 		karyawan.KaryawanID,
// 		karyawan.NamaKaryawan,
// 		karyawan.NoTelp,
// 		karyawan.Keterangan,
// 	)

// 	if err != nil {
// 		result = "Gagal"
// 		return result, errors.Wrap(err, "[DATA][InsertKaryawan]")
// 	}

// 	result = "Berhasil"

// 	return result, err

func (d Data) LoginAdmin(ctx context.Context, admin_id string, admin_password string) (string, error) {
	var (
		admin  ppdbEntity.TableAdmin
		result string
		err    error
	)

	err = (*d.stmt)[loginAdmin].QueryRowxContext(ctx, admin_id).StructScan(&admin)
	if err != nil {
		if err == sql.ErrNoRows {
			result = "Admin not found"
		} else {
			result = "Failed to query admin"
		}
		return result, errors.Wrap(err, "[DATA] [LoginAdmin]")
	}
	

	// err = bcrypt.CompareHashAndPassword([]byte(admin.AdminPassword), []byte(admin_password))
	// if err != nil {
	// 	result = "Invalid password"
	// 	return result, errors.Wrap(err, "[DATA] [LoginAdmin]")
	// }

	result = "Login successful"
	return result, nil
}
