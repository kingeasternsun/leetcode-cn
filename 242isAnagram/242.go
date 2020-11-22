package main

import "fmt"

func main() {

	fmt.Println(isAnagram("anagram", "nagaram"))
	fmt.Println(isAnagram("rat", "car"))
	fmt.Println(isAnagram("abc", "ab"))
	fmt.Println(isAnagram("ab", "cab"))
	fmt.Println(isAnagram("abc", "bac"))

}

//8ms 2.8MB
func isAnagram1(s string, t string) bool {

	sMap := make(map[rune]int, 0)
	tMap := make(map[rune]int, 0)

	for _, r := range s {
		sMap[r] = sMap[r] + 1
	}

	for _, r := range t {
		tMap[r] = tMap[r] + 1
	}

	if len(sMap) != len(tMap) {
		return false
	}

	for k, v := range sMap {

		if v != tMap[k] {
			return false
		}
	}

	return true

}

//8ms 2.8MB
func isAnagram(s string, t string) bool {

	sMap := make(map[rune]int, 0)
	// tMap := make(map[rune]int, 0)

	for _, r := range s {
		sMap[r] = sMap[r] + 1
	}

	for _, r := range t {

		tv, ok := sMap[r]
		if !ok {
			return false
		}

		tv = tv - 1
		if tv == 0 {
			delete(sMap, r)
		} else {
			sMap[r] = tv
		}

	}

	return (len(sMap) == 0)

}
