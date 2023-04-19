package stringutils

import "unicode/utf8"

func BytesCount(s string)int{
	return len(s)
}

func CharsCount(s string) int{
	return utf8.RuneCountInString(s)
}