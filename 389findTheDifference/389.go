package main

func main() {

}
func findTheDifference(s string, t string) byte {

	smap := make([]int, 26)
	tmap := make([]int, 26)

	for i := 0; i < len(s); i++ {

		smap[s[i]-'a']++
	}

	for i := 0; i < len(t); i++ {

		tmap[t[i]-'a']++
	}

	for i := 0; i < 26; i++ {

		if smap[i] != tmap[i] {

			return byte(i + 'a')
		}

	}

	return byte('a')

}
