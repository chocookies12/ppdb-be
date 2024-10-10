package ppdb

import (
	"context"
	"database/sql"
	"fmt"
	"ppdb-be/pkg/errors"
	"strconv"

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
func (d Data) GetDataAdmin(ctx context.Context, searchInput string, offset, limit int) ([]ppdbEntity.TableKelolaDataAdmin, error) {
	var (
		adminData      ppdbEntity.TableKelolaDataAdmin
		adminDataArray []ppdbEntity.TableKelolaDataAdmin
		err            error
	)

	// Execute the query with search input, limit, and offset
	rows, err := (*d.stmt)[getDataAdmin].QueryxContext(ctx, "%"+searchInput+"%", offset, limit)
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

func (d Data) GetDataAdminPagination(ctx context.Context, searchInput string) (int, error) {
    var totalCount int

    // Query untuk mendapatkan total count tanpa LIMIT
    err := (*d.stmt)[getDataAdminPagination].GetContext(ctx, &totalCount, "%"+searchInput+"%")
    if err != nil {
        return 0, errors.Wrap(err, "[DATA] [GetDataAdminPagination] Error executing count query")
    }

    return totalCount, nil
}


func (d Data) InsertDataAdmin(ctx context.Context, admin ppdbEntity.TableAdmin) (string, error) {
	var (
		err    error
		result string
		lastID string
		newID  string
	)

	// Mengambil AdminID terakhir
	err = (*d.stmt)[getLastAdminId].QueryRowxContext(ctx).Scan(&lastID)
	if err != nil && err != sql.ErrNoRows {
		result = "Gagal mengambil ID terakhir"
		return result, errors.Wrap(err, "[DATA][GetLastAdminId]")
	}

	if lastID != "" {
		// Mengambil bagian numerik dari lastID dan menambahkannya
		num, _ := strconv.Atoi(lastID[1:])
		newID = fmt.Sprintf("A%04d", num+1) // Menggunakan prefix "A" untuk AdminID

	} else {
		newID = "A0001" // ID pertama
	}

	fmt.Println("newID", newID)

	// Set AdminID baru ke admin struct
	admin.AdminID = newID

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(admin.Password), bcrypt.DefaultCost)
	if err != nil {
		return "Gagal", errors.Wrap(err, "[DATA][InsertDataAdmin] - hashing password failed")
	}

	// Set the hashed password back to the admin struct
	admin.Password = string(hashedPassword)

	// Eksekusi query untuk memasukkan data ke dalam tabel T_Admin
	_, err = (*d.stmt)[insertDataAdmin].ExecContext(ctx,
		admin.AdminID,
		admin.RoleID,
		admin.AdminName,
		admin.Password,
		admin.EmailAdmin,
	)

	if err != nil {
		result = "Gagal"
		return result, errors.Wrap(err, "[DATA][InsertDataAdmin]")
	}

	result = "Berhasil"
	return result, nil
}

// Hapus Data Admin
func (d Data) DeleteAdmin(ctx context.Context, adminID string) (string, error) {
    var (
        err    error
        result string
    )
	
    _, err = (*d.stmt)[deleteDataAdmin].ExecContext(ctx, adminID)

    if err != nil {
        result = "Gagal"
        return result, errors.Wrap(err, "[DATA][DeleteAdmin]")
    }

    result = "Berhasil"
    return result, nil
}


// Role
func (d Data) GetRole(ctx context.Context) ([]ppdbEntity.TableRole, error) {
	var (
		role      ppdbEntity.TableRole
		roleArray []ppdbEntity.TableRole
		err       error
	)

	rows, err := (*d.stmt)[getRole].QueryxContext(ctx)
	if err != nil {
		return roleArray, errors.Wrap(err, "[DATA] [GetRole]")
	}

	defer rows.Close()

	for rows.Next() {
		if err = rows.StructScan(&role); err != nil {
			return roleArray, errors.Wrap(err, "[DATA] [GetRole]")
		}
		roleArray = append(roleArray, role)
	}
	return roleArray, err
}
