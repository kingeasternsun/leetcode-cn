package main

import "strconv"

func evalRPN(tokens []string) int {

	exp := make([]int, 0)

	for i := range tokens {

		if tokens[i] != "+" && tokens[i] != "-" && tokens[i] != "*" && tokens[i] != "/" {
			n, _ := strconv.Atoi(tokens[i])
			exp = append(exp, n)
			continue
		}

		n2 := exp[len(exp)-1]
		n1 := exp[len(exp)-2]

		var n int 
		switch tokens[i] {
		case "+":
			n = n1 + n2
		case "-":
			n = n1 - n2
		case "*":
			n = n1 * n2
		case "/":
			n = n1 / n2

		}

		exp[len(exp)-2] = n
		exp = exp[:len(exp)-1]

	}

	return exp[0]

}
