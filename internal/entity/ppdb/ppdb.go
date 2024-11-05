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
	StaffID        string     `db:"staffID" json:"staff_id"`
	StaffName      string     `db:"staffName" json:"staff_name"`
	StaffGender    string     `db:"staffGender" json:"staff_gender"`
	StaffPosition  string     `db:"staffPosition" json:"staff_position"`
	StaffTmptLahir *string    `db:"staffTmptLahir" json:"staff_tmpt_lahir"`
	StaffTglLahir  *time.Time `db:"staffTglLahir" json:"staff_tgl_lahir"`
	StaffPhoto     []byte     `db:"staffPhoto" json:"staff_photo"`
	LinkStaffPhoto string     `json:"link_staff_photo"`
}

type TableEvent struct {
	EventID        string     `db:"eventID" json:"event_id"`
	EventHeader    string     `db:"eventHeader" json:"event_header"`
	EventStartDate time.Time  `db:"eventStartDate" json:"event_start_date"`
	EventEndDate   *time.Time `db:"eventEndDate" json:"event_end_date"`
	EventDesc      string     `db:"eventDesc" json:"event_desc"`
	EventImage     []byte     `db:"eventImage" json:"event_image"`
	LinkEventImage string     `json:"link_event_image"`
}

type TableStatus struct {
	StatusID   string `db:"statusID" json:"status_id"`
	StatusName string `db:"statusName" json:"status_name"`
	StatusDesc string `db:"statusDesc" json:"status_desc"`
}

type TablePesertaDidik struct {
	PesertaID         string `db:"pesertaID" json:"peserta_id,omitempty"`
	PesertaName       string `db:"pesertaName" json:"peserta_name,omitempty"`
	Password          string `db:"password" json:"password,omitempty"`
	EmailPeserta      string `db:"emailPeserta" json:"email_peserta,omitempty"`
	NoTelpHpPeserta   string `db:"noTelpHpPeserta" json:"no_telp_hp_peserta,omitempty"`
	SekolahAsalYN     string `db:"sekolahAsalYN" json:"sekolah_asal_yn,omitempty"`
	SekolahAsal       string `db:"sekolahAsal" json:"sekolah_asal,omitempty"`
	AlamatSekolahAsal string `db:"alamatSekolahAsal" json:"alamat_sekolah_asal,omitempty"`
}

type TablePembayaranFormulir struct {
	PembayaranID    string    `db:"pembayaranID" json:"pembayaran_id"`
	PesertaID       string    `db:"pesertaID" json:"peserta_id"`
	StatusID        string    `db:"statusID" json:"status_id"`
	StatusName      string    `db:"statusName" json:"status_name"`
	TglPembayaran   time.Time `db:"tglPembayaran" json:"tgl_pembayaran"`
	HargaFormulir   float64   `db:"hargaFormulir" json:"harga_formulir"`
	BuktiPembayaran []byte    `db:"buktiPembayaran" json:"bukti_pembayaran"`
}

type TableFormulir struct {
	FormulirID    string    `db:"formulirID" json:"formulir_id"`
	PesertaID     string    `db:"pesertaID" json:"peserta_id"`
	PembayaranID  string    `db:"pembayaranID" json:"pembayaran_id"`
	JurusanID     string    `db:"jurusanID" json:"jurusan_id"`
	AgamaID       string    `db:"agamaID" json:"agama_id"`
	GenderPeserta string    `db:"genderPeserta" json:"gender_peserta"`
	TempatLahir   string    `db:"tempatLahir" json:"tempat_lahir"`
	TglLahir      time.Time `db:"tglLahir" json:"tgl_lahir"`
	NISN          string    `db:"NISN" json:"nisn"`
	Kelas         string    `db:"Kelas" json:"kelas"`
	TglSubmit     time.Time `db:"tglSubmit" json:"tgl_submit"`
	StatusID      string    `db:"statusID" json:"status_id"`
}

type TableKontakPeserta struct {
	KontakID       string `db:"kontakID" json:"kontak_id"`
	FormulirID     string `db:"formulirID" json:"formulir_id"`
	AlamatTerakhir string `db:"alamatTerakhir" json:"alamat_terakhir"`
	KodePos        string `db:"kodePos" json:"kode_pos"`
	NoTelpRumah    string `db:"noTelpRumah" json:"no_telp_rumah"`
}

type TableOrtu struct {
	OrtuID        string `db:"ortuID" json:"ortu_id"`
	FormulirID    string `db:"formulirID" json:"formulir_id"`
	NamaAyah      string `db:"namaAyah" json:"nama_ayah"`
	PekerjaanAyah string `db:"pekerjaanAyah" json:"pekerjaan_ayah"`
	NoTelpHpAyah  string `db:"noTelpHpAyah" json:"no_telp_hp_ayah"`
	NamaIbu       string `db:"namaIbu" json:"nama_ibu"`
	PekerjaanIbu  string `db:"pekerjaanIbu" json:"pekerjaan_ibu"`
	NoTelpHpIbu   string `db:"noTelpHpIbu" json:"no_telp_hp_ibu"`
	NamaWali      string `db:"namaWali" json:"nama_wali"`
	PekerjaanWali string `db:"pekerjaanWali" json:"pekerjaan_wali"`
	NoTelpHpWali  string `db:"noTelpHpWali" json:"no_telp_hp_wali"`
}

type TableDataFormulir struct {
	FormulirID     string    `db:"formulirID" json:"formulir_id"`
	PesertaID      string    `db:"pesertaID" json:"peserta_id"`
	PembayaranID   string    `db:"pembayaranID" json:"pembayaran_id"`
	JurusanID      string    `db:"jurusanID" json:"jurusan_id"`
	AgamaID        string    `db:"agamaID" json:"agama_id"`
	GenderPeserta  string    `db:"genderPeserta" json:"gender_peserta"`
	TempatLahir    string    `db:"tempatLahir" json:"tempat_lahir"`
	TglLahir       time.Time `db:"tglLahir" json:"tgl_lahir"`
	NISN           string    `db:"NISN" json:"nisn"`
	Kelas          string    `db:"Kelas" json:"kelas"`
	TglSubmit      time.Time `db:"tglSubmit" json:"tgl_submit"`
	StatusID       string    `db:"statusID" json:"status_id"`
	StatusName     string    `db:"statusName" json:"status_name"`
	KontakID       string    `db:"kontakID" json:"kontak_id"`
	AlamatTerakhir string    `db:"alamatTerakhir" json:"alamat_terakhir"`
	KodePos        string    `db:"kodePos" json:"kode_pos"`
	NoTelpRumah    string    `db:"noTelpRumah" json:"no_telp_rumah"`
	OrtuID         string    `db:"ortuID" json:"ortu_id"`
	NamaAyah       string    `db:"namaAyah" json:"nama_ayah"`
	PekerjaanAyah  string    `db:"pekerjaanAyah" json:"pekerjaan_ayah"`
	NoTelpHpAyah   string    `db:"noTelpHpAyah" json:"no_telp_hp_ayah"`
	NamaIbu        string    `db:"namaIbu" json:"nama_ibu"`
	PekerjaanIbu   string    `db:"pekerjaanIbu" json:"pekerjaan_ibu"`
	NoTelpHpIbu    string    `db:"noTelpHpIbu" json:"no_telp_hp_ibu"`
	NamaWali       string    `db:"namaWali" json:"nama_wali"`
	PekerjaanWali  string    `db:"pekerjaanWali" json:"pekerjaan_wali"`
	NoTelpHpWali   string    `db:"noTelpHpWali" json:"no_telp_hp_wali"`
}

type TableBerkas struct {
	BerkasID      string    `db:"berkasID" json:"berkas_id"`
	PesertaID     string    `db:"pesertaID" json:"peserta_id"`
	StatusID      string    `db:"statusID" json:"status_id"`
	AktalLahir    []byte    `db:"aktaLahir" json:"akta_lahir"`
	PasPhoto      []byte    `db:"pasPhoto" json:"pas_photo"`
	Rapor         []byte    `db:"rapor" json:"rapor"`
	TanggalUpload time.Time `db:"tanggalUpload" json:"tanggal_upload"`
}

type TableJadwalTest struct {
	TestID    string    `db:"testID" json:"test_id"`
	PesertaID string    `db:"pesertaID" json:"peserta_id"`
	StatusID  string    `db:"statusID" json:"status_id"`
	TglTest   time.Time `db:"tglTest" json:"tgl_test"`
	WaktuTest time.Time `db:"waktuTest" json:"waktu_test"`
}
