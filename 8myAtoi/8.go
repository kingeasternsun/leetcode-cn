package main

import (
	"fmt"
	"math"
)

func main() {

	fmt.Println(myAtoi("42") == 42)
	fmt.Println(myAtoi("42 ") == 42)
	fmt.Println(myAtoi("42-"), myAtoi("42-") == 42)
	fmt.Println(myAtoi("42t"), myAtoi("42t") == 42)
	fmt.Println(myAtoi(" 42") == 42)
	fmt.Println(myAtoi("-42 "), myAtoi("-42 ") == -42)
	fmt.Println(myAtoi("+42-") == 42)
	fmt.Println(myAtoi("t42t") == 0)
	fmt.Println(myAtoi("- 42 ") == 0)
	fmt.Println(myAtoi("+ 42-") == 0)
	fmt.Println(myAtoi("t 42t") == 0)
	fmt.Println(myAtoi("  -42") == -42)
	fmt.Println(myAtoi(" --4 2 - 3") == 0)
	fmt.Println(myAtoi("4193 with words") == 4193)
	fmt.Println(myAtoi("-91283472332") == -2147483648)
	fmt.Println(myAtoi("9223372036854775808") == 2147483647)
	// fmt.Println(myAtoi("-91283472332"))
}

func myAtoi(s string) int {

	a := SM{S: s}
	a.Init()

	return a.Atoi()
}

/*
 1: 表示 start
 2： 表示 numbert
 3： 表示 符号
 4:  表示空格
 5:  表示其他

 利用状态机
*/

const (
	NUMBER = 2
	SIGN   = 3
	SPACE  = 4
	OTHER  = 5
)

type Status struct {
	CurStatus int
	Opt       int
}

type SM struct {
	S         string
	CurStatus int
	MoveMap   map[Status]int
	Sign      int64
	Number    int64
}

func (sm *SM) Init() {
	sm.MoveMap = make(map[Status]int, 0)
	sm.MoveMap[Status{CurStatus: 1, Opt: NUMBER}] = NUMBER
	sm.MoveMap[Status{CurStatus: 1, Opt: SIGN}] = SIGN
	sm.MoveMap[Status{CurStatus: 1, Opt: SPACE}] = 1
	sm.MoveMap[Status{CurStatus: 1, Opt: 5}] = 5
	// sm.MoveMap[Status{CurStatus: 1, Opt: 6}] = 5

	sm.MoveMap[Status{CurStatus: NUMBER, Opt: NUMBER}] = NUMBER
	sm.MoveMap[Status{CurStatus: NUMBER, Opt: SIGN}] = OTHER
	sm.MoveMap[Status{CurStatus: NUMBER, Opt: SPACE}] = 5
	sm.MoveMap[Status{CurStatus: NUMBER, Opt: 5}] = 5

	sm.MoveMap[Status{CurStatus: SIGN, Opt: NUMBER}] = NUMBER
	sm.MoveMap[Status{CurStatus: SIGN, Opt: SIGN}] = 5
	sm.MoveMap[Status{CurStatus: SIGN, Opt: SPACE}] = 5
	sm.MoveMap[Status{CurStatus: SIGN, Opt: 5}] = 5

	sm.MoveMap[Status{CurStatus: SPACE, Opt: NUMBER}] = NUMBER
	sm.MoveMap[Status{CurStatus: SPACE, Opt: SIGN}] = SIGN
	sm.MoveMap[Status{CurStatus: SPACE, Opt: SPACE}] = SPACE
	sm.MoveMap[Status{CurStatus: SPACE, Opt: 5}] = 5

	sm.MoveMap[Status{CurStatus: 5, Opt: NUMBER}] = 5
	sm.MoveMap[Status{CurStatus: 5, Opt: 3}] = 5
	sm.MoveMap[Status{CurStatus: 5, Opt: 5}] = 5
	sm.MoveMap[Status{CurStatus: 5, Opt: 5}] = 5

	sm.Sign = 1
	sm.CurStatus = 1

}

func (sm *SM) move(op int) int {

	var ok bool
	sm.CurStatus, ok = sm.MoveMap[Status{sm.CurStatus, op}]
	if !ok {
		sm.CurStatus = OTHER
	}

	return sm.CurStatus
}

func (sm *SM) Move(c byte) int {

	var op int

	if c == ' ' {
		op = SPACE
	} else if c == '-' {
		op = SIGN
		//要在这里判断数字符号,
		if sm.CurStatus == 1 {
			sm.Sign = -1
		}

	} else if c == '+' {
		op = SIGN
	} else if c >= '0' && c <= '9' {
		op = NUMBER
		sm.Number = sm.Number*10 + int64(c-'0')
	} else {
		op = OTHER
	}

	return sm.move(op)
}

func (sm *SM) Atoi() int {

	num := len(sm.S)
	for i := 0; i < num; i++ {
		t := sm.Move(sm.S[i])
		if t == OTHER {
			break
		}

		//要在这里提前判断是否溢出
		if sm.Number > math.MaxInt32 {

			if sm.Sign == 1 {
				return math.MaxInt32
			}
			return math.MinInt32

		}

	}

	return int(sm.Number * sm.Sign)

}
