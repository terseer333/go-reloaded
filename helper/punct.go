package helper

import "regexp"

func Punc(str string) string {

	puc := regexp.MustCompile(`[ ]*([,.:;?!]+)`)
	return puc.ReplaceAllString(str, `$1`)

	

}
