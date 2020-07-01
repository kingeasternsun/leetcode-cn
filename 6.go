package main

import "fmt"

// func main() {
// 	var s = "A"
// 	var numRows = 1
// 	r := (convert(s, numRows))
// 	if r != "A" {
// 		fmt.Println("error", r)
// 	}
// }

func convert(s string, numRows int) string {

	if numRows == 1 {
		return s
	}

	sLen := len(s)
	sb := make([]byte, sLen)

	k := 0
	for i := 0; i < numRows; i++ {

		step := 2*numRows - 2 - i*2

		for j := i; j < sLen; j = j + 2*numRows - 2 {

			// sb = append(sb,(s[j]))
			if j < sLen {
				fmt.Println(j, sLen)
				sb[k] = s[j]
				k = k + 1
			}

			if i != 0 && i != numRows-1 && j+step < sLen {

				// sb = append(sb, (s[j+step]))
				sb[k] = s[j+step]
				k = k + 1
			}

		}

	}

	return string(sb)
}
