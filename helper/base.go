package helper

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

func Base(str string) string {
	s := strings.Fields(str)
	for i := 0; i < len(s); i++ {
		if s[i] == "(hex)" {
			ja, err := strconv.ParseInt(s[i-1], 16, 64)
			if err != nil {
				log.Fatal(err)
			}
			s[i-1] = fmt.Sprint(ja)
			s = append(s[:i], s[i+1:]...)
			i--
		}

		if s[i] == "(bin)" {
			ja, err := strconv.ParseInt(s[i-1], 2, 64)
			if err != nil {
				log.Fatal(err)

			}
			s[i-1] = fmt.Sprint(ja)
			s = append(s[:i], s[i+1:]...)
			i--
		}

	}
	return strings.Join(s, " ")
}
