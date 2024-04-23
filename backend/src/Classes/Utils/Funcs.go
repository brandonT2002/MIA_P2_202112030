package utils

import "strings"

func ClearString(name string) string {
	return strings.TrimSpace(strings.ReplaceAll(name, "\x00", ""))
}

func CompareStrings(s1, s2 string) bool {
	return ClearString(s1) == ClearString(s2)
}
