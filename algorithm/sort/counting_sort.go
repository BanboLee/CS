package sort

func CountingSort(arr []int, k int) []int {
	c := make([]int, k)
	for _, v := range arr {
		c[v] = 1
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
