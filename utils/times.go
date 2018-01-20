package utils

import (
	"time"
)

// Timestamp gen unix time stamp from string to int64
func Timestamp(createdBy string) (createdAtUnix int64) {
	timestamp, err := time.Parse(time.RFC3339, createdBy)
	if err != nil {
		createdAtUnix = time.Now().Unix()
		ErrorLog(err)
	} else {
		createdAtUnix = timestamp.Unix()
	}
	return
}
