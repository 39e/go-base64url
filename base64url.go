package base64url

import (
	"encoding/base64"
	"fmt"
	"strings"
)

func Encode(rawStr string) string {
	encodeStr := base64.StdEncoding.EncodeToString([]byte(rawStr))
	r := strings.NewReplacer("=", "-", "/", "_", "+", ".")
	return r.Replace(encodeStr)
}

func Decode(encodeStr string) string {
	r := strings.NewReplacer("-", "=", "_", "/", ".", "+")
	data, err := base64.StdEncoding.DecodeString(r.Replace(encodeStr))
	if err != nil {
		fmt.Println("error: ", err)
		return ""
	}
	return string(data)
}
