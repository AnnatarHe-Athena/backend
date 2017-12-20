package model

import (
	"database/sql"
)

// DBInstance is database connection resource to global share
var DBInstance *sql.DB

// DatabaseInit have to called at init.
func DatabaseInit() {

}
