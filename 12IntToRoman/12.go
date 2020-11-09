package main

import (
	"fmt"
	"strings"
)

func main() {

	fmt.Println(intToRoman(3) == "III")
	fmt.Println(intToRoman(4) == "IV")
	fmt.Println(intToRoman(5) == "V")
	fmt.Println(intToRoman(9) == "IX")
	fmt.Println(intToRoman(10) == "X")
	fmt.Println(intToRoman(50) == "L")
	fmt.Println(intToRoman(90) == "XC")
	fmt.Println(intToRoman(58) == "LVIII")
	fmt.Println(intToRoman(1994) == "MCMXCIV")
	fmt.Println(intToRoman(3999) == "MMMCMXCIX")
}

func intToRoman(num int) string {

	m := map[int]string{1: "I", 5: "V", 10: "X", 50: "L", 100: "C", 500: "D", 1000: "M", 4: "IV", 9: "IX", 40: "XL", 90: "XC", 400: "CD", 900: "CM"}

	res := ""
	for num > 0 {

		_, width := intParse(num)

		// if h == 1 {
		// 	res += m[width]
		// 	num -= width
		// 	continue
		// }

		for _, v := range []int{9, 5, 4} {

			if num >= v*width {

				res += m[v*width]
				num -= v * width
				break
			}

		}

		if num/width > 0 {
			res += strings.Repeat(m[width], num/width)
			num -= width * (num / width)
		}

	}
	return res
}

func intParse(num int) (h int, width int) {

	width = 1
	for num >= 10 {
		width = width * 10
		num = num / 10
	}

	return h, width
}
