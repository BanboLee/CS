package sort

func ShellSort(arr []int) {
	n := len(arr)
	for i := n / 2; i > 0; i /= 2 {
		for j := i; j < n; j++ {
			tmp := arr[j]
			k := j
			for k >= i && tmp < arr[k-i] {
				arr[k] = arr[k-i]
				k -= i
			}
			arr[k] = tmp
		}
	}
}
