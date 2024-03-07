package crypto

import (
	"crypto/sha256"
	"encoding/base64"
	"fmt"
)

func EncodePassword(uname string, pswd string) string {
	if uname == "" {
		return ""
	}

	sPassword := fmt.Sprintf("%s-%s", uname, pswd)
	bPassword := sha256.Sum256([]byte(sPassword))
	return base64.StdEncoding.EncodeToString(bPassword[:])
}
