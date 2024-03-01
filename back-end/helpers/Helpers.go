package helpers

import (
	"encoding/base64"
	"os"
)

func EncodeUser() string {
	user := os.Getenv("ZINC_FIRST_ADMIN_USER")
	pass := os.Getenv("ZINC_FIRST_ADMIN_PASSWORD")
	code := base64.StdEncoding.EncodeToString([]byte(user + ":" + pass))
	return code
}
