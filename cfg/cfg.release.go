// +build release

package cfg

var CONFIG = &ConfigType{
	IsDev:               false,
	DatabaseResourceStr: "host=db user=postgres password=admin dbname=postgres sslmode=disable",
}
