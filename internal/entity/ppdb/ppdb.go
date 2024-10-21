package ppdb

import "time"

// import "gopkg.in/guregu/null.v3/zero"

// "time"

// "gopkg.in/guregu/null.v3/zero"

type TableKontakSekolah struct {
	Kontak_id         string `db:"kontakKYID" json:"kontak_id"`
	Alamat_sekolah    string `db:"alamatSekolah" json:"alamat_sekolah"`
	No_telp1          string `db:"noTelpSekolah1" json:"no_telp1"`
	No_telp2          string `db:"noTelpSekolah2" json:"no_telp2"`
	Email_sekolah     string `db:"emailSekolah" json:"email_sekolah"`
	Instagram_sekolah string `db:"instagramSekolah" json:"instagram_sekolah"`
}

type TableAdmin struct {
	AdminID    string `db:"adminID" json:"admin_id"`
	RoleID     string `db:"roleID" json:"role_id"`
	AdminName  string `db:"adminName" json:"admin_name"`
	EmailAdmin string `db:"emailAdmin" json:"email_admin"`
	Password   string `db:"password" json:"password"`
}

type TableKelolaDataAdmin struct {
	AdminID    string `db:"adminID" json:"admin_id"`
	RoleID     string `db:"roleID" json:"role_id"`
	AdminName  string `db:"adminName" json:"admin_name"`
	EmailAdmin string `db:"emailAdmin" json:"email_admin"`
	Password   string `db:"password" json:"password"`
	RoleName   string `db:"roleName" json:"role_name"`
	RoleDesc   string `db:"roleDesc" json:"role_desc"`
}

type TableRole struct {
	RoleID   string `db:"roleID" json:"role_id"`
	RoleName string `db:"roleName" json:"role_name"`
	RoleDesc string `db:"roleDesc" json:"role_desc"`
}

type TableInfoDaftar struct {
	InfoID           string `db:"infoID" json:"info_id"`
	PosterDaftar     []byte `db:"posterDaftar" json:"poster_daftar"`
	LinkPosterDaftar string `json:"link_poster_daftar"`
	AwalTahunAjar    string `db:"awalTahunAjar" json:"awal_tahun_ajar"`
	AkhirTahunAjar   string `db:"akhirTahunAjar" json:"akhir_tahun_ajar"`
}

type TableBanner struct {
	BannerID        string `db:"bannerID" json:"banner_id"`
	BannerName      string `db:"bannerName" json:"banner_name"`
	BannerImage     []byte `db:"bannerImage" json:"banner_image"`
	LinkBannerImage string `json:"link_banner_image"`
}

type TableFasilitas struct {
	FasilitasID        string `db:"fasilitasID" json:"fasilitas_id"`
	FasilitasName      string `db:"fasilitasName" json:"fasilitas_name"`
	FasilitasImage     []byte `db:"fasilitasImage" json:"fasilitas_image"`
	LinkFasilitasImage string `json:"link_fasilitas_image"`
}

type TableStaff struct {
	StaffID        string    `db:"staffID" json:"staff_id"`
	StaffName      string    `db:"staffName" json:"staff_name"`
	StaffGender    string    `db:"staffGender" json:"staff_gender"`
	StaffPosition  string    `db:"staffPosition" json:"staff_position"`
	StaffTmptLahir string    `db:"staffTmptLahir" json:"staff_tmpt_lahir"`
	StaffTglLahir  time.Time `db:"staffTglLahir" json:"staff_tgl_lahir"`
	StaffPhoto     []byte    `db:"staffPhoto" json:"staff_photo"`
	LinkStaffPhoto string    `json:"link_staff_photo"`
}
