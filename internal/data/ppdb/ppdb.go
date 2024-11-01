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
	}
	insertStmt = []statement{
		{insertDataAdmin, qInsertDataAdmin},
		{insertInfoDaftar, qInsertInfoDaftar},
		{insertBanner, qInsertBanner},
		{insertFasilitas, qInsertFasilitas},
		{insertProfileStaff, qInsertProfileStaff},
		{insertEvent, qInsertEvent},
	}
	updateStmt = []statement{
		{updateBanner, qUpdateBanner},
		{updateInfoDaftar, qUpdateInfoDaftar},
		{updateFasilitas, qUpdateFasilitas},
		{updateProfileStaff, qUpdateProfileStaff},
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
