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

	"golang.org/x/crypto/bcrypt"
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

func (d Data) LoginAdmin(ctx context.Context, emailAdmin string, password string) (string, error) {
	var (
		admin  ppdbEntity.TableAdmin
		result string
		err    error
	)

	// Query untuk mendapatkan data admin berdasarkan emailAdmin
	err = (*d.stmt)[loginAdmin].QueryRowxContext(ctx, emailAdmin).StructScan(&admin)
	if err != nil {
		if err == sql.ErrNoRows {
			result = "Admin not found"
		} else {
			result = "Failed to query admin"
		}
		return result, errors.Wrap(err, "[DATA] [LoginAdmin]")
	}

	// Pengecekan hardcoded untuk email "catherinbella38@gmail.com" dengan password "admin"
	if admin.EmailAdmin == "catherinbella38@gmail.com" && password == "admin" {
		result = "Login successful"
		return result, nil
	}

	// Proses hashing untuk akun lainnya
	err = bcrypt.CompareHashAndPassword([]byte(admin.Password), []byte(password))
	if err != nil {
		result = "Invalid password"
		return result, errors.Wrap(err, "[DATA] [LoginAdmin] Password mismatch")
	}

	result = "Login successful"
	return result, nil
}

// Kontak Sekolah
func (d Data) GetKontakSekolah(ctx context.Context) ([]ppdbEntity.TableKontakSekolah, error) {
	var (
		KontakSekolah      ppdbEntity.TableKontakSekolah
		kontakSekolahArray []ppdbEntity.TableKontakSekolah
		err                error
	)

	rows, err := (*d.stmt)[getKontakSekolah].QueryxContext(ctx)
	if err != nil {
		return kontakSekolahArray, errors.Wrap(err, "[DATA] [GetKontakSekolah]")
	}

	defer rows.Close()

	for rows.Next() {
		if err = rows.StructScan(&KontakSekolah); err != nil {
			return kontakSekolahArray, errors.Wrap(err, "[DATA] [GetKontakSekolah]")
		}
		kontakSekolahArray = append(kontakSekolahArray, KontakSekolah)
	}
	return kontakSekolahArray, err
}

// Get Data Admin
func (d Data) GetDataAdmin(ctx context.Context, searchInput string) ([]ppdbEntity.TableKelolaDataAdmin, error) {
	var (
		adminData      ppdbEntity.TableKelolaDataAdmin
		adminDataArray []ppdbEntity.TableKelolaDataAdmin
		err            error
	)

	// Execute the query with search input
	rows, err := (*d.stmt)[getDataAdmin].QueryxContext(ctx, "%"+searchInput+"%")
	if err != nil {
		return adminDataArray, errors.Wrap(err, "[DATA] [GetDataAdmin] Error executing query")
	}

	defer rows.Close()

	// Loop through the rows and map to struct
	for rows.Next() {
		if err = rows.StructScan(&adminData); err != nil {
			return adminDataArray, errors.Wrap(err, "[DATA] [GetDataAdmin] Error scanning row")
		}
		adminDataArray = append(adminDataArray, adminData)
	}

	// Return the result
	return adminDataArray, err
}
