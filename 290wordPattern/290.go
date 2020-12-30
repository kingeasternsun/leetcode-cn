package main

import (
	"fmt"
	"strings"
)

func wordPattern(pattern string, s string) bool {

	words := strings.Split(s, " ")
	var newWords []string
	for _, s := range words {
		if len(s) > 0 {
			newWords = append(newWords, s)
		}
	}
	if len(pattern) != len(newWords) {
		return false
	}

	s2cMap := make(map[string]byte, 0)
	for i, word := range newWords {

		if preByte, ok := s2cMap[word]; ok {
			if preByte != pattern[i] {
				return false
			}
		} else {
			s2cMap[word] = pattern[i]
		}

	}

	c2sMap := make(map[byte]string, 0)
	for i := 0; i < len(pattern); i++ {
		c := pattern[i]
		if preStr, ok := c2sMap[c]; ok {
			if preStr != newWords[i] {
				return false
			}
		} else {
			c2sMap[c] = newWords[i]
		}

	}

	return true

}

func main() {
	fmt.Println(wordPattern("abbc", " dog cat cat  dog"))
	fmt.Println(wordPattern("abba", " dog cat cat  dog"))
	fmt.Println(wordPattern("aaaa", " dog cat cat  dog"))
	fmt.Println(wordPattern("abba", " dog dog dog  dog"))
	fmt.Println(wordPattern("aaaa", " dog dog dog  dog"))
}
