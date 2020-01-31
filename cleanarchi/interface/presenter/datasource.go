package presenter

import (
	"github.com/po3rin/cleanarchi/domain/entity"
	"github.com/po3rin/cleanarchi/usecase/port"
	"github.com/po3rin/cleanarchi/usecase/port/server"
)

// Output Port の実装
func (p *HTTPPresenter) DownloadDataSources(datasources []entity.DataSource) (*server.DownloadDataSourcesResponse, port.Error) {
	res := &server.DownloadDataSourcesResponse{}
	res.DataSources = datasources
	return res, nil
}

func (p *HTTPPresenter) DownloadDataSource(datasource *entity.DataSource) (*server.DownloadDataSourceResponse, port.Error) {
	res := &server.DownloadDataSourceResponse{}
	res.DataSource = datasource
	return res, nil
}
