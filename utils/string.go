package utils

import (
	"fmt"
	"strings"
)

type StringUtilities struct{}

func String() *StringUtilities {
	return &StringUtilities{}
}

func (s *StringUtilities) Replacer(base string, replacements []string) string {
	str := base

	for _, replacement := range replacements {
		placeholder := "{}"
		index := strings.Index(base, placeholder)

		if index != -1 {
			str = str[:index] + replacement + str[index+len(placeholder):]
		}
	}

	return str
}

func (s *StringUtilities) Stringify(item any) string {
	return fmt.Sprintf("%v", item)
}
