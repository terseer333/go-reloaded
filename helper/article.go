package helper

import (
	"strings"
)

func Article(text string) string {

	words := strings.Fields(text)
	vowel := "aeiouhAEIOUH"
	consonant := "bcdfgjklmnpqrstvwxyzBCDFGJKLMNPQRSTVWXYZ"

	for i := 0; i < len(words); i++ {
		if (words[i] == "A" || words[i] == "a") && strings.ContainsAny(words[i+1][:1], vowel) {
			switch words[i] {
			case "A":
				words[i] = "An"
			case "a":
				words[i] = "an"
			}
		}

		if (words[i] == "An" || words[i] == "an") && strings.ContainsAny(words[i+1][:1], consonant) {
			switch words[i] {
			case "An":
				words[i] = "A"
			case "an":
				words[i] = "a"
			}
		}

	}
	return strings.Join(words, " ")
}
