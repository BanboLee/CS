package sort

func BubbleSort(arr []int) {
	for i := 1; i < len(arr); i++ {
		swapped := false
		for j := 0; j < len(arr)-i; j++ {
			if arr[j] > arr[j+1] {
				swapped = true
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
		if !swapped {
			break
		}
	}
}

func BubbleSortF(arr []float32) {
	for i := 1; i < len(arr); i++ {
		swapped := false
		for j := 0; j < len(arr)-i; j++ {
			if arr[j] > arr[j+1] {
				swapped = true
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
		if !swapped {
			break
		}
	}
}
