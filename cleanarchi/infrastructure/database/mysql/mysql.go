package mysql

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"

	"github.com/po3rin/cleanarchi/config"
	"github.com/po3rin/cleanarchi/interface/errs"
	"github.com/po3rin/cleanarchi/interface/gateway/database/rdb"
	"github.com/po3rin/cleanarchi/usecase/port"
)

type SqlHandler struct {
	Conn *sql.DB
}

// DIP に基づき、rdb.SqlHandler の Interface を実装していく
func NewSqlHandler(config config.DB) (rdb.SqlHandler, port.Error) {
	var (
		host     = config.Host
		port     = config.Port
		dbName   = config.Database
		user     = config.User
		password = config.Password
	)
	driverName := "mysql"
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", user, password, host, port, dbName)

	conn, err := sql.Open(driverName, dataSourceName)
	if err != nil {
		return nil, &errs.ErrorUnknown{}
	}
	err = conn.Ping()
	if err != nil {
		return nil, &errs.ErrorUnknown{}
	}
	sqlHandler := new(SqlHandler)
	sqlHandler.Conn = conn
	return sqlHandler, nil
}

func (handler *SqlHandler) Execute(statement string, args ...interface{}) (rdb.Result, port.Error) {
	res := SqlResult{}
	result, err := handler.Conn.Exec(statement, args...)
	if err != nil {
		return res, &errs.ErrorUnknown{}
	}
	res.Result = result
	return res, nil
}

func (handler *SqlHandler) Query(statement string, args ...interface{}) (rdb.Row, port.Error) {
	rows, err := handler.Conn.Query(statement, args...)
	if err != nil {
		return new(SqlRow), &errs.ErrorUnknown{}
	}
	row := new(SqlRow)
	row.Rows = rows
	return row, nil
}

type SqlResult struct {
	Result sql.Result
}

func (r SqlResult) LastInsertId() (int64, port.Error) {
	res, err := r.Result.LastInsertId()
	if err != nil {
		return res, &errs.ErrorUnknown{}
	}
	return res, nil
}

func (r SqlResult) RowsAffected() (int64, port.Error) {
	res, err := r.Result.RowsAffected()
	if err != nil {
		return res, &errs.ErrorUnknown{}
	}
	return res, nil
}

type SqlRow struct {
	Rows *sql.Rows
}

func (r SqlRow) Scan(dest ...interface{}) port.Error {
	if err := r.Rows.Scan(dest...); err != nil {
		return &errs.ErrorUnknown{}
	}
	return nil
}

func (r SqlRow) Next() bool {
	return r.Rows.Next()
}

func (r SqlRow) Close() port.Error {
	if err := r.Rows.Close(); err != nil {
		return &errs.ErrorUnknown{}
	}
	return nil
}
