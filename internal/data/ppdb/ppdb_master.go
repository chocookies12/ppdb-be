package ppdb

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"path/filepath"
	"ppdb-be/pkg/errors"
	"strconv"
	"time"

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

// Agama
func (d Data) GetAgama(ctx context.Context) ([]ppdbEntity.TableAgama, error) {
	var (
		agama      ppdbEntity.TableAgama
		agamaArray []ppdbEntity.TableAgama
		err        error
	)

	rows, err := (*d.stmt)[getAgama].QueryxContext(ctx)
	if err != nil {
		return agamaArray, errors.Wrap(err, "[DATA] [GetAgama]")
	}

	defer rows.Close()

	for rows.Next() {
		if err = rows.StructScan(&agama); err != nil {
			return agamaArray, errors.Wrap(err, "[DATA] [GetAgama]")
		}
		agamaArray = append(agamaArray, agama)
	}
	return agamaArray, err
}

// Jurusan
func (d Data) GetJurusan(ctx context.Context) ([]ppdbEntity.TableJurusan, error) {
	var (
		jurusan      ppdbEntity.TableJurusan
		jurusanArray []ppdbEntity.TableJurusan
		err          error
	)

	rows, err := (*d.stmt)[getJurusan].QueryxContext(ctx)
	if err != nil {
		return jurusanArray, errors.Wrap(err, "[DATA] [GetJurusan]")
	}

	defer rows.Close()

	for rows.Next() {
		if err = rows.StructScan(&jurusan); err != nil {
			return jurusanArray, errors.Wrap(err, "[DATA] [GetJurusan]")
		}
		jurusanArray = append(jurusanArray, jurusan)
	}
	return jurusanArray, err
}

// Status
func (d Data) GetStatus(ctx context.Context) ([]ppdbEntity.TableStatus, error) {
	var (
		status      ppdbEntity.TableStatus
		statusArray []ppdbEntity.TableStatus
		err         error
	)

	rows, err := (*d.stmt)[getStatus].QueryxContext(ctx)
	if err != nil {
		return statusArray, errors.Wrap(err, "[DATA] [GetStatus]")
	}

	defer rows.Close()

	for rows.Next() {
		if err = rows.StructScan(&status); err != nil {
			return statusArray, errors.Wrap(err, "[DATA] [GetStatus]")
		}
		statusArray = append(statusArray, status)
	}
	return statusArray, err
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

func (d Data) UpdateInfoDaftar(ctx context.Context, infoDaftar ppdbEntity.TableInfoDaftar, infoID string) (string, error) {
	var (
		result string
		err    error
	)

	_, err = (*d.stmt)[updateInfoDaftar].ExecContext(ctx, infoDaftar.PosterDaftar, infoDaftar.AwalTahunAjar, infoDaftar.AkhirTahunAjar, infoID)

	if err != nil {
		result = "Gagal"
		return result, errors.Wrap(err, "[DATA][UpdateInfoDaftar]")
	}

	result = "Berhasil"
	return result, err
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

func (d Data) UpdateBanner(ctx context.Context, banner ppdbEntity.TableBanner, bannerID string) (string, error) {
	var (
		result string
		err    error
	)

	_, err = (*d.stmt)[updateBanner].ExecContext(ctx, banner.BannerName, banner.BannerImage, bannerID)

	if err != nil {
		result = "Gagal"
		return result, errors.Wrap(err, "[DATA][UpdateBanner]")
	}

	result = "Berhasil"
	return result, err
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

// ini untuk get fasilitas di website admin (ada untuk paginationnya)
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

// ini untuk get fasilitas di website utama
func (d Data) GetFasilitasUtama(ctx context.Context) ([]ppdbEntity.TableFasilitas, error) {
	var (
		fasilitasArray []ppdbEntity.TableFasilitas
		err            error
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

func (d Data) UpdateFasilitas(ctx context.Context, fasilitas ppdbEntity.TableFasilitas, fasilitasID string) (string, error) {
	var (
		result string
		err    error
	)

	_, err = (*d.stmt)[updateFasilitas].ExecContext(ctx, fasilitas.FasilitasName, fasilitas.FasilitasImage, fasilitasID)

	if err != nil {
		result = "Gagal"
		return result, errors.Wrap(err, "[DATA][UpdateFasilitas]")
	}

	result = "Berhasil"
	return result, err
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

func generateImageURLFotoStaff(id string) string {
	var url = "http://localhost:8081"
	return fmt.Sprintf(url+"/ppdb/v1/data/getphotostaff?staffID=%s", id)
}

// ini untuk get staff di website admin (ada untuk paginationnya)
func (d Data) GetProfileStaff(ctx context.Context, searchInput string, offset, limit int) ([]ppdbEntity.TableStaff, error) {
	var (
		staffArray []ppdbEntity.TableStaff
		err        error
	)

	rows, err := (*d.stmt)[getProfilStaff].QueryxContext(ctx, "%"+searchInput+"%", offset, limit)
	if err != nil {
		return staffArray, errors.Wrap(err, "[DATA] [GetProfileStaff]")
	}
	defer rows.Close()

	// Pastikan direktori untuk menyimpan gambar ada
	imageDir := filepath.Join("public", "images")
	if err := EnsureDirectory(imageDir); err != nil {
		return nil, errors.Wrap(err, "[DATA] [GetProfileStaff] - Failed to ensure directory")
	}

	for rows.Next() {
		var staff ppdbEntity.TableStaff

		// Memindahkan data dari hasil query ke dalam struct
		var staffTglLahir sql.NullString // Menggunakan sql.NullString untuk menangani nilai NULL dari database
		if err = rows.Scan(&staff.StaffID, &staff.StaffName, &staff.StaffGender, &staff.StaffPosition, &staff.StaffTmptLahir, &staffTglLahir, &staff.LinkStaffPhoto); err != nil {
			return nil, errors.Wrap(err, "[DATA] [GetProfileStaff] - Failed to scan row")
		}

		// Mengubah staffTglLahir menjadi *time.Time
		if staffTglLahir.Valid {
			t, err := time.Parse("2006-01-02", staffTglLahir.String) // Menggunakan format yang sesuai
			if err != nil {
				return nil, errors.Wrap(err, "[DATA] [GetProfileStaff] - Failed to parse staffTglLahir")
			}
			staff.StaffTglLahir = &t
		} else {
			staff.StaffTglLahir = nil // Mengatur nil jika tidak ada nilai
		}

		// Menyimpan gambar dan menghasilkan URL
		filePath := filepath.Join(imageDir, staff.StaffID+".jpg")
		if err := saveImageToFile(staff.StaffPhoto, filePath); err != nil {
			return nil, errors.Wrap(err, "[DATA] [GetProfileStaff] - Failed to save image")
		}

		// Menghasilkan URL untuk foto staf
		staff.LinkStaffPhoto = generateImageURLFotoStaff(staff.StaffID)

		staffArray = append(staffArray, staff)
	}

	if err = rows.Err(); err != nil {
		return nil, errors.Wrap(err, "[DATA] [GetProfileStaff] - Row iteration error")
	}

	return staffArray, nil
}

func (d Data) GetProfileStaffPagination(ctx context.Context, searchInput string) (int, error) {
	var totalCount int

	// Query untuk mendapatkan total count tanpa LIMIT
	err := (*d.stmt)[getProfilStaffPagination].GetContext(ctx, &totalCount, "%"+searchInput+"%")
	if err != nil {
		return 0, errors.Wrap(err, "[DATA] [GetProfileStaffPagination] Error executing count query")
	}

	return totalCount, nil
}

func (d Data) DeleteProfileStaff(ctx context.Context, staffID string) (string, error) {
	var (
		err    error
		result string
	)

	_, err = (*d.stmt)[deleteProfileStaff].ExecContext(ctx, staffID)

	if err != nil {
		result = "Gagal"
		return result, errors.Wrap(err, "[DATA][DeleteProfileStaff]")
	}

	result = "Berhasil"
	return result, nil
}

func (d Data) GetProfileStaffUtama(ctx context.Context) ([]ppdbEntity.TableStaff, error) {
	var (
		staffArray []ppdbEntity.TableStaff
		err        error
	)

	rows, err := (*d.stmt)[getProfilStaffUtama].QueryxContext(ctx)
	if err != nil {
		return staffArray, errors.Wrap(err, "[DATA] [GetProfileStaff]")
	}
	defer rows.Close()

	// Pastikan direktori untuk menyimpan gambar ada
	imageDir := filepath.Join("public", "images")
	if err := EnsureDirectory(imageDir); err != nil {
		return nil, errors.Wrap(err, "[DATA] [GetProfileStaff] - Failed to ensure directory")
	}

	for rows.Next() {
		var staff ppdbEntity.TableStaff

		// Memindahkan data dari hasil query ke dalam struct
		var staffTglLahir sql.NullString // Menggunakan sql.NullString untuk menangani nilai NULL dari database
		if err = rows.Scan(&staff.StaffID, &staff.StaffName, &staff.StaffGender, &staff.StaffPosition, &staff.StaffTmptLahir, &staffTglLahir, &staff.LinkStaffPhoto); err != nil {
			return nil, errors.Wrap(err, "[DATA] [GetProfileStaff] - Failed to scan row")
		}

		// Mengubah staffTglLahir menjadi *time.Time
		if staffTglLahir.Valid {
			t, err := time.Parse("2006-01-02", staffTglLahir.String) // Menggunakan format yang sesuai
			if err != nil {
				return nil, errors.Wrap(err, "[DATA] [GetProfileStaff] - Failed to parse staffTglLahir")
			}
			staff.StaffTglLahir = &t
		} else {
			staff.StaffTglLahir = nil // Mengatur nil jika tidak ada nilai
		}

		// Menyimpan gambar dan menghasilkan URL
		filePath := filepath.Join(imageDir, staff.StaffID+".jpg")
		if err := saveImageToFile(staff.StaffPhoto, filePath); err != nil {
			return nil, errors.Wrap(err, "[DATA] [GetProfileStaff] - Failed to save image")
		}

		// Menghasilkan URL untuk foto staf
		staff.LinkStaffPhoto = generateImageURLFotoStaff(staff.StaffID)
		staffArray = append(staffArray, staff)
	}

	if err = rows.Err(); err != nil {
		return nil, errors.Wrap(err, "[DATA] [GetProfileStaff] - Row iteration error")
	}

	return staffArray, nil
}

func (d Data) UpdateProfileStaff(ctx context.Context, staff ppdbEntity.TableStaff, staffID string) (string, error) {
	var (
		result string
		err    error
	)
	fmt.Println("masuk data")

	_, err = (*d.stmt)[updateProfileStaff].ExecContext(ctx,
		staff.StaffName,
		staff.StaffGender,
		staff.StaffPosition,
		staff.StaffTmptLahir,
		staff.StaffTglLahir,
		staff.StaffPhoto,
		staffID)
	fmt.Println("masuk data2")

	if err != nil {
		result = "Gagal"
		return result, errors.Wrap(err, "[DATA][UpdateProfileStaff]")
	}
	fmt.Println("masuk data3")

	result = "Berhasil"
	return result, err
}

// EVENT
func (d Data) InsertEvent(ctx context.Context, event ppdbEntity.TableEvent) (string, error) {
	var (
		err    error
		result string
		lastID string
		newID  string
	)

	// Mengambil EventID terakhir
	err = (*d.stmt)[getLastEventId].QueryRowxContext(ctx).Scan(&lastID)
	if err != nil && err != sql.ErrNoRows {
		result = "Gagal mengambil ID terakhir"
		return result, errors.Wrap(err, "[DATA][GetLastEventId]")
	}

	if lastID != "" {
		// Mengambil bagian numerik dari lastID dan menambahkannya
		num, _ := strconv.Atoi(lastID[1:])
		newID = fmt.Sprintf("E%04d", num+1)
	} else {
		newID = "E0001" // ID pertama
	}

	fmt.Println("newID", newID)

	// Set EventID baru ke struct event
	event.EventID = newID

	// Eksekusi query untuk memasukkan data ke dalam tabel T_EventSekolah
	_, err = (*d.stmt)[insertEvent].ExecContext(ctx,
		event.EventID,
		event.EventHeader,
		event.EventStartDate,
		event.EventEndDate,
		event.EventDesc,
		event.EventImage,
	)

	if err != nil {
		result = "Gagal menyimpan data event"
		return result, errors.Wrap(err, "[DATA][InsertEvent]")
	}

	result = "Berhasil menyimpan data event"
	return result, nil
}

func (d Data) GetImageEvent(ctx context.Context, eventID string) ([]byte, error) {
	var poster []byte
	if err := (*d.stmt)[getImageEvent].QueryRowxContext(ctx, eventID).Scan(&poster); err != nil {
		return poster, errors.Wrap(err, "[DATA][GetImageEvent]")
	}

	return poster, nil
}

func generateImageURLFotoEvent(id string) string {
	var url = "http://localhost:8081"
	return fmt.Sprintf(url+"/ppdb/v1/data/getimageevent?eventID=%s", id)
}

func (d Data) GetEvent(ctx context.Context, searchInput string, offset, limit int) ([]ppdbEntity.TableEvent, error) {
	var (
		eventArray []ppdbEntity.TableEvent
		err        error
	)

	rows, err := (*d.stmt)[getEvent].QueryxContext(ctx, "%"+searchInput+"%", offset, limit)
	if err != nil {
		return eventArray, errors.Wrap(err, "[DATA] [GetEvent]")
	}
	defer rows.Close()

	// Ensure the directory for storing images exists
	imageDir := filepath.Join("public", "images")
	if err := EnsureDirectory(imageDir); err != nil {
		return nil, errors.Wrap(err, "[DATA] [GetEvent] - Failed to ensure directory")
	}

	for rows.Next() {
		var event ppdbEntity.TableEvent

		// Scan the event data from the database
		var eventStartDate sql.NullString
		var eventEndDate sql.NullString // Using sql.NullString to handle potential NULL values

		if err = rows.Scan(&event.EventID, &event.EventHeader, &eventStartDate, &eventEndDate, &event.EventDesc, &event.LinkEventImage); err != nil {
			return nil, errors.Wrap(err, "[DATA] [GetEvent] - Failed to scan row")
		}

		// Parse start date
		if eventStartDate.Valid {
			t, err := time.Parse("2006-01-02", eventStartDate.String) // Using the correct format
			if err != nil {
				return nil, errors.Wrap(err, "[DATA] [GetEvent] - Failed to parse eventStartDate")
			}
			event.EventStartDate = t // Set directly as time.Time
		} else {
			// If not valid, it remains the zero value of time.Time
			event.EventStartDate = time.Time{} // Optional: Set to a specific zero value if necessary
		}

		// Parse end date
		if eventEndDate.Valid {
			t, err := time.Parse("2006-01-02", eventEndDate.String) // Using the correct format
			if err != nil {
				return nil, errors.Wrap(err, "[DATA] [GetEvent] - Failed to parse eventEndDate")
			}
			event.EventEndDate = &t // Set as a pointer to time.Time
		} else {
			event.EventEndDate = nil // Set to nil if not valid
		}

		// Save image and generate URL if necessary
		filePath := filepath.Join(imageDir, event.EventID+".jpg")
		if err := saveImageToFile(event.EventImage, filePath); err != nil {
			return nil, errors.Wrap(err, "[DATA] [GetEvent] - Failed to save image")
		}
		// Generate URL for the event image (if applicable)
		event.LinkEventImage = generateImageURLFotoEvent(event.EventID)

		eventArray = append(eventArray, event)
	}

	if err = rows.Err(); err != nil {
		return nil, errors.Wrap(err, "[DATA] [GetEvent] - Row iteration error")
	}

	return eventArray, nil
}

func (d Data) GetEventPagination(ctx context.Context, searchInput string) (int, error) {
	var totalCount int

	// Query untuk mendapatkan total count tanpa LIMIT
	err := (*d.stmt)[getEventPagination].GetContext(ctx, &totalCount, "%"+searchInput+"%")
	if err != nil {
		return 0, errors.Wrap(err, "[DATA] [GetEventPagination] Error executing count query")
	}

	return totalCount, nil
}

func (d Data) GetEventDetail(ctx context.Context, eventID string) (ppdbEntity.TableEvent, error) {
	var (
		event ppdbEntity.TableEvent
		err   error
	)

	// Execute the query to get event details by eventID
	row := (*d.stmt)[getEventDetail].QueryRowxContext(ctx, eventID)

	// Scan event data from the row
	var eventStartDate sql.NullString
	var eventEndDate sql.NullString // Using sql.NullString to handle potential NULL values

	if err = row.Scan(&event.EventID, &event.EventHeader, &eventStartDate, &eventEndDate, &event.EventDesc, &event.LinkEventImage); err != nil {
		if err == sql.ErrNoRows {
			return event, errors.Wrap(err, "[DATA] [GetEventDetail] - Event not found")
		}
		return event, errors.Wrap(err, "[DATA] [GetEventDetail] - Failed to scan row")
	}

	// Parse start date
	if eventStartDate.Valid {
		t, err := time.Parse("2006-01-02", eventStartDate.String) // Using the correct format
		if err != nil {
			return event, errors.Wrap(err, "[DATA] [GetEventDetail] - Failed to parse eventStartDate")
		}
		event.EventStartDate = t // Set directly as time.Time
	} else {
		event.EventStartDate = time.Time{} // Optional: Set to a specific zero value if necessary
	}

	// Parse end date
	if eventEndDate.Valid {
		t, err := time.Parse("2006-01-02", eventEndDate.String) // Using the correct format
		if err != nil {
			return event, errors.Wrap(err, "[DATA] [GetEventDetail] - Failed to parse eventEndDate")
		}
		event.EventEndDate = &t // Set as a pointer to time.Time
	} else {
		event.EventEndDate = nil // Set to nil if not valid
	}

	// Generate URL for the event image (if applicable)
	event.LinkEventImage = generateImageURLFotoEvent(event.EventID)

	return event, nil
}

func (d Data) GetEventUtama(ctx context.Context) ([]ppdbEntity.TableEvent, error) {
	var (
		eventArray []ppdbEntity.TableEvent
		err        error
	)

	// Execute the query to get all events
	rows, err := (*d.stmt)[getEventUtama].QueryxContext(ctx)
	if err != nil {
		return eventArray, errors.Wrap(err, "[DATA] [GetEventUtama] - Failed to execute query")
	}
	defer rows.Close()

	for rows.Next() {
		var event ppdbEntity.TableEvent

		// Scan event data from the row
		var eventStartDate sql.NullString
		var eventEndDate sql.NullString // Using sql.NullString to handle potential NULL values

		if err = rows.Scan(&event.EventID, &event.EventHeader, &eventStartDate, &eventEndDate, &event.EventDesc, &event.LinkEventImage); err != nil {
			return nil, errors.Wrap(err, "[DATA] [GetEventUtama] - Failed to scan row")
		}

		// Parse start date
		if eventStartDate.Valid {
			t, err := time.Parse("2006-01-02", eventStartDate.String) // Using the correct format
			if err != nil {
				return nil, errors.Wrap(err, "[DATA] [GetEventUtama] - Failed to parse eventStartDate")
			}
			event.EventStartDate = t // Set directly as time.Time
		}

		// Parse end date
		if eventEndDate.Valid {
			t, err := time.Parse("2006-01-02", eventEndDate.String) // Using the correct format
			if err != nil {
				return nil, errors.Wrap(err, "[DATA] [GetEventUtama] - Failed to parse eventEndDate")
			}
			event.EventEndDate = &t // Set as a pointer to time.Time
		}

		// Generate URL for the event image (if applicable)
		event.LinkEventImage = generateImageURLFotoEvent(event.EventID)

		// Append the event to the array
		eventArray = append(eventArray, event)
	}

	if err = rows.Err(); err != nil {
		return nil, errors.Wrap(err, "[DATA] [GetEventUtama] - Row iteration error")
	}

	return eventArray, nil
}

func (d Data) UpdateEvent(ctx context.Context, event ppdbEntity.TableEvent, eventID string) (string, error) {
	var (
		result string
		err    error
	)

	_, err = (*d.stmt)[updateEvent].ExecContext(ctx,
		event.EventHeader,
		event.EventStartDate,
		event.EventEndDate,
		event.EventDesc,
		event.EventImage,
		eventID)

	if err != nil {
		result = "Gagal"
		return result, errors.Wrap(err, "[DATA][UpdateEvent]")
	}

	result = "Berhasil"
	return result, err
}

// Hapus Data Event
func (d Data) DeleteEvent(ctx context.Context, eventID string) (string, error) {
	var (
		err    error
		result string
	)

	_, err = (*d.stmt)[deleteEvent].ExecContext(ctx, eventID)

	if err != nil {
		result = "Gagal"
		return result, errors.Wrap(err, "[DATA][DeleteEvent]")
	}

	result = "Berhasil"
	return result, nil
}

// Peserta Didik
func (d Data) GetLastPesertaDidikId(ctx context.Context) (string, error) {
	var (
		err    error
		lastID string
		newID  string
	)

	err = (*d.stmt)[getLastPesertaDidikId].QueryRowxContext(ctx).Scan(&lastID)
	if err != nil && err != sql.ErrNoRows {
		return newID, errors.Wrap(err, "[DATA][GetLastPesertaDidikId]")
	}

	if lastID != "" {
		num, _ := strconv.Atoi(lastID[1:])
		newID = fmt.Sprintf("D%04d", num+1)
	} else {
		newID = "D0001"
	}

	return newID, nil
}

func (d Data) GetLastPembayaranFormulirId(ctx context.Context) (string, error) {
	var (
		err    error
		lastID string
		newID  string
	)

	err = (*d.stmt)[getLastPembayaranFormulirId].QueryRowxContext(ctx).Scan(&lastID)
	if err != nil && err != sql.ErrNoRows {
		return newID, errors.Wrap(err, "[DATA][GetLastPembayaranFormulirId]")
	}

	if lastID != "" {
		num, _ := strconv.Atoi(lastID[1:])
		newID = fmt.Sprintf("P%04d", num+1)
	} else {
		newID = "P0001"
	}

	return newID, nil
}

func (d Data) GetLastFormulirId(ctx context.Context) (string, error) {
	var (
		err    error
		lastID string
		newID  string
	)

	err = (*d.stmt)[getLastFormulirId].QueryRowxContext(ctx).Scan(&lastID)
	if err != nil && err != sql.ErrNoRows {
		return newID, errors.Wrap(err, "[DATA][GetLastFormulirId]")
	}

	if lastID != "" {
		num, _ := strconv.Atoi(lastID[1:])
		newID = fmt.Sprintf("F%04d", num+1)
	} else {
		newID = "F0001"
	}

	return newID, nil
}

func (d Data) GetLastKontakPesertaId(ctx context.Context) (string, error) {
	var (
		err    error
		lastID string
		newID  string
	)

	err = (*d.stmt)[getLastKontakPesertaId].QueryRowxContext(ctx).Scan(&lastID)
	if err != nil && err != sql.ErrNoRows {
		return newID, errors.Wrap(err, "[DATA][GetLastKontakPesertaId]")
	}

	if lastID != "" {
		num, _ := strconv.Atoi(lastID[1:])
		newID = fmt.Sprintf("K%04d", num+1)
	} else {
		newID = "K0001"
	}

	return newID, nil
}

func (d Data) GetLastOrtuId(ctx context.Context) (string, error) {
	var (
		err    error
		lastID string
		newID  string
	)

	err = (*d.stmt)[getLastOrtuId].QueryRowxContext(ctx).Scan(&lastID)
	if err != nil && err != sql.ErrNoRows {
		return newID, errors.Wrap(err, "[DATA][GetLastOrtuId]")
	}

	if lastID != "" {
		num, _ := strconv.Atoi(lastID[1:])
		newID = fmt.Sprintf("O%04d", num+1)
	} else {
		newID = "O0001"
	}

	return newID, nil
}

func (d Data) GetLastBerkasId(ctx context.Context) (string, error) {
	var (
		err    error
		lastID string
		newID  string
	)

	err = (*d.stmt)[getLastBerkasId].QueryRowxContext(ctx).Scan(&lastID)
	if err != nil && err != sql.ErrNoRows {
		return newID, errors.Wrap(err, "[DATA][GetLastBerkasId]")
	}

	if lastID != "" {
		num, _ := strconv.Atoi(lastID[1:])
		newID = fmt.Sprintf("B%04d", num+1)
	} else {
		newID = "B0001"
	}

	return newID, nil
}

func (d Data) GetLastJadwalTestId(ctx context.Context) (string, error) {
	var (
		err    error
		lastID string
		newID  string
	)

	err = (*d.stmt)[getLastJadwalTestId].QueryRowxContext(ctx).Scan(&lastID)
	if err != nil && err != sql.ErrNoRows {
		return newID, errors.Wrap(err, "[DATA][GetLastJadwalTestId]")
	}

	if lastID != "" {
		num, _ := strconv.Atoi(lastID[1:])
		newID = fmt.Sprintf("J%04d", num+1)
	} else {
		newID = "J0001"
	}

	return newID, nil
}

func (d Data) InsertPesertaDidik(ctx context.Context, pesertadidik ppdbEntity.TablePesertaDidik) (string, error) {
	var (
		err    error
		result string
	)

	// Parse the date string
	dateString := "0001-01-02"
	layout := "2006-01-02"
	defaultTime, err := time.Parse(layout, dateString)
	if err != nil {
		result = "Gagal parse"
		return result, errors.Wrap(err, "[DATA][GetLastPesertaDidikId]")
	}

	// Peserta Didik
	pesertadidik.PesertaID, err = d.GetLastPesertaDidikId(ctx)
	if err != nil && err != sql.ErrNoRows {
		result = "Gagal mengambil ID terakhir"
		return result, errors.Wrap(err, "[DATA][GetLastPesertaDidikId]")
	}

	// Pembayaran Formulir
	pembayaranformulir := ppdbEntity.TablePembayaranFormulir{}

	pembayaranformulir.PembayaranID, err = d.GetLastPembayaranFormulirId(ctx)
	if err != nil && err != sql.ErrNoRows {
		result = "Gagal mengambil ID terakhir"
		return result, errors.Wrap(err, "[DATA][GetLastPesertaDidikId]")
	}
	pembayaranformulir.PesertaID = pesertadidik.PesertaID
	pembayaranformulir.StatusID = "S0001"
	pembayaranformulir.HargaFormulir = 150000
	pembayaranformulir.TglPembayaran = defaultTime

	// Formulir
	formulir := ppdbEntity.TableFormulir{}

	formulir.FormulirID, err = d.GetLastFormulirId(ctx)
	if err != nil && err != sql.ErrNoRows {
		result = "Gagal mengambil ID terakhir"
		return result, errors.Wrap(err, "[DATA][GetLastFormulirId]")
	}
	formulir.PesertaID = pesertadidik.PesertaID
	formulir.PembayaranID = pembayaranformulir.PembayaranID
	formulir.StatusID = "S0001"
	formulir.TglLahir = defaultTime
	formulir.TglSubmit = defaultTime

	// Kontak Peserta
	kontakpeserta := ppdbEntity.TableKontakPeserta{}

	kontakpeserta.KontakID, err = d.GetLastKontakPesertaId(ctx)
	if err != nil && err != sql.ErrNoRows {
		result = "Gagal mengambil ID terakhir"
		return result, errors.Wrap(err, "[DATA][GetLastKontakPesertaId]")
	}
	kontakpeserta.FormulirID = formulir.FormulirID

	// Ortu
	ortu := ppdbEntity.TableOrtu{}

	ortu.OrtuID, err = d.GetLastOrtuId(ctx)
	if err != nil && err != sql.ErrNoRows {
		result = "Gagal mengambil ID terakhir"
		return result, errors.Wrap(err, "[DATA][GetLastOrtuId]")
	}
	ortu.FormulirID = formulir.FormulirID

	// Berkas
	berkas := ppdbEntity.TableBerkas{}

	berkas.BerkasID, err = d.GetLastBerkasId(ctx)
	if err != nil && err != sql.ErrNoRows {
		result = "Gagal mengambil ID terakhir"
		return result, errors.Wrap(err, "[DATA][GetLastBerkasId]")
	}
	berkas.PesertaID = pesertadidik.PesertaID
	berkas.StatusID = "S0001"
	berkas.TanggalUpload = defaultTime

	// JadwalTest
	jadwaltest := ppdbEntity.TableJadwalTest{}

	jadwaltest.TestID, err = d.GetLastJadwalTestId(ctx)
	if err != nil && err != sql.ErrNoRows {
		result = "Gagal mengambil ID terakhir"
		return result, errors.Wrap(err, "[DATA][GetLastJadwalTestId]")
	}
	jadwaltest.PesertaID = pesertadidik.PesertaID
	jadwaltest.StatusID = "S0001"
	jadwaltest.TglTest = defaultTime
	jadwaltest.WaktuTest = defaultTime

	// Insert Peserta Didik
	_, err = (*d.stmt)[insertPesertaDidik].ExecContext(ctx,
		pesertadidik.PesertaID,
		pesertadidik.PesertaName,
		pesertadidik.Password,
		pesertadidik.EmailPeserta,
		pesertadidik.NoTelpHpPeserta,
		pesertadidik.SekolahAsalYN,
		pesertadidik.SekolahAsal,
		pesertadidik.AlamatSekolahAsal)
	if err != nil {
		return "Gagal menyimpan data peserta didik", errors.Wrap(err, "[DATA][InsertPesertaDidik]")
	}

	// Insert Pembayaran Formulir
	_, err = (*d.stmt)[insertPembayaranFormulir].ExecContext(ctx,
		pembayaranformulir.PembayaranID,
		pembayaranformulir.PesertaID,
		pembayaranformulir.StatusID,
		pembayaranformulir.TglPembayaran,
		pembayaranformulir.HargaFormulir)
	if err != nil {
		return "Gagal menyimpan data pembayaran formulir", errors.Wrap(err, "[DATA][InsertPembayaranFormulir]")
	}

	// Insert Formulir
	_, err = (*d.stmt)[insertFormulir].ExecContext(ctx,
		formulir.FormulirID,
		formulir.PesertaID,
		formulir.PembayaranID,
		formulir.TglLahir,
		formulir.TglSubmit,
		formulir.StatusID)
	if err != nil {
		return "Gagal menyimpan data formulir", errors.Wrap(err, "[DATA][InsertFormulir]")
	}

	// Insert Kontak Peserta
	_, err = (*d.stmt)[insertKontakPeserta].ExecContext(ctx,
		kontakpeserta.KontakID,
		kontakpeserta.FormulirID)
	if err != nil {
		return "Gagal menyimpan data kontak peserta", errors.Wrap(err, "[DATA][InsertKontakPeserta]")
	}

	// Insert Ortu
	_, err = (*d.stmt)[insertOrtu].ExecContext(ctx,
		ortu.OrtuID,
		ortu.FormulirID)
	if err != nil {
		return "Gagal menyimpan data ortu", errors.Wrap(err, "[DATA][InsertOrtu]")
	}

	// Insert Berkas
	_, err = (*d.stmt)[insertBerkas].ExecContext(ctx,
		berkas.BerkasID,
		berkas.PesertaID,
		berkas.StatusID,
		berkas.TanggalUpload,
	)
	if err != nil {
		return "Gagal menyimpan data berkas", errors.Wrap(err, "[DATA][InsertBerkas]")
	}

	// Insert Jadwal Test
	_, err = (*d.stmt)[insertJadwalTest].ExecContext(ctx,
		jadwaltest.TestID,
		jadwaltest.PesertaID,
		jadwaltest.StatusID,
		jadwaltest.TglTest,
		jadwaltest.WaktuTest,
	)
	if err != nil {
		return "Gagal menyimpan data jadwal test", errors.Wrap(err, "[DATA][InsertJadwalTest]")
	}

	result = "Berhasil menyimpan data peserta didik"
	return result, nil
}

func (d Data) GetLoginCheck(ctx context.Context, login ppdbEntity.TablePesertaDidik) (ppdbEntity.TablePesertaDidik, error) {
	var (
		err    error
		result ppdbEntity.TablePesertaDidik
	)

	err = (*d.stmt)[getLoginCheck].QueryRowxContext(ctx, login.EmailPeserta).StructScan(&result)
	if err != nil {
		return result, errors.Wrap(err, "[DATA][GetLoginCheck]")
	}

	return result, nil
}

func (d Data) GetPesertaDidikDetail(ctx context.Context, idpesertadidik string) (ppdbEntity.TablePesertaDidik, error) {
	var (
		err    error
		result ppdbEntity.TablePesertaDidik
	)

	err = (*d.stmt)[getPesertaDidikDetail].QueryRowxContext(ctx, idpesertadidik).StructScan(&result)
	if err != nil {
		return result, errors.Wrap(err, "[DATA][GetLoginCheck]")
	}

	return result, nil
}

func (d Data) GetPembayaranFormulirDetail(ctx context.Context, idpesertadidik string) (ppdbEntity.TablePembayaranFormulir, error) {

	var (
		tglPembayaran      sql.NullString
		pembayaranformulir ppdbEntity.TablePembayaranFormulir
	)

	err := (*d.stmt)[getPembayaranFormulirDetail].QueryRowxContext(ctx, idpesertadidik).Scan(
		&pembayaranformulir.PembayaranID,
		&pembayaranformulir.PesertaID,
		&pembayaranformulir.PesertaName,
		&pembayaranformulir.StatusID,
		&pembayaranformulir.StatusName,
		&tglPembayaran,
		&pembayaranformulir.HargaFormulir,
		&pembayaranformulir.BuktiPembayaran)
	if err != nil {
		return pembayaranformulir, errors.Wrap(err, "[DATA][GetPembayaranFormulirDetail]")
	}

	t, err := time.Parse("2006-01-02", tglPembayaran.String)
	if err != nil {
		return pembayaranformulir, errors.Wrap(err, "[DATA] [GetPembayaranFormulirDetail] - Failed to parse date")
	}
	pembayaranformulir.TglPembayaran = t

	return pembayaranformulir, nil
}

func (d Data) GetFormulirDetail(ctx context.Context, idpesertadidik string) (ppdbEntity.TableDataFormulir, error) {

	var (
		tglLahir  sql.NullString
		tglSubmit sql.NullString
		formulir  ppdbEntity.TableDataFormulir
	)

	err := (*d.stmt)[getFormulirDetail].QueryRowxContext(ctx, idpesertadidik).Scan(
		&formulir.FormulirID,
		&formulir.PesertaID,
		&formulir.PembayaranID,
		&formulir.JurusanID,
		&formulir.AgamaID,
		&formulir.GenderPeserta,
		&formulir.NoAktaLahir,
		&formulir.TempatLahir,
		&tglLahir,
		&formulir.NISN,
		&formulir.Kelas,
		&tglSubmit,
		&formulir.StatusID,
		&formulir.StatusName,
		&formulir.KontakID,
		&formulir.AlamatTerakhir,
		&formulir.KodePos,
		&formulir.NoTelpRumah,
		&formulir.OrtuID,
		&formulir.NamaAyah,
		&formulir.PekerjaanAyah,
		&formulir.NoTelpHpAyah,
		&formulir.NamaIbu,
		&formulir.PekerjaanIbu,
		&formulir.NoTelpHpIbu,
		&formulir.NamaWali,
		&formulir.PekerjaanWali,
		&formulir.NoTelpHpWali,
		&formulir.PesertaName,
		&formulir.NoTelpHpPeserta,
		&formulir.SekolahAsal,
		&formulir.AlamatSekolahAsal,
		&formulir.JurusanName,
		&formulir.AgamaName)
	if err != nil {
		return formulir, errors.Wrap(err, "[DATA][GetFormulirDetail]")
	}

	tLahir, err := time.Parse("2006-01-02", tglLahir.String)
	if err != nil {
		return formulir, errors.Wrap(err, "[DATA] [GetFormulirDetail] - Failed to parse tgl lahir")
	}
	formulir.TglLahir = tLahir

	tSubmit, err := time.Parse("2006-01-02 15:04:05", tglSubmit.String)
	if err != nil {
		return formulir, errors.Wrap(err, "[DATA] [GetFormulirDetail] - Failed to parse tgl submit")
	}
	formulir.TglSubmit = tSubmit

	return formulir, nil
}

func (d Data) GetBerkasDetail(ctx context.Context, idpesertadidik string) (ppdbEntity.TableBerkas, error) {

	var (
		tglUpload sql.NullString
		berkas    ppdbEntity.TableBerkas
	)

	err := (*d.stmt)[getBerkasDetail].QueryRowxContext(ctx, idpesertadidik).Scan(
		&berkas.BerkasID,
		&berkas.PesertaID,
		&berkas.PesertaName,
		&berkas.StatusID,
		&berkas.StatusName,
		&berkas.AktalLahir,
		&berkas.PasPhoto,
		&berkas.Rapor,
		&tglUpload)
	if err != nil {
		return berkas, errors.Wrap(err, "[DATA][GetBerkasDetail]")
	}

	t, err := time.Parse("2006-01-02", tglUpload.String)
	if err != nil {
		return berkas, errors.Wrap(err, "[DATA] [GetBerkasDetail] - Failed to parse date")
	}
	berkas.TanggalUpload = t

	return berkas, nil
}

func (d Data) GetJadwalTestDetail(ctx context.Context, idpesertadidik string) (ppdbEntity.TableJadwalTest, error) {

	var (
		tglTest    sql.NullString
		waktuTest  sql.NullString
		jadwaltest ppdbEntity.TableJadwalTest
	)

	err := (*d.stmt)[getJadwalTestDetail].QueryRowxContext(ctx, idpesertadidik).Scan(
		&jadwaltest.TestID,
		&jadwaltest.PesertaID,
		&jadwaltest.PesertaName,
		&jadwaltest.StatusID,
		&jadwaltest.StatusName,
		&tglTest,
		&waktuTest)
	if err != nil {
		return jadwaltest, errors.Wrap(err, "[DATA][GetJadwalTestDetail]")
	}

	tTest, err := time.Parse("2006-01-02", tglTest.String)
	if err != nil {
		return jadwaltest, errors.Wrap(err, "[DATA] [GetJadwalTestDetail] - Failed to parse date")
	}
	jadwaltest.TglTest = tTest

	tWaktu, err := time.Parse("15:04:05", waktuTest.String)
	if err != nil {
		return jadwaltest, errors.Wrap(err, "[DATA] [GetJadwalTestDetail] - Failed to parse date")
	}
	jadwaltest.WaktuTest = tWaktu

	return jadwaltest, nil
}

func (d Data) UpdatePembayaranFormulir(ctx context.Context, pembayaranformulir ppdbEntity.TablePembayaranFormulir) (string, error) {

	_, err := (*d.stmt)[updatePembayaranFormulir].ExecContext(ctx,
		pembayaranformulir.StatusID,
		pembayaranformulir.BuktiPembayaran,
		pembayaranformulir.PembayaranID,
	)
	if err != nil {
		return "Gagal update data pembayaran formulir", errors.Wrap(err, "[DATA][UpdatePembayaranFormulir]")
	}

	result := "Berhasil update data pembayaran formulir"

	return result, nil
}

func (d Data) UpdateFormulir(ctx context.Context, formulir ppdbEntity.TableDataFormulir) (string, error) {

	_, err := (*d.stmt)[updateFormulir].ExecContext(ctx,
		formulir.JurusanID,
		formulir.AgamaID,
		formulir.GenderPeserta,
		formulir.NoAktaLahir,
		formulir.TempatLahir,
		formulir.TglLahir,
		formulir.NISN,
		formulir.Kelas,
		formulir.StatusID,
		formulir.FormulirID,
	)
	if err != nil {
		return "Gagal update data formulir", errors.Wrap(err, "[DATA][UpdateFormulir]")
	}

	result := "Berhasil update data formulir"

	return result, nil
}

func (d Data) UpdateKontakPeserta(ctx context.Context, kontakpeserta ppdbEntity.TableDataFormulir) (string, error) {

	_, err := (*d.stmt)[updateKontakPeserta].ExecContext(ctx,
		kontakpeserta.AlamatTerakhir,
		kontakpeserta.KodePos,
		kontakpeserta.NoTelpRumah,
		kontakpeserta.KontakID,
	)
	if err != nil {
		return "Gagal update data kontak peserta", errors.Wrap(err, "[DATA][UpdateKontakPeserta]")
	}

	result := "Berhasil update data kontak peserta"

	return result, nil
}

func (d Data) UpdateOrtu(ctx context.Context, ortu ppdbEntity.TableDataFormulir) (string, error) {

	_, err := (*d.stmt)[updateOrtu].ExecContext(ctx,
		ortu.NamaAyah,
		ortu.PekerjaanAyah,
		ortu.NoTelpHpAyah,
		ortu.NamaIbu,
		ortu.PekerjaanIbu,
		ortu.NoTelpHpIbu,
		ortu.NamaWali,
		ortu.PekerjaanWali,
		ortu.NoTelpHpWali,
		ortu.OrtuID,
	)
	if err != nil {
		return "Gagal update data ortu", errors.Wrap(err, "[DATA][UpdateOrtu]")
	}

	result := "Berhasil update data ortu"

	return result, nil
}

func (d Data) UpdateBerkas(ctx context.Context, berkas ppdbEntity.TableBerkas) (string, error) {

	_, err := (*d.stmt)[updateBerkas].ExecContext(ctx,
		berkas.StatusID,
		berkas.AktalLahir,
		berkas.PasPhoto,
		berkas.Rapor,
		berkas.BerkasID,
	)
	if err != nil {
		return "Gagal update data berkas", errors.Wrap(err, "[DATA][UpdateBerkas]")
	}

	result := "Berhasil update data berkas"

	return result, nil
}

func (d Data) UpdateJadwalTest(ctx context.Context, jadwalTest ppdbEntity.TableJadwalTest) (string, error) {

	_, err := (*d.stmt)[updateJadwalTest].ExecContext(ctx,
		jadwalTest.StatusID,
		jadwalTest.TglTest,
		jadwalTest.WaktuTest,
		jadwalTest.TestID,
	)
	if err != nil {
		return "Gagal update data jadwal test", errors.Wrap(err, "[DATA][UpdateJadwalTest]")
	}

	result := "Berhasil update data jadwal test"

	return result, nil
}

func (d Data) GetJadwalTestAll(ctx context.Context, searchInput string, offset, limit int) ([]ppdbEntity.TableJadwalTest, error) {

	var (
		tglTest         sql.NullString
		waktuTest       sql.NullString
		jadwaltest      ppdbEntity.TableJadwalTest
		jadwaltestArray []ppdbEntity.TableJadwalTest
	)

	// Menjalankan query dengan parameter searchInput, offset, dan limit
	rows, err := (*d.stmt)[getJadwalTestAll].QueryxContext(ctx, "%"+searchInput+"%", offset, limit)

	if err != nil {
		return jadwaltestArray, errors.Wrap(err, "[DATA][GetJadwalTestAll] - Query failed")
	}
	defer rows.Close()

	// Iterasi melalui hasil query dan mengisi slice jadwaltests
	for rows.Next() {

		// Scan data dari hasil query ke dalam objek jadwaltest
		err := rows.Scan(
			&jadwaltest.TestID,
			&jadwaltest.PesertaID,
			&jadwaltest.PesertaName,
			&jadwaltest.StatusID,
			&jadwaltest.StatusName,
			&tglTest,
			&waktuTest,
		)
		if err != nil {
			return jadwaltestArray, errors.Wrap(err, "[DATA][GetJadwalTestAll] - Scan failed")
		}

		// Parse tglTest dan waktuTest jika valid
		if tglTest.Valid {
			tTest, err := time.Parse("2006-01-02", tglTest.String)
			if err != nil {
				return jadwaltestArray, errors.Wrap(err, "[DATA][GetJadwalTestAll] - Failed to parse tglTest")
			}
			jadwaltest.TglTest = tTest
		}

		if waktuTest.Valid {
			tWaktu, err := time.Parse("15:04:05", waktuTest.String)
			if err != nil {
				return jadwaltestArray, errors.Wrap(err, "[DATA][GetJadwalTestAll] - Failed to parse waktuTest")
			}
			jadwaltest.WaktuTest = tWaktu
		}

		// Tambahkan jadwaltest ke dalam slice jadwaltests
		jadwaltestArray = append(jadwaltestArray, jadwaltest)
	}

	// Cek apakah ada error setelah iterasi
	if err := rows.Err(); err != nil {
		return jadwaltestArray, errors.Wrap(err, "[DATA][GetJadwalTestAll] - Rows iteration failed")
	}

	// Mengembalikan slice jadwaltests
	return jadwaltestArray, nil
}

func (d Data) GetJadwalTestPagination(ctx context.Context, searchInput string) (int, error) {
	var totalCount int

	// Query untuk mendapatkan total count jadwal ujian tanpa LIMIT
	err := (*d.stmt)[getJadwalTestPagination].GetContext(ctx, &totalCount, "%"+searchInput+"%")
	if err != nil {
		return 0, errors.Wrap(err, "[DATA] [GetJadwalTestPagination] Error executing count query")
	}

	return totalCount, nil
}

func (d Data) GetPembayaranFormulirAll(ctx context.Context, searchInput string, offset, limit int) ([]ppdbEntity.TablePembayaranFormulir, error) {

	var (
		tglPembayaran       sql.NullString
		pembayaranFormulir  ppdbEntity.TablePembayaranFormulir
		pembayaranFormulirArray []ppdbEntity.TablePembayaranFormulir
	)

	// Menjalankan query dengan parameter searchInput, offset, dan limit
	rows, err := (*d.stmt)[getPembayaranFormulirAll].QueryxContext(ctx, "%"+searchInput+"%", offset, limit)

	if err != nil {
		return pembayaranFormulirArray, errors.Wrap(err, "[DATA][GetPembayaranFormulirAll] - Query failed")
	}
	defer rows.Close()

	// Iterasi melalui hasil query dan mengisi slice pembayaranFormulirArray
	for rows.Next() {

		// Scan data dari hasil query ke dalam objek pembayaranFormulir
		err := rows.Scan(
			&pembayaranFormulir.PembayaranID,
			&pembayaranFormulir.PesertaID,
			&pembayaranFormulir.PesertaName,
			&pembayaranFormulir.StatusID,
			&pembayaranFormulir.StatusName,
			&tglPembayaran,
			&pembayaranFormulir.HargaFormulir,
			&pembayaranFormulir.BuktiPembayaran,
		)
		if err != nil {
			return pembayaranFormulirArray, errors.Wrap(err, "[DATA][GetPembayaranFormulirAll] - Scan failed")
		}

		// Parse tglPembayaran jika valid
		if tglPembayaran.Valid {
			tPembayaran, err := time.Parse("2006-01-02", tglPembayaran.String)
			if err != nil {
				return pembayaranFormulirArray, errors.Wrap(err, "[DATA][GetPembayaranFormulirAll] - Failed to parse tglPembayaran")
			}
			pembayaranFormulir.TglPembayaran = tPembayaran
		}

		// Tambahkan pembayaranFormulir ke dalam slice pembayaranFormulirArray
		pembayaranFormulirArray = append(pembayaranFormulirArray, pembayaranFormulir)
	}

	// Cek apakah ada error setelah iterasi
	if err := rows.Err(); err != nil {
		return pembayaranFormulirArray, errors.Wrap(err, "[DATA][GetPembayaranFormulirAll] - Rows iteration failed")
	}

	// Mengembalikan slice pembayaranFormulirArray
	return pembayaranFormulirArray, nil
}

func (d Data) GetPembayaranFormulirPagination(ctx context.Context, searchInput string) (int, error) {
	var totalCount int

	// Query untuk mendapatkan total count pembayaran formulir tanpa LIMIT
	err := (*d.stmt)[getPembayaranFormulirPagination].GetContext(ctx, &totalCount, "%"+searchInput+"%")
	if err != nil {
		return 0, errors.Wrap(err, "[DATA] [GetPembayaranFormulirPagination] Error executing count query")
	}

	return totalCount, nil
}
