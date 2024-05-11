package str

import "strings"

func IsStringEmpty(str string) bool {
	return len(strings.TrimSpace(str)) == 0
}
