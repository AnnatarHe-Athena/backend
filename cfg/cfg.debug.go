// +build !release

package cfg

var CONFIG = &ConfigType{
	IsDev:               true,
	DatabaseResourceStr: "host=127.0.0.1 user=postgres password=admin dbname=postgres sslmode=disable",
	Port:                ":9988",
}
