package funcs

import (
	"strings"
)

func Pascal(v string) string {
	return strings.ToUpper(v[:1]) + v[1:]
}
