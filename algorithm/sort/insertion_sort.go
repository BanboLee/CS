package sort

func InsertionSort(arr []int) {
	for i := 1; i < len(arr); i++ {
		cur := arr[i]
		j := i - 1
		for j >= 0 && arr[j] > cur {
			arr[j+1] = arr[j]
			j--
		}
		arr[j] = cur
	}
}
