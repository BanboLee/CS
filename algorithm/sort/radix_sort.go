package sort

func RadixSort(arr []int) []int {
	d, w := 8, 64
	for i := 0; i < w/d; i++ {
		count := make([]int, 1<<d)
		for _, v := range arr {
			count[(v>>(i*d))&(1<<d-1)]++
		}
		for j := 1; j < 1<<d; j++ {
			count[j] += count[j-1]
		}
		tmp := make([]int, len(arr))
		for j := len(arr) - 1; j >= 0; j-- {
			count[(arr[j]>>(i*d))&(1<<d-1)]--
			tmp[count[(arr[j]>>(i*d))&(1<<d-1)]] = arr[j]
		}

		arr = tmp
	}
	return arr
}
