package utils

import "log"

func ErrorLog(err error) {
	log.Fatalln("error here: ", err.Error())
}
