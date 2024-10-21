package ppdb

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"path/filepath"
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

func saveImageToFile(imageBytes []byte, filePath string) error {
	file, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("failed to create file: %w", err)
	}
	defer file.Close()

	_, err = file.Write(imageBytes)
	if err != nil {
		return fmt.Errorf("failed to write image to file: %w", err)
	}

	return nil
}

func EnsureDirectory(path string) error {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		if err := os.MkdirAll(path, os.ModePerm); err != nil {
			return err
		}
	}
	return nil
}

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

// Info Pendaftarn
func (d Data) InsertInfoDaftar(ctx context.Context, infoDaftar ppdbEntity.TableInfoDaftar) (string, error) {
	var (
		err    error
		result string
		lastID string
		newID  string
	)

	// Mengambil InfoID terakhir
	err = (*d.stmt)[getLastInfoId].QueryRowxContext(ctx).Scan(&lastID)
	if err != nil && err != sql.ErrNoRows {
		result = "Gagal mengambil ID terakhir"
		return result, errors.Wrap(err, "[DATA][GetLastInfoId]")
	}

	if lastID != "" {
		// Mengambil bagian numerik dari lastID dan menambahkannya
		num, _ := strconv.Atoi(lastID[1:])
		newID = fmt.Sprintf("I%04d", num+1)

	} else {
		newID = "I0001" // ID pertama
	}

	fmt.Println("newID", newID)

	// Set AdminID baru ke admin struct
	infoDaftar.InfoID = newID

	// Eksekusi query untuk memasukkan data ke dalam tabel T_InfoDaftar
	_, err = (*d.stmt)[insertInfoDaftar].ExecContext(ctx,
		infoDaftar.InfoID,
		infoDaftar.PosterDaftar,
		infoDaftar.AwalTahunAjar,
		infoDaftar.AkhirTahunAjar,
	)

	if err != nil {
		result = "Gagal"
		return result, errors.Wrap(err, "[DATA][InsertDataAdmin]")
	}

	result = "Berhasil"
	return result, nil
}

func (d Data) GetGambarInfoDaftar(ctx context.Context, infoID string) ([]byte, error) {
	var poster []byte
	if err := (*d.stmt)[getGambarInfoDaftar].QueryRowxContext(ctx, infoID).Scan(&poster); err != nil {
		return poster, errors.Wrap(err, "[DATA][GetGambarInfoDaftar]")
	}

	return poster, nil
}

func generateImageURLFotoInfoDaftar(id string) string {
	var url = "http://localhost:8081"
	return fmt.Sprintf(url+"/ppdb/v1/data/getgambarinfodaftar?infoID=%s", id)
}

func (d Data) GetInfoDaftar(ctx context.Context) ([]ppdbEntity.TableInfoDaftar, error) {
	var (
		infoDaftarArray []ppdbEntity.TableInfoDaftar
		err             error
	)

	rows, err := (*d.stmt)[getInfoDaftar].QueryxContext(ctx)
	if err != nil {
		return infoDaftarArray, errors.Wrap(err, "[DATA] [GetInfoDaftar]")
	}

	defer rows.Close()

	// Pastikan direktori untuk menyimpan gambar ada
	imageDir := filepath.Join("public", "images")
	if err := EnsureDirectory(imageDir); err != nil {
		return nil, errors.Wrap(err, "[DATA] [GetInfoDaftar] - Failed to ensure directory")
	}

	for rows.Next() {
		var infoDaftar ppdbEntity.TableInfoDaftar

		// Memindahkan data dari hasil query ke dalam struct
		if err = rows.Scan(&infoDaftar.InfoID, &infoDaftar.LinkPosterDaftar, &infoDaftar.AwalTahunAjar, &infoDaftar.AkhirTahunAjar); err != nil {
			return nil, errors.Wrap(err, "[DATA] [GetInfoDaftar] - Failed to scan row")
		}

		// Menyimpan gambar dan menghasilkan URL
		filePath := filepath.Join(imageDir, infoDaftar.InfoID+".jpg")
		if err := saveImageToFile(infoDaftar.PosterDaftar, filePath); err != nil {
			return nil, errors.Wrap(err, "[DATA] [GetInfoDaftar] - Failed to save image")
		}

		infoDaftar.LinkPosterDaftar = generateImageURLFotoInfoDaftar(infoDaftar.InfoID)
		infoDaftarArray = append(infoDaftarArray, infoDaftar)
	}

	if err = rows.Err(); err != nil {
		return nil, errors.Wrap(err, "[DATA] [GetInfoDaftar] - Row iteration error")
	}

	return infoDaftarArray, nil
}

// Banner sekolah
func (d Data) InsertBanner(ctx context.Context, banner ppdbEntity.TableBanner) (string, error) {
	var (
		err    error
		result string
		lastID string
		newID  string
	)

	err = (*d.stmt)[getLastBannerId].QueryRowxContext(ctx).Scan(&lastID)
	if err != nil && err != sql.ErrNoRows {
		result = "Gagal mengambil ID terakhir"
		return result, errors.Wrap(err, "[DATA][GetLastBannerId]")
	}

	if lastID != "" {
		num, _ := strconv.Atoi(lastID[1:])
		newID = fmt.Sprintf("B%04d", num+1)

	} else {
		newID = "B0001" // ID pertama
	}

	fmt.Println("newID", newID)

	// Set AdminID baru ke admin struct
	banner.BannerID = newID

	// Eksekusi query untuk memasukkan data ke dalam tabel T_InfoDaftar
	_, err = (*d.stmt)[insertBanner].ExecContext(ctx,
		banner.BannerID,
		banner.BannerName,
		banner.BannerImage,
	)

	if err != nil {
		result = "Gagal"
		return result, errors.Wrap(err, "[DATA][InsertBanner]")
	}

	result = "Berhasil"
	return result, nil
}

func (d Data) GetGambarBanner(ctx context.Context, bannerID string) ([]byte, error) {
	var poster []byte
	if err := (*d.stmt)[getGambarBanner].QueryRowxContext(ctx, bannerID).Scan(&poster); err != nil {
		return poster, errors.Wrap(err, "[DATA][GetGambarBanner]")
	}

	return poster, nil
}

func generateImageURLFotoBanner(id string) string {
	var url = "http://localhost:8081"
	return fmt.Sprintf(url+"/ppdb/v1/data/getgambarbanner?bannerID=%s", id)
}

func (d Data) GetBanner(ctx context.Context) ([]ppdbEntity.TableBanner, error) {
	var (
		bannerArray []ppdbEntity.TableBanner
		err         error
	)

	rows, err := (*d.stmt)[getBanner].QueryxContext(ctx)
	if err != nil {
		return bannerArray, errors.Wrap(err, "[DATA] [GetBanner]")
	}

	defer rows.Close()

	// Pastikan direktori untuk menyimpan gambar ada
	imageDir := filepath.Join("public", "images")
	if err := EnsureDirectory(imageDir); err != nil {
		return nil, errors.Wrap(err, "[DATA] [GetBanner] - Failed to ensure directory")
	}

	for rows.Next() {
		var banner ppdbEntity.TableBanner

		// Memindahkan data dari hasil query ke dalam struct
		if err = rows.Scan(&banner.BannerID, &banner.BannerName, &banner.LinkBannerImage); err != nil {
			return nil, errors.Wrap(err, "[DATA] [GetBanner] - Failed to scan row")
		}

		// Menyimpan gambar dan menghasilkan URL
		filePath := filepath.Join(imageDir, banner.BannerID+".jpg")
		if err := saveImageToFile(banner.BannerImage, filePath); err != nil {
			return nil, errors.Wrap(err, "[DATA] [GetBanner] - Failed to save image")
		}

		banner.LinkBannerImage = generateImageURLFotoBanner(banner.BannerID)
		bannerArray = append(bannerArray, banner)
	}

	if err = rows.Err(); err != nil {
		return nil, errors.Wrap(err, "[DATA] [GetBanner] - Row iteration error")
	}

	return bannerArray, nil
}

// Hapus Data Banner
func (d Data) DeleteBanner(ctx context.Context, bannerID string) (string, error) {
	var (
		err    error
		result string
	)

	_, err = (*d.stmt)[deleteBanner].ExecContext(ctx, bannerID)

	if err != nil {
		result = "Gagal"
		return result, errors.Wrap(err, "[DATA][DeleteBanner]")
	}

	result = "Berhasil"
	return result, nil
}

// Fasilitas Sekolah
func (d Data) InsertFasilitas(ctx context.Context, fasilitas ppdbEntity.TableFasilitas) (string, error) {
	var (
		err    error
		result string
		lastID string
		newID  string
	)

	err = (*d.stmt)[getLastFasilitasId].QueryRowxContext(ctx).Scan(&lastID)
	if err != nil && err != sql.ErrNoRows {
		result = "Gagal mengambil ID terakhir"
		return result, errors.Wrap(err, "[DATA][GetLastFasilitasId]")
	}

	if lastID != "" {
		num, _ := strconv.Atoi(lastID[1:])
		newID = fmt.Sprintf("F%04d", num+1)

	} else {
		newID = "F0001" // ID pertama
	}

	fmt.Println("newID", newID)

	// Set AdminID baru ke admin struct
	fasilitas.FasilitasID = newID

	// Eksekusi query untuk memasukkan data ke dalam tabel T_Fasilitas
	_, err = (*d.stmt)[insertFasilitas].ExecContext(ctx,
		fasilitas.FasilitasID,
		fasilitas.FasilitasName,
		fasilitas.FasilitasImage,
	)

	if err != nil {
		result = "Gagal"
		return result, errors.Wrap(err, "[DATA][InsertFasilitas]")
	}

	result = "Berhasil"
	return result, nil
}

func (d Data) GetGambarFasilitas(ctx context.Context, fasilitasID string) ([]byte, error) {
	var poster []byte
	if err := (*d.stmt)[getGambarFasilitas].QueryRowxContext(ctx, fasilitasID).Scan(&poster); err != nil {
		return poster, errors.Wrap(err, "[DATA][GetGambarFasilitas]")
	}

	return poster, nil
}

func generateImageURLFotoFasilitas(id string) string {
	var url = "http://localhost:8081"
	return fmt.Sprintf(url+"/ppdb/v1/data/getgambarfasilitas?fasilitasID=%s", id)
}

//ini untuk get fasilitas di website admin (ada untuk paginationnya)
func (d Data) GetFasilitas(ctx context.Context, searchInput string, offset, limit int) ([]ppdbEntity.TableFasilitas, error) {
	var (
		fasilitasArray []ppdbEntity.TableFasilitas
		err            error
	)

	rows, err := (*d.stmt)[getFasilitas].QueryxContext(ctx, "%"+searchInput+"%", offset, limit)
	if err != nil {
		return fasilitasArray, errors.Wrap(err, "[DATA] [GetFasilitas]")
	}
	defer rows.Close()

	// Pastikan direktori untuk menyimpan gambar ada
	imageDir := filepath.Join("public", "images")
	if err := EnsureDirectory(imageDir); err != nil {
		return nil, errors.Wrap(err, "[DATA] [GetFasilitas] - Failed to ensure directory")
	}

	for rows.Next() {
		var fasilitas ppdbEntity.TableFasilitas

		// Memindahkan data dari hasil query ke dalam struct
		if err = rows.Scan(&fasilitas.FasilitasID, &fasilitas.FasilitasName, &fasilitas.LinkFasilitasImage); err != nil {
			return nil, errors.Wrap(err, "[DATA] [GetFasilitas] - Failed to scan row")
		}

		// Menyimpan gambar dan menghasilkan URL
		filePath := filepath.Join(imageDir, fasilitas.FasilitasID+".jpg")
		if err := saveImageToFile(fasilitas.FasilitasImage, filePath); err != nil {
			return nil, errors.Wrap(err, "[DATA] [GetFasilitas] - Failed to save image")
		}

		fasilitas.LinkFasilitasImage = generateImageURLFotoFasilitas(fasilitas.FasilitasID)
		fasilitasArray = append(fasilitasArray, fasilitas)
	}

	if err = rows.Err(); err != nil {
		return nil, errors.Wrap(err, "[DATA] [GetFasilitas] - Row iteration error")
	}

	return fasilitasArray, nil
}

func (d Data) GetFasilitasPagination(ctx context.Context, searchInput string) (int, error) {
	var totalCount int

	// Query untuk mendapatkan total count tanpa LIMIT
	err := (*d.stmt)[getFasilitasPagination].GetContext(ctx, &totalCount, "%"+searchInput+"%")
	if err != nil {
		return 0, errors.Wrap(err, "[DATA] [GetFasilitasPagination] Error executing count query")
	}

	return totalCount, nil
}

//ini untuk get fasilitas di website utama 
func (d Data) GetFasilitasUtama(ctx context.Context) ([]ppdbEntity.TableFasilitas, error) {
	var (
		fasilitasArray []ppdbEntity.TableFasilitas
		err         error
	)

	rows, err := (*d.stmt)[getFasilitasUtama].QueryxContext(ctx)
	if err != nil {
		return fasilitasArray, errors.Wrap(err, "[DATA] [GetFasilitasUtama]")
	}

	defer rows.Close()

	// Pastikan direktori untuk menyimpan gambar ada
	imageDir := filepath.Join("public", "images")
	if err := EnsureDirectory(imageDir); err != nil {
		return nil, errors.Wrap(err, "[DATA] [GetFasilitasUtama] - Failed to ensure directory")
	}

	for rows.Next() {
		var fasilitas ppdbEntity.TableFasilitas

		// Memindahkan data dari hasil query ke dalam struct
		if err = rows.Scan(&fasilitas.FasilitasID, &fasilitas.FasilitasName, &fasilitas.LinkFasilitasImage); err != nil {
			return nil, errors.Wrap(err, "[DATA] [GetFasilitasUtama] - Failed to scan row")
		}

		// Menyimpan gambar dan menghasilkan URL
		filePath := filepath.Join(imageDir, fasilitas.FasilitasID+".jpg")
		if err := saveImageToFile(fasilitas.FasilitasImage, filePath); err != nil {
			return nil, errors.Wrap(err, "[DATA] [GetFasilitasUtama] - Failed to save image")
		}

		fasilitas.LinkFasilitasImage = generateImageURLFotoFasilitas(fasilitas.FasilitasID)
		fasilitasArray = append(fasilitasArray, fasilitas)
	}

	if err = rows.Err(); err != nil {
		return nil, errors.Wrap(err, "[DATA] [GetFasilitasUtama] - Row iteration error")
	}

	return fasilitasArray, nil
}



// Hapus Data Fasilitas
func (d Data) DeleteFasilitas(ctx context.Context, fasilitasID string) (string, error) {
	var (
		err    error
		result string
	)

	_, err = (*d.stmt)[deleteFasilitas].ExecContext(ctx, fasilitasID)

	if err != nil {
		result = "Gagal"
		return result, errors.Wrap(err, "[DATA][DeleteFasilitas]")
	}

	result = "Berhasil"
	return result, nil
}

// Profile Staff
func (d Data) InsertProfileStaff(ctx context.Context, staff ppdbEntity.TableStaff) (string, error) {
    var (
        err    error
        result string
        lastID string
        newID  string
    )

    // Mengambil StaffID terakhir
    err = (*d.stmt)[getLastStaffId].QueryRowxContext(ctx).Scan(&lastID)
    if err != nil && err != sql.ErrNoRows {
        result = "Gagal mengambil ID terakhir"
        return result, errors.Wrap(err, "[DATA][GetLastStaffId]")
    }

    if lastID != "" {
        // Mengambil bagian numerik dari lastID dan menambahkannya
        num, _ := strconv.Atoi(lastID[1:])
        newID = fmt.Sprintf("S%04d", num+1)
    } else {
        newID = "S0001" // ID pertama
    }

    fmt.Println("newID", newID)

    // Set StaffID baru ke struct staff
    staff.StaffID = newID

    // Eksekusi query untuk memasukkan data ke dalam tabel T_ProfileStaff
    _, err = (*d.stmt)[insertProfileStaff].ExecContext(ctx,
        staff.StaffID,
        staff.StaffName,
        staff.StaffGender,
        staff.StaffPosition,
        staff.StaffTmptLahir,
        staff.StaffTglLahir,
        staff.StaffPhoto,
    )

    if err != nil {
        result = "Gagal menyimpan data staff"
        return result, errors.Wrap(err, "[DATA][InsertProfileStaff]")
    }

    result = "Berhasil menyimpan data staff"
    return result, nil
}

func (d Data) GetPhotoStaff(ctx context.Context, staffID string) ([]byte, error) {
	var poster []byte
	if err := (*d.stmt)[getPhotoStaff].QueryRowxContext(ctx, staffID).Scan(&poster); err != nil {
		return poster, errors.Wrap(err, "[DATA][GetPhotoStaff]")
	}

	return poster, nil
}


