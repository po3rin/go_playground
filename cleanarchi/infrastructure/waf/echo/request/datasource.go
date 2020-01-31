package request

type DataSourceRequestParams struct {
	Name string `path:"name" v-get:"required"`
}
