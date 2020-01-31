package interactor

import (
	"github.com/po3rin/cleanarchi/usecase/port/repository"
	"github.com/po3rin/cleanarchi/usecase/port/server"

	"github.com/po3rin/cleanarchi/config"
	"github.com/po3rin/cleanarchi/usecase/port"
)

type DataSourceInteractor struct {
	Config               *config.Config
	OutputPort           server.DataSourceOutputPort
	DataSourceRepository repository.DataSourceRepository
}

func NewDataSourceInteractor(
	config *config.Config,
	outputPort server.DataSourceOutputPort,
	dataSourceRepository repository.DataSourceRepository,
) *DataSourceInteractor {
	return &DataSourceInteractor{
		// Container:            newUsecaseContainer(config), // TODO
		OutputPort:           outputPort,
		DataSourceRepository: dataSourceRepository,
	}
}

// Input Port の実装
func (i *DataSourceInteractor) DownloadDataSources() (*server.DownloadDataSourcesResponse, port.Error) {
	res, err := i.DataSourceRepository.FindAll()
	if err != nil {
		return nil, err
	}
	// Output Port の使用
	return i.OutputPort.DownloadDataSources(res)
}

func (i *DataSourceInteractor) DownloadDataSource(params *server.DownloadDataSourceRequestParams) (*server.DownloadDataSourceResponse, port.Error) {
	res, err := i.DataSourceRepository.Find(params.DataSourceID)
	if err != nil {
		return nil, err
	}
	return i.OutputPort.DownloadDataSource(res)
}
