package rdb

import (
	"github.com/po3rin/cleanarchi/config"
	"github.com/po3rin/cleanarchi/domain/entity"
	"github.com/po3rin/cleanarchi/interface/errs"
	"github.com/po3rin/cleanarchi/usecase/port"
)

type DataSourceRDBRepositoryAdapter struct {
	config config.Config
	SqlHandler
}

func NewDataSourceRDBRepository(
	config config.Config,
) *DataSourceRDBRepositoryAdapter {
	return &DataSourceRDBRepositoryAdapter{
		config: config,
	}
}

func (repo *DataSourceRDBRepositoryAdapter) FindAll() ([]entity.DataSource, port.Error) {
	var datasources []entity.DataSource
	row, err := repo.Query("SELECT id, type FROM datasource_m")
	defer row.Close()
	if err != nil {
		return nil, &errs.ErrorUnknown{}
	}
	var (
		id       string
		dataType string
	)
	for row.Next() {
		if err = row.Scan(&id, &dataType); err != nil {
			return nil, &errs.ErrorUnknown{}
		}
		datasources = append(datasources, entity.DataSource{
			ID:   entity.DataSourceID(id),
			Type: dataType,
		})
	}
	return datasources, nil
}

func (repo *DataSourceRDBRepositoryAdapter) Find(identifier entity.DataSourceID) (*entity.DataSource, port.Error) {
	// 例えば、同じ RDB でも MySQL と PostgreSQL でクエリを変更するとかであれば、ファイルを分ける必要が出てくるかもしれない
	row, err := repo.Query("SELECT id, type FROM datasource_m WHERE id = ?", identifier)
	defer row.Close()
	if err != nil {
		return nil, &errs.ErrorUnknown{}
	}
	if !row.Next() {
		return nil, &errs.ErrorResourceNotFound{}
	}
	var (
		id       string
		dataType string
	)
	if err = row.Scan(&id, &dataType); err != nil {
		return nil, &errs.ErrorUnknown{}
	}
	datasource := entity.DataSource{
		ID:   entity.DataSourceID(id),
		Type: dataType,
	}
	return &datasource, nil
}
