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
					   LIMIT 1;`

	getRole  = "GetRole"
	qGetRole = `Select roleID, roleName, roleDesc FROM T_Role`

	//query insert
	insertDataAdmin  = "InsertDataAdmin"
	qInsertDataAdmin = `INSERT INTO T_Admin (adminID, roleID, adminName, password, emailAdmin)
						VALUES (?, ?, ?, ?, ?)`

	//query delete
	deleteDataAdmin  = "DeleteDataAdmin"
	qDeleteDataAdmin = `DELETE FROM T_Admin WHERE adminID = ?`
)

var (
	readStmt = []statement{
		{loginAdmin, qLoginAdmin},
		{getKontakSekolah, qGetKontakSekolah},
		{getDataAdmin, qGetDataAdmin},
		{getDataAdminPagination, qGetDataAdminPagination},
		{getLastAdminId, qGetLastAdminId},

		{getRole, qGetRole},
	}
	insertStmt = []statement{
		{insertDataAdmin, qInsertDataAdmin},
	}
	updateStmt = []statement{}
	deleteStmt = []statement{
		{deleteDataAdmin, qDeleteDataAdmin},
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
