package ppdb

import (
	"context"
	"log"

	"github.com/jmoiron/sqlx"
	"github.com/opentracing/opentracing-go"

	jaegerLog "ppdb-be/pkg/log"
)

type (
	// Data ...
	Data struct {
		db   *sqlx.DB
		stmt *map[string]*sqlx.Stmt

		tracer opentracing.Tracer
		logger jaegerLog.Factory
	}

	// statement ...
	statement struct {
		key   string
		query string
	}
)

const (
	//query get

	getDataAdminPagination  = "GetDataAdminPagination"
	qGetDataAdminPagination = `SELECT  count(*)
                    FROM T_Admin AS a JOIN T_Role AS r ON a.roleID = r.roleID WHERE 
                    a.adminName LIKE ?`

	getKontakSekolah  = "GetKontakSekolah"
	qGetKontakSekolah = `Select kontakKYID, alamatSekolah, noTelpSekolah1, noTelpSekolah2, emailSekolah, instagramSekolah FROM T_KontakSekolah`

	loginAdmin  = "LoginAdmin"
	qLoginAdmin = `Select adminID, roleID, adminName, emailAdmin, password FROM T_Admin WHERE emailAdmin = ? `

	getDataAdmin  = "GetDataAdmin"
	qGetDataAdmin = `SELECT  a.adminID, a.adminName, a.emailAdmin, r.roleID, r.roleName, r.roleDesc
                    FROM T_Admin AS a JOIN T_Role AS r ON a.roleID = r.roleID WHERE 
                    a.adminName LIKE ? LIMIT ?, ?`

	getLastAdminId  = "GetLastAdminId"
	qGetLastAdminId = `SELECT adminID FROM T_Admin
						ORDER BY adminID DESC
					   LIMIT 1`

	getRole  = "GetRole"
	qGetRole = `Select roleID, roleName, roleDesc FROM T_Role`

	getLastInfoId  = "GetLastInfoId"
	qGetLastInfoId = `SELECT infoID FROM T_InfoPendaftaran ORDER BY infoID DESC LIMIT 1`

	getGambarInfoDaftar  = "GetGambarInfodaftar"
	qGetGambarInfoDaftar = `SELECT posterDaftar from T_InfoPendaftaran WHERE infoID= ?`

	getInfoDaftar  = "GetInfoDaftar"
	qGetInfoDaftar = `SELECT infoID, posterDaftar, awalTahunAjar, akhirTahunAjar from T_InfoPendaftaran`

	getLastBannerId  = "GetLastBannerId"
	qGetLastBannerId = `SELECT bannerID FROM T_BannerSekolah ORDER BY bannerID DESC LIMIT 1`

	getGambarBanner  = "GetGambarBanner"
	qGetGambarBanner = `SELECT bannerImage from T_BannerSekolah WHERE bannerID =?`

	getBanner  = "GetBanner"
	qGetBanner = `SELECT bannerID, bannerName, bannerImage from T_BannerSekolah `

	getLastFasilitasId  = "GetLastFasilitasId"
	qGetLastFasilitasId = `SELECT fasilitasID FROM T_Fasilitas ORDER BY fasilitasID DESC LIMIT 1`

	getGambarFasilitas  = "GetGambarFasilitas"
	qGetGambarFasilitas = `SELECT fasilitasImage from T_Fasilitas WHERE fasilitasID =?`

	getFasilitas  = "GetFasilitas"
	qGetFasilitas = `SELECT fasilitasID, fasilitasName, fasilitasImage from T_Fasilitas WHERE 
                     fasilitasName LIKE ? LIMIT ?, ? `

	getFasilitasPagination  = "GetFasilitasPagination"
	qGetFasilitasPagination = `SELECT  count(*) FROM T_Fasilitas WHERE fasilitasName LIKE ?`

	getFasilitasUtama  = "GetFasilitasUtama"
	qGetFasilitasUtama = `SELECT fasilitasID, fasilitasName, fasilitasImage from T_Fasilitas`

	getLastStaffId  = "GetLastStaffId"
	qGetLastStaffId = `SELECT staffID FROM T_ProfileStaff ORDER BY staffID DESC LIMIT 1`

	getPhotoStaff  = "GetPhotoStaff"
	qGetPhotoStaff = `SELECT staffPhoto FROM T_ProfileStaff WHERE staffID = ?`

	getProfilStaff   = "GetProfilStaff"
	qGetProfileStaff = `SELECT staffID, staffName, staffGender, staffPosition, staffTmptLahir, staffTglLahir, staffPhoto
						FROM T_ProfileStaff WHERE staffName LIKE ? LIMIT ?, ?`

	getProfilStaffPagination  = "GetProfilStaffPagination"
	qGetProfilStaffPagination = `SELECT count(*) FROM T_ProfileStaff WHERE staffName LIKE ? `

	getProfilStaffUtama   = "GetProfilStaffUtama"
	qGetProfileStaffUtama = `SELECT staffID, staffName, staffGender, staffPosition, staffTmptLahir, staffTglLahir, staffPhoto
						FROM T_ProfileStaff`

	getLastEventId  = "GetLastEventId"
	qGetLastEventId = `SELECT eventID from T_EventSekolah ORDER BY eventID DESC LIMIT 1`

	getImageEvent  = "GetImageEvent"
	qGetImageEvent = `SELECT eventImage FROM T_EventSekolah WHERE eventID = ?`

	getEvent  = "GetEvent"
	qGetEvent = `SELECT eventID, eventHeader, eventStartDate, eventEndDate, eventDesc, eventImage 
				FROM T_EventSekolah WHERE eventHeader LIKE ? LIMIT ?,?`

	getEventPagination  = "GetEventPagination"
	qGetEventPagination = `SELECT count(*) FROM T_EventSekolah WHERE eventHeader LIKE ?`

	getEventDetail  = "GetEventDetail"
	qGetEventDetail = `SELECT eventID, eventHeader, eventStartDate, eventEndDate, eventDesc, eventImage 
				FROM T_EventSekolah WHERE eventID LIKE ?`

	getEventUtama  = "GetEventUtama"
	qGetEventUtama = `SELECT eventID, eventHeader, eventStartDate, eventEndDate, eventDesc, eventImage 
					FROM T_EventSekolah`

	getStatus  = "GetStatus"
	qGetStatus = `SELECT statusID, statusName, statusDesc FROM T_Status`

	getLastPesertaDidikId  = "GetLastPesertaDidikId"
	qGetLastPesertaDidikId = `SELECT pesertaID
								FROM T_PesertaDidik
								ORDER BY pesertaID DESC LIMIT 1`

	getLastPembayaranFormulirId  = "GetLastPembayaranFormulirId"
	qGetLastPembayaranFormulirId = `SELECT pembayaranID
								FROM T_PembayaranFormulir
								ORDER BY  pembayaranID DESC LIMIT 1`

	getPembayaranFormulirDetail  = "GetPembayaranFormulirDetail"
	qGetPembayaranFormulirDetail = `SELECT pembayaranID, pesertaID, statusID, 
										IFNULL(CAST(tglPembayaran AS DATE), "0001-01-01") AS tglPembayaran, hargaFormulir, buktiPembayaran
									FROM T_PembayaranFormulir
									WHERE pesertaID = ?`

	getLastFormulirId  = "GetLastFormulirId"
	qGetLastFormulirId = `SELECT formulirID
							FROM T_Formulir
							ORDER BY  formulirID DESC LIMIT 1`

	getLastKontakPesertaId  = "GetLastKontakPesertaId"
	qGetLastKontakPesertaId = `SELECT kontakID
								FROM T_KontakPeserta
								ORDER BY  kontakID DESC LIMIT 1`

	getLastOrtuId  = "GetLastOrtuId"
	qGetLastOrtuId = `SELECT ortuID
						FROM T_Ortu
						ORDER BY  ortuID DESC LIMIT 1`

	getFormulirDetail  = "GetFormulirDetail"
	qGetFormulirDetail = `SELECT f.formulirID, pesertaID, pembayaranID, jurusanID, agamaID, genderPeserta, tempatLahir, 
							IFNULL(CAST(tglLahir AS DATE), '0001-01-01') AS tglLahir, 
							NISN, Kelas, tglSubmit, statusID, kontakID, alamatTerakhir, kodePos, noTelpRumah,
							ortuID, namaAyah, pekerjaanAyah, noTelpHpAyah, namaIbu, pekerjaanIbu, noTelpHpIbu, namaWali, pekerjaanWali, noTelpHpWali
						FROM T_Formulir f
							JOIN T_KontakPeserta kp ON f.formulirID = kp.formulirID
							JOIN T_Ortu o ON f.formulirID = o.formulirID
						WHERE pesertaID = ?`

	getLastBerkasId  = "GetLastBerkasId"
	qGetLastBerkasId = `SELECT berkasID
						FROM T_Berkas
						ORDER BY berkasID DESC LIMIT 1`

	getBerkasDetail  = "GetBerkasDetail"
	qGetBerkasDetail = `SELECT berkasID, pesertaID, statusID, aktaLahir, pasPhoto, rapor, 
							IFNULL(CAST(tanggalUpload AS DATE), '0001-01-01') AS tanggalUpload
						FROM T_Berkas
						WHERE pesertaID = ?`

	getLastJadwalTestId  = "GetLastJadwalTestId"
	qGetLastJadwalTestId = `SELECT testID
							FROM T_JadwalTest
							ORDER BY testID DESC LIMIT 1`

	getJadwalTestDetail  = "GetJadwalTestDetail"
	qGetJadwalTestDetail = `SELECT testID, pesertaID, statusID, IFNULL(CAST(tglTest AS DATE), '0001-01-01') AS tglTest,  
								IFNULL(CAST(waktuTest AS TIME), '0001-01-01 00:00:00') AS waktuTest
							FROM T_JadwalTest
							WHERE pesertaID = ?`

	getLoginCheck  = "GetLoginCheck"
	qGetLoginCheck = `SELECT pesertaID, emailPeserta, password
						FROM T_PesertaDidik
						WHERE emailPeserta = ?`

	//query insert
	insertDataAdmin  = "InsertDataAdmin"
	qInsertDataAdmin = `INSERT INTO T_Admin (adminID, roleID, adminName, password, emailAdmin)
						VALUES (?, ?, ?, ?, ?)`

	insertInfoDaftar  = "InsertInfoDaftar"
	qInsertInfoDaftar = `INSERT INTO T_InfoPendaftaran (infoID, posterDaftar, awalTahunAjar, akhirTahunAjar)
						VALUES (?, ?, ?, ?)`

	insertBanner  = "InsertBanner"
	qInsertBanner = `INSERT INTO T_BannerSekolah (bannerID, bannerName, bannerImage)
					VALUES (?, ?, ?)`

	insertFasilitas  = "InsertFasilitas"
	qInsertFasilitas = `INSERT INTO T_Fasilitas (fasilitasID, fasilitasName, fasilitasImage)
						VALUES (?, ?, ?)`

	insertProfileStaff  = "InsertProfileStaff"
	qInsertProfileStaff = `INSERT INTO T_ProfileStaff (staffID, staffName, staffGender, staffPosition, 
						staffTmptLahir, staffTglLahir, staffPhoto ) VALUES (?, ?, ?, ?, ?, ?, ?)`

	insertEvent  = "InsertEvent"
	qInsertEvent = `INSERT INTO T_EventSekolah (eventID, eventHeader, eventStartDate, eventEndDate, eventDesc, eventImage)
					VALUES(?, ?, ?, ?, ?, ?)`

	insertPesertaDidik  = "InsertPesertaDidik"
	qInsertPesertaDidik = `INSERT INTO T_PesertaDidik
								(pesertaID, pesertaName, password, emailPeserta, noTelpHpPeserta, sekolahAsalYN, sekolahAsal, alamatSekolahAsal)
							VALUES(?, ?, ?, ?, ?, ?, ?, ?)`

	insertPembayaranFormulir  = "InsertPembayaranFormulir"
	qInsertPembayaranFormulir = `INSERT INTO T_PembayaranFormulir
									(pembayaranID, pesertaID, statusID, tglPembayaran, hargaFormulir, buktiPembayaran)
								VALUES(?, ?, ?, ?, ?, "")`

	insertFormulir  = "InsertFormulir"
	qInsertFormulir = `INSERT INTO u868654674_ppdb.T_Formulir
							(formulirID, pesertaID, pembayaranID, jurusanID, agamaID, genderPeserta, tempatLahir, tglLahir, NISN, Kelas, tglSubmit, statusID)
						VALUES(?, ?, ?, "", "", "", "", ?, "", "", ?, ?)`

	insertKontakPeserta  = "InsertKontakPeserta"
	qInsertKontakPeserta = `INSERT INTO T_KontakPeserta
								(kontakID, formulirID, alamatTerakhir, kodePos, noTelpRumah)
							VALUES(?, ?, "", "", "")`

	insertOrtu  = "InsertOrtu"
	qInsertOrtu = `INSERT INTO T_Ortu
						(ortuID, formulirID, namaAyah, pekerjaanAyah, noTelpHpAyah, namaIbu, pekerjaanIbu, noTelpHpIbu, 
						namaWali, pekerjaanWali, noTelpHpWali)
					VALUES(?, ?, "", "", "", "", "", "", "", "", "");`

	insertBerkas  = "InsertBerkas"
	qInsertBerkas = `INSERT INTO u868654674_ppdb.T_Berkas
						(berkasID, pesertaID, statusID, aktaLahir, pasPhoto, rapor, tanggalUpload)
					VALUES(?, ?, ?, "", "", "", ?);`

	insertJadwalTest  = "InsertJadwalTest"
	qInsertJadwalTest = `INSERT INTO u868654674_ppdb.T_JadwalTest
							(testID, pesertaID, statusID, tglTest, waktuTest)
						VALUES(?, ?, ?, ?, CAST(? AS TIME));`

	//query delete
	deleteDataAdmin  = "DeleteDataAdmin"
	qDeleteDataAdmin = `DELETE FROM T_Admin WHERE adminID = ?`

	deleteBanner  = "DeleteBanner"
	qDeleteBanner = `DELETE FROM T_BannerSekolah WHERE bannerID = ?`

	deleteFasilitas  = "DeleteFasilitas"
	qDeleteFasilitas = `DELETE FROM T_Fasilitas WHERE fasilitasID = ?`

	deleteProfileStaff  = "DeleteProfileStaff"
	qDeleteProfileStaff = `DELETE FROM T_ProfileStaff WHERE staffID = ?	`

	deleteEvent  = "DeleteEvent"
	qDeleteEvent = `DELETE FROM T_EventSekolah WHERE eventID = ?`

	//query update

	updateBanner  = "UpdateBanner"
	qUpdateBanner = `UPDATE T_BannerSekolah
					SET bannerName = ?, bannerImage = ?
					WHERE bannerID = ?`

	updateInfoDaftar  = "UpdateInfoDaftar"
	qUpdateInfoDaftar = `UPDATE T_InfoPendaftaran 
		SET posterDaftar = ?, awalTahunAjar = ?, akhirTahunAjar = ? 
		WHERE infoID = ?`

	updateFasilitas  = "UpdateFasilitas"
	qUpdateFasilitas = `UPDATE T_Fasilitas 
							SET fasilitasName = ?, fasilitasImage = ?
							WHERE fasilitasID = ?`

	updateProfileStaff  = "UpdateProfileStaff"
	qUpdateProfileStaff = `UPDATE T_ProfileStaff 
						 SET staffName = ?, staffGender = ?, staffPosition = ?, 
						 staffTmptLahir = ?, staffTglLahir = ?, staffPhoto = ?
						WHERE staffID = ?`

	updateEvent  = "UpdateEvent"
	qUpdateEvent = `UPDATE T_EventSekolah 
					SET eventHeader = ?, eventStartDate = ?, eventEndDate = ?, eventDesc = ?, eventImage = ?
					WHERE eventID = ?`
)

var (
	readStmt = []statement{
		{loginAdmin, qLoginAdmin},
		{getKontakSekolah, qGetKontakSekolah},
		{getDataAdmin, qGetDataAdmin},
		{getDataAdminPagination, qGetDataAdminPagination},
		{getLastAdminId, qGetLastAdminId},
		{getLastInfoId, qGetLastInfoId},
		{getLastBannerId, qGetLastBannerId},

		{getGambarInfoDaftar, qGetGambarInfoDaftar},
		{getInfoDaftar, qGetInfoDaftar},

		{getGambarBanner, qGetGambarBanner},
		{getBanner, qGetBanner},

		{getRole, qGetRole},
		{getStatus, qGetStatus},

		{getLastFasilitasId, qGetLastFasilitasId},
		{getGambarFasilitas, qGetGambarFasilitas},
		{getFasilitas, qGetFasilitas},
		{getFasilitasPagination, qGetFasilitasPagination},
		{getFasilitasUtama, qGetFasilitasUtama},

		{getLastStaffId, qGetLastStaffId},
		{getPhotoStaff, qGetPhotoStaff},
		{getProfilStaff, qGetProfileStaff},
		{getProfilStaffPagination, qGetProfilStaffPagination},
		{getProfilStaffUtama, qGetProfileStaffUtama},

		{getLastEventId, qGetLastEventId},
		{getImageEvent, qGetImageEvent},
		{getEvent, qGetEvent},
		{getEventPagination, qGetEventPagination},
		{getEventDetail, qGetEventDetail},
		{getEventUtama, qGetEventUtama},

		{getLastPesertaDidikId, qGetLastPesertaDidikId},

		{getLastPembayaranFormulirId, qGetLastPembayaranFormulirId},
		{getPembayaranFormulirDetail, qGetPembayaranFormulirDetail},

		{getLastFormulirId, qGetLastFormulirId},
		{getLastKontakPesertaId, qGetLastKontakPesertaId},
		{getLastOrtuId, qGetLastOrtuId},
		{getFormulirDetail, qGetFormulirDetail},

		{getLastBerkasId, qGetLastBerkasId},
		{getBerkasDetail, qGetBerkasDetail},

		{getLastJadwalTestId, qGetLastJadwalTestId},
		{getJadwalTestDetail, qGetJadwalTestDetail},

		{getLoginCheck, qGetLoginCheck},
	}
	insertStmt = []statement{
		{insertDataAdmin, qInsertDataAdmin},
		{insertInfoDaftar, qInsertInfoDaftar},
		{insertBanner, qInsertBanner},
		{insertFasilitas, qInsertFasilitas},
		{insertProfileStaff, qInsertProfileStaff},
		{insertEvent, qInsertEvent},
		{insertPesertaDidik, qInsertPesertaDidik},
		{insertPembayaranFormulir, qInsertPembayaranFormulir},
		{insertFormulir, qInsertFormulir},
		{insertKontakPeserta, qInsertKontakPeserta},
		{insertOrtu, qInsertOrtu},
		{insertBerkas, qInsertBerkas},
		{insertJadwalTest, qInsertJadwalTest},
	}
	updateStmt = []statement{
		{updateBanner, qUpdateBanner},
		{updateInfoDaftar, qUpdateInfoDaftar},
		{updateFasilitas, qUpdateFasilitas},
		{updateProfileStaff, qUpdateProfileStaff},
		{updateEvent, qUpdateEvent},
	}
	deleteStmt = []statement{
		{deleteDataAdmin, qDeleteDataAdmin},
		{deleteBanner, qDeleteBanner},
		{deleteFasilitas, qDeleteFasilitas},
		{deleteProfileStaff, qDeleteProfileStaff},
		{deleteEvent, qDeleteEvent},
	}
)

// New ...
func New(db *sqlx.DB, tracer opentracing.Tracer, logger jaegerLog.Factory) *Data {
	var (
		stmts = make(map[string]*sqlx.Stmt)
	)
	d := &Data{
		db:     db,
		tracer: tracer,
		logger: logger,
		stmt:   &stmts,
	}

	d.InitStmt()
	return d
}

func (d *Data) InitStmt() {
	var (
		err   error
		stmts = make(map[string]*sqlx.Stmt)
	)

	for _, v := range readStmt {
		stmts[v.key], err = d.db.PreparexContext(context.Background(), v.query)
		if err != nil {
			log.Fatalf("[DB] Failed to initialize select statement key %v, err : %v", v.key, err)
		}
	}

	for _, v := range insertStmt {
		stmts[v.key], err = d.db.PreparexContext(context.Background(), v.query)
		if err != nil {
			log.Fatalf("[DB] Failed to initialize insert statement key %v, err : %v", v.key, err)
		}
	}

	for _, v := range updateStmt {
		stmts[v.key], err = d.db.PreparexContext(context.Background(), v.query)
		if err != nil {
			log.Fatalf("[DB] Failed to initialize update statement key %v, err : %v", v.key, err)
		}
	}

	for _, v := range deleteStmt {
		stmts[v.key], err = d.db.PreparexContext(context.Background(), v.query)
		if err != nil {
			log.Fatalf("[DB] Failed to initialize delete statement key %v, err : %v", v.key, err)
		}
	}

	*d.stmt = stmts
}
