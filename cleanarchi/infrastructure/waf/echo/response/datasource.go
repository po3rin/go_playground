package response

import (
	"net/http"

	"github.com/po3rin/cleanarchi/domain/entity"
	"github.com/po3rin/cleanarchi/usecase/port/server"
)

type DataSource struct {
	Name entity.DataSourceID `json:"name"`
	Type string              `json:"type"`
}

type DataSourceResult struct {
	DataSource *DataSource `json:"datasource"`
}

type DataSourcesResult struct {
	DataSources []DataSource `json:"datasources"`
}

type Error struct {
	StatusCode int
	Message    string
}

func DataSourceResponseAdapter(_res *server.DownloadDataSourceResponse) (DataSourceResult, int) {
	var res DataSourceResult

	if _res.DataSource == nil {
		res = DataSourceResult{
			DataSource: nil,
		}
	} else {
		res = DataSourceResult{
			DataSource: &DataSource{
				Name: _res.DataSource.ID,
				Type: _res.DataSource.Type,
			},
		}
	}
	return res, http.StatusOK
}

// この Adapter がないと JSON を返すという前提が Use Cases の内側に入ってきてしまうので、ここであえてマッピングしてます
func DataSourcesResponseAdapter(_res *server.DownloadDataSourcesResponse) (DataSourcesResult, int) {
	res := DataSourcesResult{}
	res.DataSources = make([]DataSource, 0)

	for i := 0; i < len(_res.DataSources); i++ {
		datasource := &_res.DataSources[i]
		res.DataSources = append(res.DataSources, DataSource{
			Name: datasource.ID,
			Type: datasource.Type,
		})
	}
	return res, http.StatusOK
}
