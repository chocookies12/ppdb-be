package ppdb

import (
	"context"
	"math"

	// ppdbEntity "ppdb-be/internal/entity/ppdb"
	ppdbEntity "ppdb-be/internal/entity/ppdb"
	"ppdb-be/pkg/errors"
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