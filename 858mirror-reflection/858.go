package main

func mirrorReflection(p int, q int) int {

	if q == 0 {
		return 0
	}

	if p == q {
		return 1
	}

	//刚好是整除
	if p%q == 0 {
		n := p / q
		if n&1 == 1 {
			return 1
		}
		return 2
	}

	g := gcd(p, q)
	hNum := q / g //垂直方向需要几个房间
	vNum := p / g //水平方向需要几个房间

	if hNum&1 == 1 {

		if vNum&1 == 1 {
			return 1
		}

		return 2

	}

	if vNum&1 == 1 {
		return 0
	}

	return 0

}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}

	return gcd(b, a%b)
}
