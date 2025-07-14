package stringops

import "strings"

func trim(s string) string {
	str := strings.TrimSpace(s)
	return str
}
