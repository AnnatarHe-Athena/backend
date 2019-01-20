package model

import (
	"log"

	"github.com/jmoiron/sqlx"

	"github.com/douban-girls/backend/cfg"
	_ "github.com/lib/pq"
)

// DBInstance is database connection resource to global share
var DBInstance *sqlx.DB

// DatabaseInit have to called at init.
func DatabaseInit() (err error) {
	DBInstance, err = sqlx.Open("postgres", cfg.CONFIG.DatabaseResourceStr)
	if err != nil {
		log.Panic("connect to database server error: ", err)
	} else {
		log.Println("connect to the database successfully with: ", cfg.CONFIG.DatabaseResourceStr)
	}

	// DBInstance.DB.SetMaxOpenConns(50)
	return err
}
