package server

import (
	"github.com/po3rin/cleanarchi/domain/entity"
	"github.com/po3rin/cleanarchi/usecase/port"
)

/*
 * Input Port
 *  └─ Interactor で実装、Controller で使用される
 */
type DataSourceInputPort interface {
	DownloadDataSources() (*DownloadDataSourcesResponse, port.Error)
	DownloadDataSource(*DownloadDataSourceRequestParams) (*DownloadDataSourceResponse, port.Error)
}

type DownloadDataSourceRequestParams struct {
	DataSourceID entity.DataSourceID
}

/*
 * Output Port
 *  └─ Presenter で実装、Interactor で使用される
 */
type DataSourceOutputPort interface {
	DownloadDataSources([]entity.DataSource) (*DownloadDataSourcesResponse, port.Error)
	DownloadDataSource(*entity.DataSource) (*DownloadDataSourceResponse, port.Error)
}

type DownloadDataSourceResponse struct {
	DataSource *entity.DataSource
}

type DownloadDataSourcesResponse struct {
	DataSources []entity.DataSource
}
