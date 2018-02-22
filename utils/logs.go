package utils

import (
	"fmt"
	"log"
)

func ErrorLog(err error) {
	if err != nil {
		log.Println("error here: ", err.Error())
	}
}

func RedisKeyGen(ip string) string {
	return "ip:" + ip + ":requested"
}

var colorMap = map[string]string{
	"black":   "30",
	"red":     "31",
	"green":   "32",
	"yellow":  "33",
	"blue":    "34",
	"default": "39",
}

// LogColor will print variables with color
func LogColor(color string, v ...interface{}) {
	fmt.Print("\033[" + colorMap[color] + "m")

	log.Println(v...)

	fmt.Print("\033[" + colorMap["default"] + "m")
}
