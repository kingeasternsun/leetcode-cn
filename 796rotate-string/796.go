package main

import "strings"

func rotateString(A string, B string) bool {

	if len(A) != len(B) {
		return false
	}

	return strings.Contains(A+A, B)

}
