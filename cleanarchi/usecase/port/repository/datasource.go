package repository

import (
	"github.com/po3rin/cleanarchi/domain/entity"
	"github.com/po3rin/cleanarchi/usecase/port"
)

type DataSourceRepository interface {
	FindAll() ([]entity.DataSource, port.Error)
	Find(entity.DataSourceID) (*entity.DataSource, port.Error)
}
