package utils

import (
	"crypto/sha256"
	"encoding/hex"
	"io"
)

func GenPassword(pwd string) string {
	return sha256Encode(pwd)
}

func sha256Encode(pwd string) string {
	h := sha256.New()
	io.WriteString(h, pwd)
	return hex.EncodeToString(h.Sum(nil))
}

// TODO: a scrypt encoded password
func scryptEncode(pwd string) string {
	return ""
	// realPasword, err := scrypt.Key([]byte(pwd), []byte(salt), 16384, 8, 1, 32)
	// if err != nil {
	// 	revel.INFO.Println("error in crypt password", err)
	// 	return pwd
	// }
	// return string(realPasword)

}
