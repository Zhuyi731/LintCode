package main

import (
	"fmt"
	"unicode"
)

func main() {
	fmt.Println(shortestCompletingWord("1s3 PSt", []string{"step", "steps", "stripe", "stepple"}))
	fmt.Println(shortestCompletingWord("1s3 456", []string{"looks", "pest", "stew", "show"}))
	fmt.Println(shortestCompletingWord("Ah71752", []string{"suggest", "letter", "of", "husband", "easy", "education", "drug", "prevent", "writer", "old"}))
	fmt.Println(shortestCompletingWord("OgEu755", []string{"enough", "these", "play", "wide", "wonder", "box", "arrive", "money", "tax", "thus"}))
	fmt.Println(shortestCompletingWord("iMSlpe4", []string{"claim", "consumer", "student", "camera", "public", "never", "wonder", "simple", "thought", "use"}))
}

func shortestCompletingWord(licensePlate string, words []string) string {
	sourceMap := map[rune]int{}
	for _, r := range licensePlate {
		if unicode.IsLower(r) {
			sourceMap[r]++
		} else if unicode.IsUpper(r) {
			sourceMap[unicode.ToLower(r)]++
		}
	}

	currentMin := ""
	for _, word := range words {
		nowMp := map[rune]int{}
		for _, r := range word {
			nowMp[r]++
		}
		can := true
		for k, v := range sourceMap {
			if v > nowMp[k] {
				can = false
				break
			}
		}
		if !can {
			continue
		}
		if currentMin == "" || len(word) < len(currentMin) {
			currentMin = word
		}
	}
	return currentMin
}
