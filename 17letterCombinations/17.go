package main

func main() {

}

//dfs
func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return nil
	}
	m := map[byte][]byte{
		'2': []byte{'a', 'b', 'c'},
		'3': []byte{'d', 'e', 'f'},
		'4': []byte{'g', 'h', 'i'},
		'5': []byte{'j', 'k', 'l'},
		'6': []byte{'m', 'n', 'o'},
		'7': []byte{'p', 'q', 'r', 's'},
		'8': []byte{'t', 'u', 'v'},
		'9': []byte{'w', 'x', 'y', 'z'},
	}

	dbytes := []byte(digits)
	res := make([]string, 0)
	tmp := make([]byte, 0)
	dfs(m, dbytes, tmp, &res)
	return res
}

func dfs(m map[byte][]byte, digits []byte, tmp []byte, res *[]string) {

	if len(digits) == 0 {
		*res = append(*res, string(tmp))
		return
	}

	for _, b := range m[digits[0]] {
		dfs(m, digits[1:], append(tmp, b), res)
	}

}
