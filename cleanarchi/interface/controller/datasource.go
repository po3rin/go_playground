package controller

import (
	"github.com/po3rin/cleanarchi/config"
	"github.com/po3rin/cleanarchi/usecase/port"

	"github.com/po3rin/cleanarchi/interface/gateway/database/rdb"
	"github.com/po3rin/cleanarchi/interface/presenter"
	"github.com/po3rin/cleanarchi/usecase/interactor"
	"github.com/po3rin/cleanarchi/usecase/port/server"
)

type DataSourceController struct {
	InputPort server.DataSourceInputPort
}

func NewDataSourceController(config *config.Config) *DataSourceController {
	return &DataSourceController{
		InputPort: interactor.NewDataSourceInteractor(
			config,
			presenter.NewHTTPPresenter(config),
			rdb.NewDataSourceRDBRepository(*config),
		),
	}
}

func (c *DataSourceController) DownloadDataSources() (*server.DownloadDataSourcesResponse, port.Error) {
	// Input Port の使用
	return c.InputPort.DownloadDataSources()
}

func (c *DataSourceController) DownloadDataSource(params *server.DownloadDataSourceRequestParams) (*server.DownloadDataSourceResponse, port.Error) {
	return c.InputPort.DownloadDataSource(params)
}
