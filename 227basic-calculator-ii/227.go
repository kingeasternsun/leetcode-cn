package main

import "strconv"

func main() {
	// calculate("11+32*14/25+4/2")
	calculate("3+2*2")
}

func calculate(s string) int {

	stack := []string{}

	count := len(s)
	var tmpBytes []byte
	for i := 0; i < count; i++ {
		if s[i] == ' ' {
			continue
		}

		switch s[i] {
		case '*', '/', '+', '-':
			stack = append(stack, string(tmpBytes))
			tmpBytes = []byte{s[i]}
			stack = append(stack, string(tmpBytes))
			tmpBytes = []byte{}
		default:
			// stack = append(stack, string(tmpBytes))
			tmpBytes = append(tmpBytes, s[i])

		}

	}
	if len(tmpBytes) > 0 {
		stack = append(stack, string(tmpBytes))
	}

	if len(stack) == 0 {
		return 0
	}

	stack2 := []string{stack[0]}
	for i := 1; i < len(stack); i++ {

		stc2Len := len(stack2)
		switch stack2[stc2Len-1] {
		case "*":
			lv, _ := strconv.Atoi(stack2[stc2Len-2])
			lr, _ := strconv.Atoi(stack[i])
			stack2 = stack2[:len(stack2)-2]
			stack2 = append(stack2, strconv.Itoa(lv*lr))
		case "/":
			lv, _ := strconv.Atoi(stack2[stc2Len-2])
			lr, _ := strconv.Atoi(stack[i])
			stack2 = stack2[:len(stack2)-2]
			stack2 = append(stack2, strconv.Itoa(lv/lr))
		default:
			stack2 = append(stack2, stack[i])

		}

	}

	res, _ := strconv.Atoi(stack2[0])
	for i := 1; i < len(stack2); i++ {

		switch stack2[i-1] {
		case "+":
			lv, _ := strconv.Atoi(stack2[i])
			res = res + lv

		case "-":
			lv, _ := strconv.Atoi(stack2[i])
			res = res - lv
		default:

		}

	}

	return res

}
