package ppdb

import (
	"context"
	// ppdbEntity "ppdb-be/internal/entity/ppdb"
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

func (s Service) LoginAdmin(ctx context.Context, admin_id string, admin_password string) (string, error) {
	var (
		result string
	)

	result, err := s.ppdb.LoginAdmin(ctx, admin_id, admin_password)

	if err != nil {
		result = "Gagal Login"
		return result, errors.Wrap(err, "[Service][LoginAdmin]")
	}
	result = "Berhasil Login"

	return result, err

}
