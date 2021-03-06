package utils

import (
	"regexp"
	"strings"
)

func FindInStringByRange(str string, from string, to string) string {
	fromIndex := 0
	toIndex := 0

	substr := strings.Index(str, from)
	if substr > 0 {
		fromIndex = substr + len(from)
	}

	substr2 := strings.Index(str, to)
	if substr2 > 0 {
		toIndex = substr2
	}

	if toIndex == 0 && fromIndex == 0 {
		return ""
	}

	if toIndex == 0 {
		return str[fromIndex:]
	}

	if fromIndex == 0 {
		return str[:toIndex]
	}

	return str[fromIndex:toIndex]
}

func UcrFirst(str string) string {
	if str == "" {
		return str
	}

	firstChar := string([]rune(str)[0])
	str = str[len(firstChar):]

	return strings.ToLower(firstChar) + str
}

var uppercaseAcronym = map[string]string{
	"ID": "id",
}

// Converts a string to CamelCase
func toCamelInitCase(s string, initCase bool) string {
	s = strings.TrimSpace(s)
	if s == "" {
		return s
	}
	if a, ok := uppercaseAcronym[s]; ok {
		s = a
	}

	n := strings.Builder{}
	n.Grow(len(s))
	capNext := initCase
	for i, v := range []byte(s) {
		vIsCap := v >= 'A' && v <= 'Z'
		vIsLow := v >= 'a' && v <= 'z'
		if capNext {
			if vIsLow {
				v += 'A'
				v -= 'a'
			}
		} else if i == 0 {
			if vIsCap {
				v += 'a'
				v -= 'A'
			}
		}
		if vIsCap || vIsLow {
			n.WriteByte(v)
			capNext = false
		} else if vIsNum := v >= '0' && v <= '9'; vIsNum {
			n.WriteByte(v)
			capNext = true
		} else {
			capNext = v == '_' || v == ' ' || v == '-' || v == '.'
		}
	}
	return n.String()
}

func ToCamelCase(s string) string {
	return toCamelInitCase(s, true)
}

func ToLowerCamelCase(s string) string {
	return toCamelInitCase(s, false)
}

func ToSnakeCase(str string) string {
	var matchFirstCap = regexp.MustCompile("(.)([A-Z][a-z]+)")
	var matchAllCap = regexp.MustCompile("([a-z0-9])([A-Z])")

	snake := matchFirstCap.ReplaceAllString(str, "${1}_${2}")
	snake = matchAllCap.ReplaceAllString(snake, "${1}_${2}")

	return strings.ToLower(snake)
}

func ToKebabCase(str string) string {
	snake := ToSnakeCase(str)
	kebab := strings.ReplaceAll(snake, "_", "-")

	return strings.ToLower(kebab)
}
