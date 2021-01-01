package leetcode

//20ms 6MB
func canPlaceFlowers1(flowerbed []int, n int) bool {
	if n == 0 {
		return true
	}

	for i, v := range flowerbed {

		if n == 0 {
			return true
		}

		//贪心方法，如果当前可以放置，就提前放置

		//提前处理不能放置的情况
		//已经种花了
		if v == 1 {
			continue
		}

		//左边已经种花了
		if i > 0 && flowerbed[i-1] == 1 {
			continue
		}

		//右边已经种花了
		if i+1 < len(flowerbed) && flowerbed[i+1] == 1 {
			continue
		}

		//当前放花
		flowerbed[i] = 1
		n--

	}
	return n == 0
}

//24ms 6MB
func canPlaceFlowers(flowerbed []int, n int) bool {
	if n == 0 {
		return true
	}

	i := 0
	for i < len(flowerbed) {

		//如果没有可以插入的花朵，那么就可以提前返回了
		if n == 0 {
			return true
		}

		//贪心方法，如果当前可以放置，就提前放置

		//提前处理不能放置的情况

		//当前位置已经种花了，不仅当前位置不能种花，下个位置肯定也不能种花了,直接跳两次
		if flowerbed[i] == 1 {
			i += 2
			continue
		}

		//左边已经种花了,当前位置不能种花，下个位置也许可以，跳一次
		if i > 0 && flowerbed[i-1] == 1 {
			i++
			continue
		}

		//右边已经种花了，不仅当前位置不能种花，下个位置不能种花，下下个位置也不能种花，所以跳三次
		if i+1 < len(flowerbed) && flowerbed[i+1] == 1 {
			i += 3
			continue
		}

		//当前放花,下个位置可定不能放花，下下个位置可能放,跳两次
		flowerbed[i] = 1
		n--
		i += 2
	}
	return n == 0
}
