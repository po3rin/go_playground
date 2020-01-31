package echo

import (
	"fmt"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/labstack/gommon/log"

	conf "github.com/po3rin/cleanarchi/config"
	"github.com/po3rin/cleanarchi/infrastructure/database/mysql"
	"github.com/po3rin/cleanarchi/interface/controller"
)

type Server struct {
	*echo.Echo
	*conf.Config
}

func createServer(config *conf.Config) (*Server, error) {
	return &Server{
		Echo:   echo.New(),
		Config: config,
	}, nil
}

func (s *Server) setRouter() {
	s.Echo.Use(
		middleware.Logger(),
		middleware.Recover(),
		func(h echo.HandlerFunc) echo.HandlerFunc {
			return func(c echo.Context) error {
				return h(&Context{c})
			}
		},
	)
	s.Echo.Binder = &CustomBinder{}
	s.Echo.HTTPErrorHandler = CustomHTTPErrorHandler

	handler, err := mysql.NewSqlHandler(s.Config.DB)
	if err != nil {
		s.Echo.Logger.Fatal(err)
	}

	api := s.Echo.Group(s.getEndpointRoot())
	{
		datasourceController := controller.NewDataSourceController(s.Config)
		// datasources
		api.GET("/datasources", s.GetDataSources(datasourceController))
		api.GET("/datasources/:name", s.GetDataSource(datasourceController))
	}
}

func (s *Server) getPrefix() string {
	return fmt.Sprintf("%s", s.Config.Meta.Version)
}

func (s *Server) getEndpointRoot() string {
	return fmt.Sprintf("/api/%s", s.getPrefix())
}

func (s *Server) getAddr() string {
	return fmt.Sprintf("%s:%d", s.Config.Server.Host, s.Config.Server.Port)
}

func (s *Server) run() {
	s.Echo.Logger.Fatal(s.Echo.Start(s.getAddr()))
}

func Run() {
	config, err := conf.LoadConfig()
	if err != nil {
		panic(err)
	}

	s, err := createServer(&config)
	if err != nil {
		panic(err)
	}
	if config.Server.Debug {
		s.Echo.Logger.SetLevel(log.DEBUG)
	}
	s.setRouter()

	s.run()
}
