package findrepeateddnasequences

func findRepeatedDnaSequences(s string) []string {

	seqMap := make(map[string]int, 0)
	for i := 0; i+10 <= len(s); i++ {
		seqMap[s[i:i+10]]++
	}

	ret := make([]string, 0)
	for k, v := range seqMap {
		if v > 1 {
			ret = append(ret, k)
		}
	}
	return ret

}
