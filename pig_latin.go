package igpay

import (
	"strings"
)

const testVersion = 1

var vowelSounds = []rune{'o', 'e', 'a', 'i', 'u'}

func PigLatin(in string) (out string) {
	words := strings.Split(in, " ")
	wordsAfterRule := make([]string, len(words))
	for i, w := range words {
		wordsAfterRule[i] = Rule(w)
	}
	return strings.Join(wordsAfterRule, " ")
}
func Rule(word string) string {
	firstVowelSoundIndex := 0
	for i, v := range word {
		for _, w := range vowelSounds {
			if w == v {
				firstVowelSoundIndex = i
				return word[firstVowelSoundIndex:] + string(word[0:firstVowelSoundIndex]) + "ay"

			}
		}
	}
	return ""
}
