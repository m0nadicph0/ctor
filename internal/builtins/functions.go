package builtins

import (
	md5s "crypto/md5"
	"encoding/hex"
	uuid2 "github.com/google/uuid"
	"strings"
	"text/template"
)

var BuiltinFunctions = template.FuncMap{
	"upper":   strings.ToUpper,
	"lower":   strings.ToLower,
	"uuid":    uuid,
	"squeeze": squeeze,
	"md5":     md5,
}

func uuid() string {
	return uuid2.New().String()
}

func squeeze(s string) string {
	return strings.Join(strings.Split(s, " "), "")
}

func md5(s string) string {
	hash := md5s.New()
	_, err := hash.Write([]byte(s))
	if err != nil {
		return ""
	}
	checksum := hash.Sum(nil)
	return hex.EncodeToString(checksum)
}
