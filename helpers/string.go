package helpers

import (
	"strconv"
)

func Unidecode(s string) string {
	r, _ := strconv.Unquote(s)
	return r
}
