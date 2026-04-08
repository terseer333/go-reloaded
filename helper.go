package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

func HexToDec(str string) string {
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
	}
	return strings.Join(s, " ")
}
func BinToDec(str string) string {
	s := strings.Fields(str)
	for i := 0; i < len(s); i++ {
		if s[i] == "(bin)" {
			ja, err := strconv.ParseInt(s[i-1], 16, 64)
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
func ToUp(str string) string {
	s := strings.Fields(str)
	for i := 0; i < len(s); i++ {
		if s[i] == "(up)" {
			if i > 0 {
				s[i-1] = strings.ToUpper(s[i-1])

				s = append(s[:1], s[i+1:]...)
				i--
			}

		}

	}
	return strings.Join(s, " ")
}
