package main

import "fmt"

func main() {
	fmt.Println(findLongestWord("abpcplea", []string{"ale", "apple", "monkey", "plea"}) == "apple")
	fmt.Println(findLongestWord("abpcplea", []string{"a", "b", "c"}) == "a")
	fmt.Println(findLongestWord("abpcplea", []string{"tt", "", ""}) == "")
	fmt.Println(findLongestWord("abpcplea", []string{"tt", "", "c"}) == "c")
}

func findLongestWord(s string, d []string) string {

	if len(s) == 0 {
		return ""
	}
	if len(d) == 0 {
		return ""
	}

	var res string

	for _, t := range d {

		if match(s, t) {

			if len(t) > len(res) {
				res = t
			} else if len(t) == len(res) && t < res {
				res = t
			}
		}

	}

	return res
}

func match(s, d string) bool {

	if len(d) == 0 {
		return false
	}

	if len(s) < len(d) {
		return false
	}

	i := 0
	j := 0
	for i < len(s) && j < len(d) {

		if s[i] == d[j] {
			i++
			j++
			continue
		}
		i++

	}

	if j == len(d) {
		return true
	}

	return false
}
