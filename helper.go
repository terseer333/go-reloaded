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
func ToUp(str string) string {
	s := strings.Fields(str)
	for i := 0; i < len(s); i++ {
		if s[i] == "(up)" {

			s[i-1] = strings.ToUpper(s[i-1])

			s = append(s[:i], s[i+1:]...)
			i--
		}

	}
	return strings.Join(s, " ")
}

func ToLow(str string) string {
	s := strings.Fields(str)
	for i := 0; i < len(s); i++ {
		if s[i] == "(low)" {
			s[i-1] = strings.ToLower(s[i-1])
			s = append(s[:i], s[i+1:]...)
			i++

		}
	}
	return strings.Join(s, " ")
}
func ToCap(str string) string {
	s := strings.Fields(str)
	for i := 0; i < len(s); i++ {
		if s[i] == "(cap)" {
			s[i-1] = strings.ToUpper(s[i-1][:1]) + strings.ToLower(s[i-1][1:])
			s = append(s[:i], s[i+1:]...)
			i++

		}
	}
	return strings.Join(s, " ")
}

func LastN(str string) string {
	s := strings.Fields(str)
	for i := 0; i < len(s); i++ {
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
