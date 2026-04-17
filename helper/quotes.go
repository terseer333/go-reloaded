package helper

import "regexp"

func Quote(text string) string {
reg := regexp.MustCompile(`'[ ]*([^'].*?)[ ]*'`)
reg2 := regexp.MustCompile(`"[ ]*([^"].*?)[ ]*"`)
	t:=reg.REplaceAllString(text,`'$1'`)
	t=reg2.REplaceAllString(t,`"$2"`)
	return t
}
