package main

func main() {

}

// 914. 卡牌分组 https://leetcode-cn.com/problems/x-of-a-kind-in-a-deck-of-cards/
func hasGroupsSizeX1(deck []int) bool {

	cnt := make([]int, 10001)
	count := len(deck)
	if count <= 1 {
		return false
	}
	for i := 0; i < count; i++ {
		cnt[deck[i]] = cnt[deck[i]] + 1
	}

	g := -1
	for i := 0; i < 10001; i++ {
		if cnt[i] == 1 {
			return false
		}

		if g == -1 {
			g = cnt[i]
		} else {
			g = gcd(g, cnt[i])
		}

	}

	if g == 1 {
		return false
	}

	return true

}

func hasGroupsSizeX(deck []int) bool {

	cnt := make(map[int]int, 0)
	count := len(deck)
	if count <= 1 {
		return false
	}
	for i := 0; i < count; i++ {
		cnt[deck[i]] = cnt[deck[i]] + 1
	}

	g := -1
	for k := range cnt {
		if cnt[k] == 1 {
			return false
		}

		if g == -1 {
			g = cnt[k]
		} else {
			g = gcd(g, cnt[k])
		}

	}

	if g == 1 {
		return false
	}

	return true

}
func gcd(x, y int) int {
	if x > y {
		y, x = x, y
	}

	if x == 0 {
		return y
	}

	return gcd(y%x, x)
}
