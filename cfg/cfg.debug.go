// +build !release

package cfg

var CONFIG = &ConfigType{
	IsDev:               true,
	DatabaseResourceStr: "host=db user=postgres password=admin dbname=postgres sslmode=disable",
	Port:                ":9988",
	RavenDSN:            "hello",
}
