package echo

import (
	"net/http"

	"github.com/labstack/echo"

	"github.com/po3rin/cleanarchi/domain/entity"
	"github.com/po3rin/cleanarchi/infrastructure/waf/echo/request"
	"github.com/po3rin/cleanarchi/infrastructure/waf/echo/response"
	"github.com/po3rin/cleanarchi/interface/controller"
	"github.com/po3rin/cleanarchi/usecase/port"
)

func (s *Server) GetDataSources(controller *controller.DataSourceController) echo.HandlerFunc {
	return C(func(c *Context) error {
		_res, err := controller.DownloadDataSources()
		if err != nil {
			c.Logger().Error(err)
			return echo.NewHTTPError(http.StatusBadRequest)
		}
		res, status := response.DataSourcesResponseAdapter(_res)

		return c.JSON(status, res)
	})
}

func (s *Server) GetDataSource(controller *controller.DataSourceController) echo.HandlerFunc {
	return C(func(c *Context) error {
		req := new(request.DataSourceRequestParams)
		if err := c.BindValidate(req); err != nil {
			c.Logger().Error(err)
			return echo.NewHTTPError(http.StatusBadRequest)
		}

		// クライアントからのリクエストを内部に流して、
		// 内部からのレスポンスをクライアントに流す。
		_req := port.DownloadDataSourceRequestParams{
			DataSourceID: entity.DataSourceID(req.Name),
		}
		_res, err := controller.DownloadDataSource(&_req)
		if err != nil {
			c.Logger().Error(err)
			return echo.NewHTTPError(http.StatusBadRequest)
		}
		res, status := response.DataSourceResponseAdapter(_res)

		return c.JSON(status, res)
	})
}
