package delpadd

import "strings"

func RemovePadding(encoded string) string {
	return strings.TrimRight(encoded, "=")
}
