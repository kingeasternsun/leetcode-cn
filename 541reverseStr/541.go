/*
 * @Description:
 * @Version: 2.0
 * @Author: kingeasternsun
 * @Date: 2021-02-05 17:36:11
 * @LastEditors: kingeasternsun
 * @LastEditTime: 2021-02-06 09:59:39
 * @FilePath: \541reverseStr\541.go
 */
package leetcode

func reverseStr(s string, k int) string {
	if k == 1 || len(s) <= 1 {
		return s
	}

	in := []byte(s)
	cnt := len(in) / (2 * k)
	for i := 0; i < len(in)/(2*k); i++ {
		reverse(in, i*2*k, i*2*k+k-1)
	}

	if len(in)%(2*k) >= k {
		reverse(in, cnt*2*k, cnt*2*k+k-1)
	} else {
		reverse(in, cnt*2*k, len(in)-1)
	}

	return string(in)

}

func reverse(word []byte, i, j int) {
	if i >= j {
		return
	}
	for ; i < j; i, j = i+1, j-1 {
		word[i], word[j] = word[j], word[i]
	}
	return
}

//另外一种
func reverseStr1(s string, k int) string {
	if k == 1 || len(s) <= 1 {
		return s
	}

	in := []byte(s)
	cnt := len(in) / (2 * k)
	for i := 0; i < len(in)/(2*k); i++ {
		reverse1(in[i*2*k : i*2*k+k])
	}

	if len(in)%(2*k) >= k {
		reverse1(in[cnt*2*k : cnt*2*k+k])
	} else {
		reverse1(in[cnt*2*k:])
	}

	return string(in)

}

func reverse1(word []byte) {
	if len(word) <= 1 {
		return
	}

	for i, j := 0, len(word)-1; i < j; i, j = i+1, j-1 {
		word[i], word[j] = word[j], word[i]
	}
	return
}
