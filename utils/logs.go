package utils

import "log"

func ErrorLog(err error) {
	if err != nil {
		log.Fatalln("error here: ", err.Error())
	}
}
