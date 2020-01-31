package presenter

import (
	"github.com/po3rin/cleanarchi/config"
	"github.com/po3rin/cleanarchi/usecase/port/server"
)

type HTTPPresenter struct {
	config *config.Config
	server.DataSourceOutputPort
}

func NewHTTPPresenter(config *config.Config) *HTTPPresenter {
	return &HTTPPresenter{
		config: config,
	}
}
