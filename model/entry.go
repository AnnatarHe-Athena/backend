package model

import (
	"database/sql"
	"log"

	"github.com/douban-girls/backend/cfg"
	_ "github.com/lib/pq"
)

// DBInstance is database connection resource to global share
var DBInstance *sql.DB

// DatabaseInit have to called at init.
func DatabaseInit() (err error) {
	DBInstance, err = sql.Open("postgres", cfg.CONFIG.DatabaseResourceStr)
	if err != nil {
		log.Panic("connect to database server error: ", err)
	} else {
		log.Println("connect to the database successfully with: ", cfg.CONFIG.DatabaseResourceStr)
	}
	return err
}
