package rdb

import (
	"github.com/po3rin/cleanarchi/usecase/port"
)

type SqlHandler interface {
	Execute(string, ...interface{}) (Result, port.Error)
	Query(string, ...interface{}) (Row, port.Error)
}

type Result interface {
	LastInsertId() (int64, port.Error)
	RowsAffected() (int64, port.Error)
}

type Row interface {
	Scan(...interface{}) port.Error
	Next() bool
	Close() port.Error
}
