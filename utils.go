package blackfriday

import "strings"

func ReplaceChars(str, chr string, rep rune) string {
	return strings.Map(func(r rune) rune {
		if strings.IndexRune(chr, r) < 0 {
			return r
		} else {
			return rep
		}
	}, str)
}

func StripChars(str, chr string) string {
	return strings.Map(func(r rune) rune {
		if strings.IndexRune(chr, r) < 0 {
			return r
		}
		return -1
	}, str)
}

func StripLeadingChars(str, chr string) string {
	for i, ch := range str {
		if strings.IndexRune(chr, ch) < 0 {
			return str[i:]
		}
	}

	return str
}
