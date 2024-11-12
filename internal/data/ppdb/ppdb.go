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

	getAgama  = "GetAgama"
	qGetAgama = `SELECT agamaID, agamaName FROM T_Agama`

	getJurusan  = "GetJurusan"
	qGetJurusan = `SELECT jurusanID, jurusanName FROM T_Jurusan`

	getLastInfoId  = "GetLastInfoId"
	qGetLastInfoId = `SELECT infoID FROM T_InfoPendaftaran ORDER BY infoID DESC LIMIT 1`

	getGambarInfoDaftar  = "GetGambarInfodaftar"
	qGetGambarInfoDaftar = `SELECT posterDaftar from T_InfoPendaftaran WHERE infoID= ?`

	getInfoDaftar  = "GetInfoDaftar"
	qGetInfoDaftar = `SELECT infoID, posterDaftar, awalTahunAjar, akhirTahunAjar, noRekening, namaBank,
					 pemilikRekening from T_InfoPendaftaran`

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

	getPesertaDidikDetail  = "GetPesertaDidikDetail"
	qGetPesertaDidikDetail = `SELECT pesertaID, pesertaName, password, emailPeserta, noTelpHpPeserta, sekolahAsalYN, sekolahAsal, alamatSekolahAsal
								FROM T_PesertaDidik
								WHERE pesertaID = ?`

	getLastPembayaranFormulirId  = "GetLastPembayaranFormulirId"
	qGetLastPembayaranFormulirId = `SELECT pembayaranID
								FROM T_PembayaranFormulir
								ORDER BY  pembayaranID DESC LIMIT 1`

	getPembayaranFormulirDetail  = "GetPembayaranFormulirDetail"
	qGetPembayaranFormulirDetail = `SELECT 
										p.pembayaranID, 
										p.pesertaID, 
										pd.pesertaName,
										p.statusID, 
										s.statusName,
										IFNULL(CAST(p.tglPembayaran AS DATE), "0001-01-01") AS tglPembayaran, 
										p.hargaFormulir, 
										p.buktiPembayaran
									FROM 
										T_PembayaranFormulir p
									JOIN 
										T_Status s ON p.statusID = s.statusID
									JOIN 
										T_PesertaDidik pd ON p.pesertaID = pd.pesertaID
									WHERE 
										p.pesertaID = ?`

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
	qGetFormulirDetail = `SELECT 
						f.formulirID, 
						f.pesertaID, 
						f.pembayaranID, 
						f.jurusanID, 
						f.agamaID, 
						f.genderPeserta, 
						f.noAktaLahir,
						f.tempatLahir, 
						IFNULL(CAST(f.tglLahir AS DATE), '0001-01-01') AS tglLahir, 
						f.NISN, 
						f.Kelas, 
						f.urutanAnak, 
						f.jumlahSaudara,
						f.tglSubmit, 
						s.statusID, 
						s.statusName, 
						kp.kontakID, 
						kp.alamatTerakhir, 
						kp.kodePos, 
						kp.noTelpRumah, 
						o.ortuID, 
						o.namaAyah, 
						o.pekerjaanAyah, 
						o.noTelpHpAyah, 
						o.namaIbu, 
						o.pekerjaanIbu, 
						o.noTelpHpIbu, 
						o.namaWali, 
						o.pekerjaanWali, 
						o.noTelpHpWali,
						pd.pesertaName,
						pd.noTelpHpPeserta,
						pd.sekolahAsal,
						pd.alamatSekolahAsal,
						IFNULL(j.jurusanName, '') AS jurusanName,    
						IFNULL(a.agamaName, '') AS agamaName         
					FROM 
						T_Formulir f
					JOIN 
						T_KontakPeserta kp ON f.formulirID = kp.formulirID
					JOIN 
						T_Ortu o ON f.formulirID = o.formulirID
					LEFT JOIN 
						T_Status s ON f.statusID = s.statusID
					JOIN 
						T_PesertaDidik pd ON f.pesertaID = pd.pesertaID
					LEFT JOIN 
						T_Jurusan j ON f.jurusanID = j.jurusanID
					LEFT JOIN 
						T_Agama a ON f.agamaID = a.agamaID          
					WHERE 
						f.pesertaID = ?`

	getLastBerkasId  = "GetLastBerkasId"
	qGetLastBerkasId = `SELECT berkasID
						FROM T_Berkas
						ORDER BY berkasID DESC LIMIT 1`

	getBerkasDetail  = "GetBerkasDetail"
	qGetBerkasDetail = `SELECT 
						b.berkasID,
						b.pesertaID,
						pd.pesertaName,
						b.statusID,
						s.statusName,
						b.aktaLahir,
						b.pasPhoto,
						b.rapor,
						IFNULL(CAST(b.tanggalUpload AS DATE), '0001-01-01') AS tanggalUpload
					FROM 
						T_Berkas b
					JOIN 
						T_Status s ON b.statusID = s.statusID
					JOIN 
						T_PesertaDidik pd ON b.pesertaID = pd.pesertaID
					WHERE 
						b.pesertaID = ?	`

	getLastJadwalTestId  = "GetLastJadwalTestId"
	qGetLastJadwalTestId = `SELECT testID
							FROM T_JadwalTest
							ORDER BY testID DESC LIMIT 1`

	getJadwalTestDetail  = "GetJadwalTestDetail"
	qGetJadwalTestDetail = `SELECT 
							jt.testID,
							jt.pesertaID,
							pd.pesertaName,
							jt.statusID,
							s.statusName,
							IFNULL(CAST(jt.tglTest AS DATE), '0001-01-01') AS tglTest,  
							IFNULL(CAST(jt.waktuTest AS TIME), '00:00:00') AS waktuTest
						FROM 
							T_JadwalTest jt
						JOIN 
							T_Status s ON jt.statusID = s.statusID
						JOIN 
							T_PesertaDidik pd ON jt.pesertaID = pd.pesertaID
						WHERE 
							jt.pesertaID = ?`

	getLoginCheck  = "GetLoginCheck"
	qGetLoginCheck = `SELECT pesertaID, pesertaName, sekolahAsalYN, emailPeserta, password
						FROM T_PesertaDidik
						WHERE emailPeserta = ?`

	getJadwalTestAll  = "GetJadwalTestAll"
	qGetJadwaltestAll = `SELECT 
						jt.testID,
						jt.pesertaID,
						pd.pesertaName,
						jt.statusID,
						s.statusName,
						IFNULL(CAST(jt.tglTest AS DATE), '0001-01-01') AS tglTest,  
						IFNULL(CAST(jt.waktuTest AS TIME), '00:00:00') AS waktuTest
					FROM 
						T_JadwalTest jt
					JOIN 
						T_Status s ON jt.statusID = s.statusID
					JOIN 
						T_PesertaDidik pd ON jt.pesertaID = pd.pesertaID
					WHERE 
						pd.pesertaName LIKE ? LIMIT ?, ?`

	getJadwalTestPagination  = "GetJadwalTestPagination"
	qGetJadwalTestPagination = `SELECT 
                              COUNT(*) AS totalCount
                            FROM 
                              T_JadwalTest jt
                            JOIN 
                              T_Status s ON jt.statusID = s.statusID
                            JOIN 
                              T_PesertaDidik pd ON jt.pesertaID = pd.pesertaID
                            WHERE 
                              pd.pesertaName LIKE ?`

	getPembayaranFormulirAll  = "GetPembayaranFormulirAll"
	qGetPembayaranFormulirAll = `SELECT 
								p.pembayaranID, 
								p.pesertaID, 
								pd.pesertaName,
								p.statusID, 
								s.statusName,
								IFNULL(CAST(p.tglPembayaran AS DATE), '0001-01-01') AS tglPembayaran, 
								p.hargaFormulir, 
								p.buktiPembayaran
							FROM 
								T_PembayaranFormulir p
							JOIN 
								T_Status s ON p.statusID = s.statusID
							JOIN 
								T_PesertaDidik pd ON p.pesertaID = pd.pesertaID
							WHERE 
								pd.pesertaName LIKE ? LIMIT ?, ?`

	getPembayaranFormulirPagination  = "GetPembayaranFormulirPagination"
	qGetPembayaranFormulirPagination = `SELECT 
                                        COUNT(*) AS totalCount
                                    FROM 
                                        T_PembayaranFormulir p
                                    JOIN 
                                        T_Status s ON p.statusID = s.statusID
                                    JOIN 
                                        T_PesertaDidik pd ON p.pesertaID = pd.pesertaID
                                    WHERE 
                                        pd.pesertaName LIKE ?`

	//query insert
	insertDataAdmin  = "InsertDataAdmin"
	qInsertDataAdmin = `INSERT INTO T_Admin (adminID, roleID, adminName, password, emailAdmin)
						VALUES (?, ?, ?, ?, ?)`

	insertInfoDaftar  = "InsertInfoDaftar"
	qInsertInfoDaftar = `INSERT INTO T_InfoPendaftaran (infoID, posterDaftar, awalTahunAjar, akhirTahunAjar, noRekening, namaBank, pemilikRekening)
						VALUES (?, ?, ?, ?, ?, ?, ?)`

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
							(formulirID, pesertaID, pembayaranID, jurusanID, agamaID, genderPeserta, noAktaLahir, tempatLahir, tglLahir, NISN, Kelas, urutanAnak, jumlahSaudara, tglSubmit, statusID)
						VALUES(?, ?, ?, "", "", "", "", "", ?, "", "", ?, ?, ?, ?)`

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
		SET posterDaftar = ?, awalTahunAjar = ?, akhirTahunAjar = ?,  noRekening = ?, namaBank=?, pemilikRekening=?
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

	updatePembayaranFormulir  = "UpdatePembayaranFormulir"
	qUpdatePembayaranFormulir = `UPDATE T_PembayaranFormulir
								SET statusID=?, tglPembayaran=NOW() + INTERVAL 7 HOUR, buktiPembayaran=?
								WHERE pembayaranID=?`

	updateFormulir  = "UpdateFormulir"
	qUpdateFormulir = `UPDATE T_Formulir
						SET jurusanID=?, agamaID=?, genderPeserta=?, noAktaLahir=?, tempatLahir=?, tglLahir=?, NISN=?, Kelas=?, urutanAnak=?, jumlahSaudara=?, tglSubmit=NOW() + INTERVAL 7 HOUR, statusID=?
						WHERE formulirID=?`

	updateKontakPeserta  = "UpdateKontakPeserta"
	qUpdateKontakPeserta = `UPDATE T_KontakPeserta
							SET alamatTerakhir=?, kodePos=?, noTelpRumah=?
							WHERE kontakID=?`

	updateOrtu  = "UpdateOrtu"
	qUpdateOrtu = `UPDATE T_Ortu
					SET namaAyah=?, pekerjaanAyah=?, noTelpHpAyah=?, namaIbu=?, pekerjaanIbu=?, 
						noTelpHpIbu=?, namaWali=?, pekerjaanWali=?, noTelpHpWali=?
					WHERE ortuID=?;`

	updateBerkas  = "UpdateBerkas"
	qUpdateBerkas = `UPDATE T_Berkas
						SET statusID=?, aktaLahir=?, pasPhoto=?, rapor=?, tanggalUpload=NOW() + INTERVAL 7 HOUR
						WHERE berkasID=?`

	updateJadwalTest  = "UpdateJadwalTest"
	qUpdateJadwalTest = `UPDATE T_JadwalTest
							SET statusID=?, tglTest=?, waktuTest=CAST(? AS TIME)
							WHERE testID=?`
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
		{getAgama, qGetAgama},
		{getJurusan, qGetJurusan},

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
		{getPesertaDidikDetail, qGetPesertaDidikDetail},

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

		{getJadwalTestAll, qGetJadwaltestAll},
		{getJadwalTestPagination, qGetJadwalTestPagination},
		{getPembayaranFormulirAll, qGetPembayaranFormulirAll},
		{getPembayaranFormulirPagination, qGetPembayaranFormulirPagination},
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
		{updatePembayaranFormulir, qUpdatePembayaranFormulir},
		{updateFormulir, qUpdateFormulir},
		{updateKontakPeserta, qUpdateKontakPeserta},
		{updateOrtu, qUpdateOrtu},
		{updateBerkas, qUpdateBerkas},
		{updateJadwalTest, qUpdateJadwalTest},
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
