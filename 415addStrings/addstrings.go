package main

import "fmt"

func main() {
	fmt.Println(addStrings("13", "99"))
}

func addStrings(num1 string, num2 string) string {

	i := len(num1) - 1
	j := len(num2) - 1

	k := 0

	res := []byte{}

	for i >= 0 || j >= 0 || k > 0 {

		if i >= 0 && j >= 0 {
			a := int(num1[i] - '0')
			b := int(num2[j] - '0')

			s := a + b + k

			res = append(res, '0'+byte(s%10))
			k = s / 10
			i--
			j--

			continue
		}

		if i >= 0 {
			a := int(num1[i] - '0')
			// b := int(num2[j] - '0')

			s := a + k

			res = append(res, '0'+byte(s%10))
			k = s / 10
			i--
			continue
		}

		if j >= 0 {
			// a := int(num1[i] - '0')
			b := int(num2[j] - '0')

			s := b + k

			res = append(res, '0'+byte(s%10))
			k = s / 10
			j--

			continue
		}

		if k > 0 {
			res = append(res, '0'+byte(k))
			k = 0
		}

	}

	i = 0
	j = len(res) - 1

	// fmt.Println(string(res))

	for i < j {
		res[i], res[j] = res[j], res[i]
		i++
		j--
	}

	return string(res)

}
