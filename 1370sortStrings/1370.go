package main

import "fmt"

func main() {

	fmt.Println(sortString("aaaabbbbcccc") == "abccbaabccba")
	fmt.Println(sortString("leetcode") == "cdelotee")
	fmt.Println(sortString("a") == "a")
	fmt.Println(sortString("") == "")
	fmt.Println(sortString("ggggggg") == "ggggggg")
}

func sortString(s string) string {

	scnt := len(s)
	cmap := make([]int, 26)
	for i := 0; i < scnt; i++ {
		cmap[int(s[i]-'a')]++
	}

	res := make([]byte, 0)
	flat := true
	for {

		if len(res) == scnt {
			return string(res)
		}

		if flat {
			for i := 0; i < 26; i++ {

				if cmap[i] > 0 {
					res = append(res, 'a'+byte(i))
					cmap[i]--
				}

			}
		} else {
			for i := 25; i >= 0; i-- {

				if cmap[i] > 0 {
					res = append(res, byte(i)+'a')
					cmap[i]--
				}

			}
		}
		flat = !flat

	}

	return string(res)

}
