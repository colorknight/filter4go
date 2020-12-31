package util

import (
	"filter4go/metadata"
	"strings"
)

func TranslateEscape(escStr *string) string {
	if escStr == nil || len(*escStr) == 0 {
		return ""
	}
	var bd strings.Builder
	bd.WriteRune(metadata.QUOTE)
	for _, c := range *escStr {
		if c == metadata.QUOTE {
			bd.WriteRune(metadata.QUOTE)
		}
		bd.WriteRune(c)
	}
	return bd.String()
}

func UntranslateEscape(str *string) string {
	if str == nil || len(*str) == 0 {
		return ""
	}
	b := []rune(*str)
	l := len(*str)
	if b[0] != metadata.QUOTE || b[l-1] != metadata.QUOTE {
		return *str
	}
	var bd strings.Builder
	for i := 1; i < l-1; i++ {
		if b[i] == metadata.QUOTE && b[i+1] == metadata.QUOTE {
			i++
			if i == l-1 {
				break
			}
			bd.WriteRune(b[i])
			continue
		}
		bd.WriteRune(b[i])
	}
	return bd.String()
}
