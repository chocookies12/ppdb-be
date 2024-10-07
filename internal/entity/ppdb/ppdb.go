package ppdb

// import "gopkg.in/guregu/null.v3/zero"

// "time"

// "gopkg.in/guregu/null.v3/zero"

// type TableAdmin struct {
// 	AdminID       string `db:"admin_id" json:"admin_id"`
// 	AdminName     string `db:"admin_name" json:"admin_name"`
// 	AdminGmail    string `db:"admin_email" json:"admin_email"`
// 	AdminPassword string `db:"admin_password" json: "admin_password"`
// }

type TableKontakSekolah struct {
	Kontak_id         string `db:"kontakKYID" json:"kontak_id"`
	Alamat_sekolah    string `db:"alamatSekolah" json:"alamat_sekolah"`
	No_telp1          string `db:"noTelpSekolah1" json:"no_telp1"`
	No_telp2          string `db:"noTelpSekolah2" json:"no_telp2"`
	Email_sekolah     string `db:"emailSekolah" json:"email_sekolah"`
	Instagram_sekolah string `db:"instagramSekolah" json:"instagram_sekolah"`
}
