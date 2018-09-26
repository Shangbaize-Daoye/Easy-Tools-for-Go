package string_handler

import (
	"strings"
)

func Map(mapping func(rune) rune, s *string) {
	*s = strings.Map(mapping, *s)
}

func Repeat(s *string, count int) {
	*s = strings.Repeat(*s, count)
}

func Replace(s *string, old, new string, n int) {
	*s = strings.Replace(*s, old, new, n)
}

func Capitalize(s *string) {
	if len(*s) != 0 {
		*s = strings.ToUpper((*s)[:1]) + strings.ToLower((*s)[1:])
	}
}

func Title(s *string) {
	*s = strings.Title(*s)
}

func Lower(s *string) {
	*s = strings.ToLower(*s)
}

func Upper(s *string) {
	*s = strings.ToUpper(*s)
}

func Trim(s *string, cutset string) {
	*s = strings.Trim(*s, cutset)
}

func TrimLeft(s *string, cutset string) {
	*s = strings.TrimLeft(*s, cutset)
}

func TrimPrefix(s *string, prefix string) {
	*s = strings.TrimPrefix(*s, prefix)
}

func TrimRight(s *string, cutset string) {
	*s = strings.TrimRight(*s, cutset)
}

func TrimSpace(s *string) {
	*s = strings.TrimSpace(*s)
}

func TrimSuffix(s *string, suffix string) {
	*s = strings.TrimSuffix(*s, suffix)
}
