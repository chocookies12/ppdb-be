package ppdb

// import "gopkg.in/guregu/null.v3/zero"

// "time"

// "gopkg.in/guregu/null.v3/zero"

type TableAdmin struct {
	AdminID       string `db:"admin_id" json:"admin_id"`
	AdminName     string `db:"admin_name" json:"admin_name"`
	AdminGmail    string `db:"admin_email" json:"admin_email"`
	AdminPassword string `db:"admin_password" json: "admin_password"`
}
