package helper

import "regexp"

func Punc(str string) string {

	puc := regexp.MustCompile(`([a-zA-Z0-9])\s+([,.!:;?]+)([a-zA-Z0-9])`)
	str = puc.ReplaceAllString(str, "$1$2 $3")

	puc2 := regexp.MustCompile(`\s+([,.:;?!])`)
	str = puc2.ReplaceAllString(str, "$1")

	space := regexp.MustCompile(`\s{2,}(\s[a-zA-Z0-9])`)
	str = space.ReplaceAllString(str, "$1")
	return str
}
