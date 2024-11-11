package ppdb

import (
	"bytes"
	"context"
	"fmt"
	"image"
	"io/ioutil"
	"log"
	"math"
	"os"
	"strconv"
	"time"

	// ppdbEntity "ppdb-be/internal/entity/ppdb"
	"ppdb-be/internal/entity/ppdb"
	ppdbEntity "ppdb-be/internal/entity/ppdb"
	"ppdb-be/pkg/errors"

	"github.com/jung-kurt/gofpdf"
	"golang.org/x/crypto/bcrypt"
	// "encoding/json"
	// "fmt"
	// "log"
	// "strconv"
	// "time"
)

// func (s Service) GetKaryawan(ctx context.Context) ([]JoEntity.GetKaryawan, interface{}, error) {
// 	var (
// 		total int
// 	)

// 	metadata := make(map[string]interface{})
// 	karyawanArray, err := s.ppdb.GetKaryawan(ctx)

// 	if err != nil {
// 		return karyawanArray, metadata, errors.Wrap(err, "[Service][GetKaryawan]")
// 	}

// 	total, err = s.ppdb.GetCountKaryawan(ctx)

// 	if err != nil {
// 		return karyawanArray, metadata, errors.Wrap(err, "[Service][GetCountKaryawan]")
// 	}
// 	metadata["total_data"] = total

// 	return karyawanArray, metadata, nil

// }

// func (s Service) InsertKaryawan(ctx context.Context, karyawan JoEntity.InsertKaryawan) (string, error) {

// 	var (
// 		result string
// 	)
// 	result, err := s.ppdb.InsertKaryawan(ctx, karyawan.Insertkaryawan)

// 	if err != nil {
// 		result = "Gagal"
// 		return result, errors.Wrap(err, "[Service][InsertKaryawan]")
// 	}

// 	result = "Berhasil"

// 	return result, err
// }

func (s Service) LoginAdmin(ctx context.Context, emailAdmin string, password string) (string, error) {
	var (
		result string
	)

	result, err := s.ppdb.LoginAdmin(ctx, emailAdmin, password)

	if err != nil {
		result = "Gagal Login"
		return result, errors.Wrap(err, "[Service][LoginAdmin]")
	}
	result = "Berhasil Login"

	return result, err

}

func (s Service) GetKontakSekolah(ctx context.Context) ([]ppdbEntity.TableKontakSekolah, error) {

	kontakArray, err := s.ppdb.GetKontakSekolah(ctx)

	if err != nil {
		return kontakArray, errors.Wrap(err, "[Service][GetKontakSekolah]")
	}

	return kontakArray, nil

}

// kelola admin
func (s Service) GetDataAdminSlim(ctx context.Context, searchInput string, page, length int) ([]ppdbEntity.TableKelolaDataAdmin, interface{}, error) {
	limit := length
	offset := (page - 1) * length
	var lastPage int
	var metadata = make(map[string]int)
	admins := []ppdbEntity.TableKelolaDataAdmin{}

	// Pagination
	if page > 0 && length > 0 {
		// Get total count of admins for pagination
		count, err := s.ppdb.GetDataAdminPagination(ctx, searchInput)
		if err != nil {
			return admins, metadata, errors.Wrap(err, "[SERVICE][GetDataAdminSlim] Error getting pagination count")
		}

		// Calculate last page based on count and length
		lastPage = int(math.Ceil(float64(count) / float64(length)))

		// Prepare metadata
		metadata["total_data"] = count
		metadata["total_page"] = lastPage

		// Get paginated admin data
		admins, err = s.ppdb.GetDataAdmin(ctx, searchInput, offset, limit)
		if err != nil {
			return admins, metadata, errors.Wrap(err, "[SERVICE][GetDataAdminSlim] Error getting paginated admin data")
		}

		return admins, metadata, nil
	}

	// If pagination parameters are not provided, return all data
	admins, err := s.ppdb.GetDataAdmin(ctx, searchInput, 0, 0) // Adjust method for retrieving all records
	if err != nil {
		return admins, metadata, errors.Wrap(err, "[SERVICE][GetDataAdminSlim] Error getting all admin data")
	}

	return admins, metadata, nil
}

func (s Service) InsertDataAdmin(ctx context.Context, admin ppdbEntity.TableAdmin) (string, error) {
	var (
		result string
	)

	// Panggil fungsi InsertDataAdmin dari data layer
	result, err := s.ppdb.InsertDataAdmin(ctx, admin)

	if err != nil {
		result = "Gagal"
		return result, errors.Wrap(err, "[Service][InsertDataAdmin]")
	}

	result = "Berhasil"

	return result, nil
}

func (s Service) DeleteAdmin(ctx context.Context, adminID string) (string, error) {
	result, err := s.ppdb.DeleteAdmin(ctx, adminID)

	if err != nil {
		return result, errors.Wrap(err, "[Service][DeleteAdmin]")
	}

	return result, nil
}

func (s Service) GetRole(ctx context.Context) ([]ppdbEntity.TableRole, error) {

	roleArray, err := s.ppdb.GetRole(ctx)

	if err != nil {
		return roleArray, errors.Wrap(err, "[Service][GetRole]")
	}

	return roleArray, nil

}

func (s Service) GetAgama(ctx context.Context) ([]ppdbEntity.TableAgama, error) {

	agamaArray, err := s.ppdb.GetAgama(ctx)

	if err != nil {
		return agamaArray, errors.Wrap(err, "[Service][GetAgama]")
	}

	return agamaArray, nil

}

func (s Service) GetJurusan(ctx context.Context) ([]ppdbEntity.TableJurusan, error) {

	jurusanArray, err := s.ppdb.GetJurusan(ctx)

	if err != nil {
		return jurusanArray, errors.Wrap(err, "[Service][GetJurusan]")
	}

	return jurusanArray, nil

}

func (s Service) GetStatus(ctx context.Context) ([]ppdbEntity.TableStatus, error) {

	statusArray, err := s.ppdb.GetStatus(ctx)

	if err != nil {
		return statusArray, errors.Wrap(err, "[Service][GetStatus]")
	}

	return statusArray, nil

}

func (s Service) InsertInfoDaftar(ctx context.Context, infoDaftar ppdbEntity.TableInfoDaftar) (string, error) {
	var (
		result string
	)

	// Panggil fungsi InsertInfoDaftar dari data layer
	result, err := s.ppdb.InsertInfoDaftar(ctx, infoDaftar)

	if err != nil {
		result = "Gagal"
		return result, errors.Wrap(err, "[Service][InsertInfoDaftar]")
	}

	result = "Berhasil"

	return result, nil
}

func (s Service) GetGambarInfoDaftar(ctx context.Context, infoID string) ([]byte, error) {
	poster, err := s.ppdb.GetGambarInfoDaftar(ctx, infoID)
	if err != nil {
		return poster, errors.Wrap(err, "[SERVICE][GetGambarInfoDaftar]")
	}

	return poster, err
}

func (s Service) GetInfoDaftar(ctx context.Context) ([]ppdbEntity.TableInfoDaftar, error) {
	infoDaftar, err := s.ppdb.GetInfoDaftar(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "[SERVICE] [GetInfoDaftar]")
	}

	return infoDaftar, nil
}

func (s Service) UpdateInfoDaftar(ctx context.Context, infoDaftar ppdbEntity.TableInfoDaftar, infoID string) (string, error) {
	var (
		result string
	)

	result, err := s.ppdb.UpdateInfoDaftar(ctx, infoDaftar, infoID)
	if err != nil {
		result = "Gagal"
		return result, errors.Wrap(err, "[Service][UpdateInfoDaftar]")
	}

	result = "Berhasil"
	return result, err
}

// Banner
func (s Service) InsertBanner(ctx context.Context, banner ppdbEntity.TableBanner) (string, error) {
	var (
		result string
	)

	result, err := s.ppdb.InsertBanner(ctx, banner)

	if err != nil {
		result = "Gagal"
		return result, errors.Wrap(err, "[Service][InsertBanner]")
	}

	result = "Berhasil"

	return result, nil
}

func (s Service) GetGambarBanner(ctx context.Context, bannerID string) ([]byte, error) {
	poster, err := s.ppdb.GetGambarBanner(ctx, bannerID)
	if err != nil {
		return poster, errors.Wrap(err, "[SERVICE][GetGambarBanner]")
	}

	return poster, err
}

func (s Service) GetBanner(ctx context.Context) ([]ppdbEntity.TableBanner, error) {
	banner, err := s.ppdb.GetBanner(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "[SERVICE] [GetBanner]")
	}

	return banner, nil
}

func (s Service) DeleteBanner(ctx context.Context, bannerID string) (string, error) {
	result, err := s.ppdb.DeleteBanner(ctx, bannerID)

	if err != nil {
		return result, errors.Wrap(err, "[Service][DeleteBanner]")
	}

	return result, nil
}

func (s Service) UpdateBanner(ctx context.Context, banner ppdbEntity.TableBanner, bannerID string) (string, error) {
	var (
		result string
	)

	result, err := s.ppdb.UpdateBanner(ctx, banner, bannerID)
	if err != nil {
		result = "Gagal"
		return result, errors.Wrap(err, "[Service][UpdateBanner]")
	}

	result = "Berhasil"
	return result, err
}

// Fasilitas
func (s Service) InsertFasilitas(ctx context.Context, fasilitas ppdbEntity.TableFasilitas) (string, error) {
	var (
		result string
	)

	result, err := s.ppdb.InsertFasilitas(ctx, fasilitas)

	if err != nil {
		result = "Gagal"
		return result, errors.Wrap(err, "[Service][InsertFasilitas]")
	}

	result = "Berhasil"

	return result, nil
}

func (s Service) GetGambarFasilitas(ctx context.Context, fasilitasID string) ([]byte, error) {
	poster, err := s.ppdb.GetGambarFasilitas(ctx, fasilitasID)
	if err != nil {
		return poster, errors.Wrap(err, "[SERVICE][GetGambarFasilitas]")
	}

	return poster, err
}

func (s Service) GetFasilitasSlim(ctx context.Context, searchInput string, page, length int) ([]ppdbEntity.TableFasilitas, interface{}, error) {
	limit := length
	offset := (page - 1) * length
	var lastPage int
	metadata := make(map[string]int)
	fasilitas := []ppdbEntity.TableFasilitas{}

	// Pagination
	if page > 0 && length > 0 {
		// Get total count of fasilitas for pagination
		count, err := s.ppdb.GetFasilitasPagination(ctx, searchInput)
		if err != nil {
			return fasilitas, metadata, errors.Wrap(err, "[SERVICE][GetFasilitas] Error getting pagination count")
		}

		// Calculate last page based on count and length
		lastPage = int(math.Ceil(float64(count) / float64(length)))

		// Prepare metadata
		metadata["total_data"] = count
		metadata["total_page"] = lastPage

		// Get paginated fasilitas data
		fasilitas, err = s.ppdb.GetFasilitas(ctx, searchInput, offset, limit)
		if err != nil {
			return fasilitas, metadata, errors.Wrap(err, "[SERVICE][GetFasilitas] Error getting paginated fasilitas data")
		}

		return fasilitas, metadata, nil
	}

	// If page or length is invalid, get all data without pagination
	fasilitas, err := s.ppdb.GetFasilitas(ctx, searchInput, offset, limit)
	if err != nil {
		return fasilitas, metadata, errors.Wrap(err, "[SERVICE][GetFasilitas] Error getting fasilitas data")
	}

	return fasilitas, metadata, nil
}

func (s Service) GetFasilitas(ctx context.Context) ([]ppdbEntity.TableFasilitas, error) {
	fasilitas, err := s.ppdb.GetFasilitasUtama(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "[SERVICE] [GetFasilitas]")
	}

	return fasilitas, nil
}

func (s Service) DeleteFasilitas(ctx context.Context, fasilitasID string) (string, error) {
	result, err := s.ppdb.DeleteFasilitas(ctx, fasilitasID)

	if err != nil {
		return result, errors.Wrap(err, "[Service][DeleteFasilitas]")
	}

	return result, nil
}

func (s Service) UpdateFasilitas(ctx context.Context, fasilitas ppdbEntity.TableFasilitas, fasilitasID string) (string, error) {
	var (
		result string
	)

	result, err := s.ppdb.UpdateFasilitas(ctx, fasilitas, fasilitasID)
	if err != nil {
		result = "Gagal"
		return result, errors.Wrap(err, "[Service][UpdateFasilitas]")
	}

	result = "Berhasil"
	return result, err
}

// Profile Staff
func (s Service) InsertProfileStaff(ctx context.Context, staff ppdbEntity.TableStaff) (string, error) {
	var (
		result string
	)

	// Panggil fungsi InsertProfileStaff dari data layer
	result, err := s.ppdb.InsertProfileStaff(ctx, staff)

	if err != nil {
		result = "Gagal menyimpan data staff"
		return result, errors.Wrap(err, "[Service][InsertProfileStaff]")
	}

	result = "Berhasil menyimpan data staff"
	return result, nil
}

func (s Service) GetPhotoStaff(ctx context.Context, staffID string) ([]byte, error) {
	poster, err := s.ppdb.GetPhotoStaff(ctx, staffID)
	if err != nil {
		return poster, errors.Wrap(err, "[SERVICE][GetPhotoStaff]")
	}

	return poster, err
}

func (s Service) GetProfileStaffSlim(ctx context.Context, searchInput string, page, length int) ([]ppdbEntity.TableStaff, interface{}, error) {
	limit := length
	offset := (page - 1) * length
	var lastPage int
	metadata := make(map[string]int)
	staff := []ppdbEntity.TableStaff{}

	// Pagination
	if page > 0 && length > 0 {
		// Get total count of staff for pagination
		count, err := s.ppdb.GetProfileStaffPagination(ctx, searchInput)
		if err != nil {
			return staff, metadata, errors.Wrap(err, "[SERVICE][GetProfileStaffSlim] Error getting pagination count")
		}

		// Calculate last page based on count and length
		lastPage = int(math.Ceil(float64(count) / float64(length)))

		// Prepare metadata
		metadata["total_data"] = count
		metadata["total_page"] = lastPage

		// Get paginated staff data
		staff, err = s.ppdb.GetProfileStaff(ctx, searchInput, offset, limit)
		if err != nil {
			return staff, metadata, errors.Wrap(err, "[SERVICE][GetProfileStaffSlim] Error getting paginated staff data")
		}

		return staff, metadata, nil
	}

	// Jika page atau length tidak valid, dapatkan semua data tanpa pagination
	staff, err := s.ppdb.GetProfileStaff(ctx, searchInput, offset, limit)
	if err != nil {
		return staff, metadata, errors.Wrap(err, "[SERVICE][GetProfileStaffSlim] Error getting staff data")
	}

	return staff, metadata, nil
}

func (s Service) DeleteProfileStaff(ctx context.Context, staffID string) (string, error) {
	result, err := s.ppdb.DeleteProfileStaff(ctx, staffID)

	if err != nil {
		return result, errors.Wrap(err, "[Service][DeleteProfileStaff]")
	}

	return result, nil
}

func (s Service) GetProfileStaffUtama(ctx context.Context) ([]ppdbEntity.TableStaff, error) {
	staff, err := s.ppdb.GetProfileStaffUtama(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "[SERVICE] [GetProfileStaffUtama]")
	}

	return staff, nil
}

func (s Service) UpdateProfileStaff(ctx context.Context, staff ppdbEntity.TableStaff, staffID string) (string, error) {
	var (
		result string
	)
	fmt.Println("masuk service")

	result, err := s.ppdb.UpdateProfileStaff(ctx, staff, staffID)
	if err != nil {
		result = "Gagal"
		return result, errors.Wrap(err, "[Service][UpdateProfileStaff]")
	}

	result = "Berhasil"
	return result, err
}

// Event
func (s Service) InsertEvent(ctx context.Context, event ppdbEntity.TableEvent) (string, error) {
	var (
		result string
	)

	// Panggil fungsi InsertEvent dari data layer
	result, err := s.ppdb.InsertEvent(ctx, event)

	if err != nil {
		result = "Gagal menyimpan data event"
		return result, errors.Wrap(err, "[Service][InsertEvent]")
	}

	result = "Berhasil menyimpan data event"
	return result, nil
}

func (s Service) GetImageEvent(ctx context.Context, eventID string) ([]byte, error) {
	poster, err := s.ppdb.GetImageEvent(ctx, eventID)
	if err != nil {
		return poster, errors.Wrap(err, "[SERVICE][GetImageEvent]")
	}

	return poster, err
}

func (s Service) GetEventSlim(ctx context.Context, searchInput string, page, length int) ([]ppdbEntity.TableEvent, interface{}, error) {
	limit := length
	offset := (page - 1) * length
	var lastPage int
	metadata := make(map[string]int)
	event := []ppdbEntity.TableEvent{}

	// Pagination
	if page > 0 && length > 0 {
		// Get total count of event for pagination
		count, err := s.ppdb.GetEventPagination(ctx, searchInput)
		if err != nil {
			return event, metadata, errors.Wrap(err, "[SERVICE][GetEventSlim] Error getting pagination count")
		}

		// Calculate last page based on count and length
		lastPage = int(math.Ceil(float64(count) / float64(length)))

		// Prepare metadata
		metadata["total_data"] = count
		metadata["total_page"] = lastPage

		// Get paginated event data
		event, err = s.ppdb.GetEvent(ctx, searchInput, offset, limit)
		if err != nil {
			return event, metadata, errors.Wrap(err, "[SERVICE][GetEventSlim] Error getting paginated event data")
		}

		return event, metadata, nil
	}

	// Jika page atau length tidak valid, dapatkan semua data tanpa pagination
	event, err := s.ppdb.GetEvent(ctx, searchInput, offset, limit)
	if err != nil {
		return event, metadata, errors.Wrap(err, "[SERVICE][GetEventSlim] Error getting event data")
	}

	return event, metadata, nil
}

func (s Service) GetEventDetail(ctx context.Context, eventID string) (ppdbEntity.TableEvent, error) {
	event, err := s.ppdb.GetEventDetail(ctx, eventID)
	if err != nil {
		return ppdbEntity.TableEvent{}, errors.Wrap(err, "[SERVICE] [GetEventDetail]")
	}

	return event, nil
}

func (s Service) GetEventUtama(ctx context.Context) ([]ppdbEntity.TableEvent, error) {
	events, err := s.ppdb.GetEventUtama(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "[SERVICE] [GetEventUtama]")
	}

	return events, nil
}

func (s Service) DeleteEvent(ctx context.Context, eventID string) (string, error) {
	result, err := s.ppdb.DeleteEvent(ctx, eventID)

	if err != nil {
		return result, errors.Wrap(err, "[Service][DeleteEvent]")
	}

	return result, nil
}

func (s Service) UpdateEvent(ctx context.Context, event ppdbEntity.TableEvent, eventID string) (string, error) {
	var (
		result string
	)

	result, err := s.ppdb.UpdateEvent(ctx, event, eventID)
	if err != nil {
		result = "Gagal"
		return result, errors.Wrap(err, "[Service][UpdateEvent]")
	}

	result = "Berhasil"
	return result, err
}

// Peserta Didik
func (s Service) InsertPesertaDidik(ctx context.Context, pesertadidik ppdbEntity.TablePesertaDidik) (string, error) {
	var (
		err    error
		result string
	)

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(pesertadidik.Password), bcrypt.DefaultCost)
	if err != nil {
		result = "Gagal menyimpan data peserta didik"
		return result, errors.Wrap(err, "[SERVICE][InsertUser][Hash]")
	}

	pesertadidik.Password = string(hashedPassword)

	result, err = s.ppdb.InsertPesertaDidik(ctx, pesertadidik)
	if err != nil {
		result = "Gagal menyimpan data peserta didik"
		return result, errors.Wrap(err, "[Service][InsertEvent]")
	}

	result = "Berhasil menyimpan data peserta didik"
	return result, nil
}

func (s Service) GetLoginCheck(ctx context.Context, login ppdbEntity.TablePesertaDidik) (ppdbEntity.TablePesertaDidik, error) {
	var (
		err    error
		result ppdb.TablePesertaDidik
	)

	result, err = s.ppdb.GetLoginCheck(ctx, login)
	if err != nil {
		return result, errors.Wrap(err, "[Service][GetLoginCheck]")
	}

	err = bcrypt.CompareHashAndPassword([]byte(result.Password), []byte(login.Password))
	if err != nil {
		return result, errors.Wrap(err, "[SERVICE][GetLoginCheck][CompareHash]")
	}

	return result, nil
}

func (s Service) GetPembayaranFormulirDetail(ctx context.Context, idpesertadidik string) (ppdbEntity.TablePembayaranFormulir, error) {

	pembayaranformulir, err := s.ppdb.GetPembayaranFormulirDetail(ctx, idpesertadidik)
	if err != nil {
		return pembayaranformulir, errors.Wrap(err, "[SERVICE][GetPembayaranFormulirDetail]")
	}

	return pembayaranformulir, nil
}

func (s Service) GetFormulirDetail(ctx context.Context, idpesertadidik string) (ppdbEntity.TableDataFormulir, error) {

	formulir, err := s.ppdb.GetFormulirDetail(ctx, idpesertadidik)
	if err != nil {
		return formulir, errors.Wrap(err, "[SERVICE][GetFormulirDetail]")
	}

	return formulir, nil
}

func (s Service) GetBerkasDetail(ctx context.Context, idpesertadidik string) (ppdbEntity.TableBerkas, error) {

	berkas, err := s.ppdb.GetBerkasDetail(ctx, idpesertadidik)
	if err != nil {
		return berkas, errors.Wrap(err, "[SERVICE][GetBerkasDetail]")
	}

	return berkas, nil
}

func (s Service) GetJadwalTestDetail(ctx context.Context, idpesertadidik string) (ppdbEntity.TableJadwalTest, error) {

	jadwaltest, err := s.ppdb.GetJadwalTestDetail(ctx, idpesertadidik)
	if err != nil {
		return jadwaltest, errors.Wrap(err, "[SERVICE][GetJadwalTestDetail]")
	}

	return jadwaltest, nil
}

func (s Service) UpdatePembayaranFormulir(ctx context.Context, pembayaranformulir ppdbEntity.TablePembayaranFormulir) (string, error) {
	var (
		err    error
		result string
	)

	result, err = s.ppdb.UpdatePembayaranFormulir(ctx, pembayaranformulir)
	if err != nil {
		result = "Gagal update data pembayaran formulir"
		return result, errors.Wrap(err, "[Service][UpdatePembayaranFormulir]")
	}

	result = "Berhasil update data pembayaran formulir"
	return result, nil
}

func (s Service) UpdateFormulir(ctx context.Context, formulir ppdbEntity.TableDataFormulir) (string, error) {
	var (
		err    error
		result string
	)

	result, err = s.ppdb.UpdateFormulir(ctx, formulir)
	if err != nil {
		result = "Gagal update data formulir"
		return result, errors.Wrap(err, "[Service][UpdateFormulir]")
	}

	result, err = s.ppdb.UpdateKontakPeserta(ctx, formulir)
	if err != nil {
		result = "Gagal update data kontak peserta"
		return result, errors.Wrap(err, "[Service][UpdateFormulir]")
	}

	result, err = s.ppdb.UpdateOrtu(ctx, formulir)
	if err != nil {
		result = "Gagal update data ortu"
		return result, errors.Wrap(err, "[Service][UpdateFormulir]")
	}

	result = "Berhasil update data formulir"
	return result, nil
}

func (s Service) UpdateBerkas(ctx context.Context, berkas ppdbEntity.TableBerkas) (string, error) {
	var (
		err    error
		result string
	)

	result, err = s.ppdb.UpdateBerkas(ctx, berkas)
	if err != nil {
		result = "Gagal update data berkas"
		return result, errors.Wrap(err, "[Service][UpdateBerkas]")
	}

	result = "Berhasil update data berkas"
	return result, nil
}

func (s Service) UpdateJadwalTest(ctx context.Context, jadwalTest ppdbEntity.TableJadwalTest) (string, error) {
	var (
		err    error
		result string
	)

	result, err = s.ppdb.UpdateJadwalTest(ctx, jadwalTest)
	if err != nil {
		result = "Gagal update data jadwal test"
		return result, errors.Wrap(err, "[Service][UpdateJadwalTest]")
	}

	result = "Berhasil update data jadwal test"
	return result, nil
}

func GetImageHeight(imagePath string, targetWidth float64) (float64, error) {
	file, err := os.Open(imagePath)
	if err != nil {
		return 0, fmt.Errorf("unable to open image: %w", err)
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		return 0, fmt.Errorf("unable to decode image: %w", err)
	}

	// Get image dimensions
	imgWidth := float64(img.Bounds().Dx())
	imgHeight := float64(img.Bounds().Dy())

	// Calculate height while maintaining aspect ratio
	aspectRatio := imgWidth / imgHeight
	height := targetWidth / aspectRatio

	return height, nil
}

func (s Service) GetGeneratedKartuTest(ctx context.Context, idpesertadidik string) ([]byte, error) {

	var (
		err error
	)

	docPdf := bytes.Buffer{}

	currentYear := time.Now().Year()
	nextYear := currentYear + 1

	currentYearString := strconv.Itoa(currentYear)
	nextYearString := strconv.Itoa(nextYear)

	pesertadidik, err := s.ppdb.GetPesertaDidikDetail(ctx, idpesertadidik)
	if err != nil {
		return docPdf.Bytes(), errors.Wrap(err, "[SERVICE][GeneratePDF]")
	}

	formulir, err := s.ppdb.GetFormulirDetail(ctx, idpesertadidik)
	if err != nil {
		return docPdf.Bytes(), errors.Wrap(err, "[SERVICE][GeneratePDF]")
	}

	berkas, err := s.ppdb.GetBerkasDetail(ctx, idpesertadidik)
	if err != nil {
		return docPdf.Bytes(), errors.Wrap(err, "[SERVICE][GeneratePDF]")
	}

	jadwaltest, err := s.ppdb.GetJadwalTestDetail(ctx, idpesertadidik)
	if err != nil {
		return docPdf.Bytes(), errors.Wrap(err, "[SERVICE][GeneratePDF]")
	}

	pdf := gofpdf.New("P", "mm", "A4", "")

	pdf.AddPage()
	pdf.SetFont("Arial", "B", 12)

	pdf.CellFormat(190, 20, "TANDA TERIMA PENDAFTARAN SISWA BARU SMA - TAHUN PELAJARAN "+currentYearString+"/"+nextYearString, "0", 0, "C", false, 0, "")
	pdf.Ln(18)

	pdf.Line(10, pdf.GetY(), 200, pdf.GetY())

	tmpFile, err := ioutil.TempFile("", "pasphoto-*.jpg")
	if err != nil {
		return docPdf.Bytes(), fmt.Errorf("error creating temp file for image: %v", err)
	}
	defer os.Remove(tmpFile.Name())

	_, err = tmpFile.Write(berkas.PasPhoto)
	if err != nil {
		return docPdf.Bytes(), fmt.Errorf("error writing image to temp file: %v", err)
	}

	err = tmpFile.Close()
	if err != nil {
		return docPdf.Bytes(), fmt.Errorf("error closing temp file: %v", err)
	}

	imageWidth := 50.0
	imageHeight, err := GetImageHeight(tmpFile.Name(), imageWidth)
	if err != nil {
		return docPdf.Bytes(), fmt.Errorf("error calculating image height: %v", err)
	}

	pdf.Ln(5)
	imageX := 10.0
	imageY := pdf.GetY()

	pdf.Image(tmpFile.Name(), imageX, imageY, imageWidth, imageHeight, true, "", 0, "")

	pdf.SetFont("Arial", "", 12)

	namaX := imageX + imageWidth + 10
	namaY := imageY + ((imageHeight - 27) / 2)

	pdf.SetY(namaY)
	pdf.SetX(namaX)

	pdf.CellFormat(40, 12, "Nama Lengkap", "", 0, "L", false, 0, "")
	radius := 2.0

	pdf.SetX(namaX + 42)
	pdf.CellFormat(60, 12, pesertadidik.PesertaName, "", 1, "L", false, 0, "")

	pdf.SetDrawColor(189, 189, 189)
	pdf.RoundedRectExt(namaX+40, namaY, 90, 12, radius, radius, radius, radius, "D")

	kelasX := namaX
	kelasY := namaY + 15

	pdf.SetY(kelasY)
	pdf.SetX(kelasX)
	pdf.CellFormat(40, 12, "Mendaftar Kelas", "", 0, "L", false, 0, "")

	pdf.SetX(kelasX + 42)
	pdf.CellFormat(60, 12, formulir.Kelas+" "+formulir.JurusanName, "", 1, "L", false, 0, "")

	pdf.RoundedRectExt(kelasX+40, kelasY, 90, 12, radius, radius, radius, radius, "D")

	pdf.SetY(imageY + imageHeight + 5)
	pdf.SetFont("Arial", "B", 12)
	pdf.CellFormat(40, 10, "CATATAN :", "", 1, "L", false, 0, "")
	pdf.SetFont("Arial", "", 12)
	pdf.CellFormat(190, 6, "- Tanda terima ini harap dicetak, ditanda-tangan oleh calon peserta didik dan dibawa saat test seleksi", "", 1, "L", false, 0, "")
	pdf.CellFormat(30, 6, "   pada tanggal", "", 0, "L", false, 0, "")
	pdf.SetFont("Arial", "B", 12)
	pdf.CellFormat(40, 6, jadwaltest.TglTest.Format("02-01-2006"), "", 1, "L", false, 0, "")
	pdf.SetFont("Arial", "", 12)
	pdf.CellFormat(190, 10, "- Tanda terima ini ditunjukkan pada saat tes seleksi", "", 1, "L", false, 0, "")

	err = pdf.Output(&docPdf)
	if err != nil {
		log.Fatalf("Error creating PDF: %s", err)
	}

	return docPdf.Bytes(), err
}

func (s Service) GetGeneratedFormulir(ctx context.Context) ([]byte, error) {

	var (
		err error
	)

	docPdf := bytes.Buffer{}

	currentYear := time.Now().Year()
	nextYear := currentYear + 1

	currentYearString := strconv.Itoa(currentYear)
	nextYearString := strconv.Itoa(nextYear)

	pdf := gofpdf.New("P", "mm", "A4", "")

	pdf.AddPage()
	pdf.SetFont("Arial", "", 9)

	imageY := pdf.GetY()
	pdf.SetY(imageY + 2.5)
	imageX := 10.0
	pdf.Image("public/images/logoKY.png", imageX, imageY, 0, 15, true, "", 0, "")

	headerY := 10.0
	headerX := imageX + 20
	pdf.SetY(headerY)
	pdf.SetX(headerX)

	pdf.CellFormat(40, 4, "SMA Kristen Yusuf", "", 1, "L", false, 0, "")
	pdf.SetX(headerX)
	pdf.CellFormat(40, 4, "Jl Arwana II No. 16 Jembatan Dua", "", 2, "L", false, 0, "")
	pdf.SetX(headerX)
	pdf.CellFormat(40, 4, "Jakarta Utara (14450)", "", 1, "L", false, 0, "")
	pdf.SetX(headerX)
	pdf.CellFormat(40, 4, "Telp: 021-6693111, 021-6682017", "", 1, "L", false, 0, "")
	pdf.SetX(headerX)
	pdf.CellFormat(40, 4, "WA: 0838-7000-4500", "", 1, "L", false, 0, "")

	lineY := headerY + 22

	pdf.SetLineWidth(0.5)
	pdf.SetDashPattern([]float64{1.5}, 1.5)
	pdf.Line(10, lineY, 200, lineY)

	pdf.SetFont("Arial", "B", 12)

	pdf.SetY(pdf.GetY()+1.5)

	pdf.Ln(3)

	pdf.CellFormat(190, 7, "FORMULIR PENDAFTARAN SISWA BARU SMA - TAHUN PELAJARAN "+currentYearString+"/"+nextYearString, "", 1, "C", false, 0, "")

	pdf.SetFont("Arial", "", 10)
	pdf.SetLineWidth(0.2)

	pdf.Ln(3)

	pdf.CellFormat(70, 2, "a. Nama (sesuai Akta Kelahiran)", "", 0, "L", false, 0, "")
	pdf.CellFormat(1, 2, ":", "", 1, "L", false, 0, "")
	pdf.Line(83, pdf.GetY(), 200, pdf.GetY())

	pdf.CellFormat(70, 15, "b. Jenis Kelamin", "", 0, "L", false, 0, "")
	pdf.CellFormat(40, 15, ": Laki-Laki / Perempuan", "0", 0, "L", false, 0, "")
	pdf.SetFont("Arial", "I", 10)
	pdf.CellFormat(40, 15, "(coret yang tidak perlu)", "0", 1, "L", false, 0, "")

	pdf.SetFont("Arial", "", 10)
	pdf.CellFormat(70, 2, "c. Tempat, Tanggal Lahir", "", 0, "L", false, 0, "")
	pdf.CellFormat(1, 2, ":", "", 1, "L", false, 0, "")
	pdf.Line(83, pdf.GetY(), 200, pdf.GetY())

	pdf.SetFont("Arial", "", 10)

	pdf.CellFormat(70, 15, "d. Akta Kelahiran", "", 0, "L", false, 0, "")
	aktaLahirY := pdf.GetY()+9
	pdf.CellFormat(60, 15, ": No.", "", 0, "L", false, 0, "")
	pdf.Line(88, aktaLahirY, 138, aktaLahirY)
	pdf.CellFormat(60, 15, "Tanggal", "", 1, "L", false, 0, "")
	pdf.Line(155, aktaLahirY, 200, aktaLahirY)

	pdf.CellFormat(70, 2, "e. NISN (Nomor Induk Siswa Nasional)", "", 0, "L", false, 0, "")
	pdf.CellFormat(1, 2, ":", "", 1, "L", false, 0, "")
	pdf.Line(83, pdf.GetY(), 200, pdf.GetY())

	agamaY := pdf.GetY() + 9
	pdf.CellFormat(70, 15, "f. Agama", "", 0, "L", false, 0, "")
	pdf.CellFormat(1, 15, ":", "", 1, "L", false, 0, "")
	pdf.Line(83, agamaY, 200, agamaY)
	// pdf.Line(83, pdf.GetY()+2, 200, pdf.GetY()+2)

	// pdf.Line(10, pdf.GetY(), 200, pdf.GetY())

	// tmpFile, err := ioutil.TempFile("", "pasphoto-*.jpg")
	// if err != nil {
	// 	return docPdf.Bytes(), fmt.Errorf("error creating temp file for image: %v", err)
	// }
	// defer os.Remove(tmpFile.Name())

	// _, err = tmpFile.Write(berkas.PasPhoto)
	// if err != nil {
	// 	return docPdf.Bytes(), fmt.Errorf("error writing image to temp file: %v", err)
	// }

	// err = tmpFile.Close()
	// if err != nil {
	// 	return docPdf.Bytes(), fmt.Errorf("error closing temp file: %v", err)
	// }

	// imageWidth := 50.0
	// imageHeight, err := GetImageHeight(tmpFile.Name(), imageWidth)
	// if err != nil {
	// 	return docPdf.Bytes(), fmt.Errorf("error calculating image height: %v", err)
	// }

	// pdf.Ln(5)
	// imageX := 10.0
	// imageY := pdf.GetY()

	// pdf.Image(tmpFile.Name(), imageX, imageY, imageWidth, imageHeight, true, "", 0, "")

	// pdf.SetFont("Arial", "", 12)

	// namaX := imageX + imageWidth + 10
	// namaY := imageY + ((imageHeight - 27) / 2)

	// pdf.SetY(namaY)
	// pdf.SetX(namaX)

	// pdf.CellFormat(40, 12, "Nama Lengkap", "", 0, "L", false, 0, "")
	// radius := 2.0

	// pdf.SetX(namaX + 42)
	// pdf.CellFormat(60, 12, pesertadidik.PesertaName, "", 1, "L", false, 0, "")

	// pdf.SetDrawColor(189, 189, 189)
	// pdf.RoundedRectExt(namaX+40, namaY, 90, 12, radius, radius, radius, radius, "D")

	// kelasX := namaX
	// kelasY := namaY + 15

	// pdf.SetY(kelasY)
	// pdf.SetX(kelasX)
	// pdf.CellFormat(40, 12, "Mendaftar Kelas", "", 0, "L", false, 0, "")

	// pdf.SetX(kelasX + 42)
	// pdf.CellFormat(60, 12, formulir.Kelas+" "+formulir.JurusanName, "", 1, "L", false, 0, "")

	// pdf.RoundedRectExt(kelasX+40, kelasY, 90, 12, radius, radius, radius, radius, "D")

	// pdf.SetY(imageY + imageHeight + 5)
	// pdf.SetFont("Arial", "B", 12)
	// pdf.CellFormat(40, 10, "CATATAN :", "", 1, "L", false, 0, "")
	// pdf.SetFont("Arial", "", 12)
	// pdf.CellFormat(190, 6, "- Tanda terima ini harap dicetak, ditanda-tangan oleh calon peserta didik dan dibawa saat test seleksi", "", 1, "L", false, 0, "")
	// pdf.CellFormat(30, 6, "   pada tanggal", "", 0, "L", false, 0, "")
	// pdf.SetFont("Arial", "B", 12)
	// pdf.CellFormat(40, 6, jadwaltest.TglTest.Format("02-01-2006"), "", 1, "L", false, 0, "")
	// pdf.SetFont("Arial", "", 12)
	// pdf.CellFormat(190, 10, "- Tanda terima ini ditunjukkan pada saat tes seleksi", "", 1, "L", false, 0, "")

	err = pdf.Output(&docPdf)
	if err != nil {
		log.Fatalf("Error creating PDF: %s", err)
	}

	return docPdf.Bytes(), err
}
