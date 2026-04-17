package helper

import (
	"fmt"

	"log"
	"strconv"
	"strings"
)

func baseN(str string) string {
	s := strings.Fields(str)
	for i := 0; i < len(s); i++ {
		if s[i] == "(hex," && i+1 < len(s) {
			num := strings.TrimSuffix(s[i+1], ")")
			s = append(s[:i], s[i+2:]...)
			i--

			count, _ := strconv.Atoi(num)

			start := i - count + 1
			if start < 0 {
				start = 0
			}
			for j := start; j <= i; j++ {
				ja, err := strconv.ParseInt(s[j], 16, 64)
				if err != nil {
					log.Fatal(err)
					if j < 0 {
						j = 0
					}
				}
				s[j] = fmt.Sprint(ja)
			}

		}
	}
	return strings.Join(s, " ")
}
