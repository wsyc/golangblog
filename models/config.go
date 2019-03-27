package models

type Config struct {

	App map[string]string

	Session map[string]string

	Datasource map[string](map[string]string)

	Static map[string]string

	StaticFile map[string]string

	Logger map[string]string

	Templete map[string]string

	All map[string]string

}