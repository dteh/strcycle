package strcycle

import (
	"strconv"
	"strings"
)

// StringModifier applies a modification to a string and returns the result
type StringModifier func(string) string

// UPPER case a string
func UPPER(s string) string {
	return strings.ToUpper(s)
}

// LOWER case a string
func LOWER(s string) string {
	return strings.ToLower(s)
}

// SingleEncode a string with a modifier -- a => %61
func singleEncode(s string, mod func(string) string) string {
	s = mod(s)
	ret := ""
	for _, char := range s {
		ret += `%` + strconv.FormatInt(int64(char), 16)
	}
	return ret
}

func doubleEncodePayload(s string) string {
	return strings.Replace(s, "%", "%25", -1)
}

// DoubleEncode a string -- a => %2561
func doubleEncode(s string, mod func(string) string) string {
	return doubleEncodePayload(singleEncode(s, mod))
}

// DoubleEncodeUpper case of s
func DoubleEncodeUpper(s string) string {
	return doubleEncode(s, UPPER)
}

// SingleEncodeUpper case of s
func SingleEncodeUpper(s string) string {
	return singleEncode(s, UPPER)
}

// DoubleEncodeLower case of s
func DoubleEncodeLower(s string) string {
	return doubleEncode(s, LOWER)
}

// SingleEncodeLower case of s
func SingleEncodeLower(s string) string {
	return singleEncode(s, LOWER)
}
