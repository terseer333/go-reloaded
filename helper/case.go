package helper

import (
	"strconv"
	"strings"
)

func Case(str string) string {
	s := strings.Fields(str)
	for i := 0; i < len(s); i++ {
		if s[i] == "(up)" && i+1 < len(s) {

			s[i-1] = strings.ToUpper(s[i-1])

			s = append(s[:i], s[i+1:]...)
			i--
		}

		if s[i] == "(low)" && i > 0 {
			s[i-1] = strings.ToLower(s[i-1])
			s = append(s[:i], s[i+1:]...)
			i--

		}

		if s[i] == "(cap)" && i > 0 {
			s[i-1] = strings.ToUpper(s[i-1][:1]) + strings.ToLower(s[i-1][1:])
			s = append(s[:i], s[i+1:]...)
			i--

		}

		if s[i] == "(low," && i+1 < len(s) {

			num := strings.TrimSuffix(s[i+1], ")")
			s = append(s[:i], s[i+2:]...)
			i--
			count, _ := strconv.Atoi(num)

			start := i - count + 1
			if start < 0 {
				start = 0
			}
			for j := start; j <= i; j++ {
				s[j] = strings.ToLower(s[j])

			}
		}
		if s[i] == "(cap," && i+1 < len(s) {
			num := strings.TrimSuffix(s[i+1], ")")
			s = append(s[:i], s[i+2:]...)
			i--
			count, _ := strconv.Atoi(num)

			start := i - count + 1
			if start < 0 {
				start = 0
			}
			for j := start; j <= i; j++ {
				s[j] = strings.ToUpper(s[j][:1]) + strings.ToLower(s[j][1:])

			}
		}
		if s[i] == "(up," && i+1 < len(s) {
			num := strings.TrimSuffix(s[i+1], ")")
			s = append(s[:i], s[i+2:]...)
			i--
			count, _ := strconv.Atoi(num)

			start := i - count + 1
			if start < 0 {
				start = 0
			}
			for j := start; j <= i; j++ {
				s[j] = strings.ToUpper(s[j])

			}
		}

	}
	return strings.Join(s, " ")
}
