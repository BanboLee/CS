package sort

func CountingSort(arr []int, k int) []int {
	c := make([]int, k)
	for _, v := range arr {
		c[v]++
	}

	for i := 1; i < k; i++ {
		c[i] += c[i-1]
	}

	b := make([]int, len(arr))
	for i := len(arr) - 1; i >= 0; i-- {
		c[arr[i]]--           // 小于等于arr[i]的数减一
		b[c[arr[i]]] = arr[i] // 填充回去
	}
	return b
}

func CountingSortNegative(arr []int) []int {
	if len(arr) == 0 {
		return nil
	}

	min, max := arr[0], arr[0]
	for _, v := range arr {
		if v > max {
			max = v
		} else if v < min {
			min = v
		}
	}

	ans := make([]int, len(arr))
	count := make([]int, max-min+1)
	for _, v := range arr {
		count[v-min]++
	}
	for i := 1; i < len(count); i++ {
		count[i] += count[i-1]
	}
	for i := len(arr) - 1; i >= 0; i-- {
		count[arr[i]-min]--
		ans[count[arr[i]-min]] = arr[i]
	}

	return ans
}
