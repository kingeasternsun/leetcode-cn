package main

import "strings"

func reverseWords(s string) string {

	subs := strings.Split(s, " ")
	var revSubs []string
	// var res string
	for i := len(subs) - 1; i >= 0; i-- {
		if len(subs[i]) > 0 {
			revSubs = append(revSubs, subs[i])

		}
	}
	res := strings.Join(revSubs, " ")
	return res
}
