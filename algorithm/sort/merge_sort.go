package sort

func merge(arr1, arr2 []int) []int {
	ans := make([]int, len(arr1)+len(arr2))
	l1, l2 := 0, 0
	for l1 < len(arr1) || l2 < len(arr2) {
		if l1 == len(arr1) {
			copy(ans[l1+l2:], arr2[l2:])
			break
		}
		if l2 == len(arr2) {
			copy(ans[l1+l2:], arr1[l1:])
			break
		}
		if arr1[l1] <= arr2[l2] {
			ans[l1+l2] = arr1[l1]
			l1++
		} else {
			ans[l1+l2] = arr2[l2]
			l2++
		}
	}
	return ans
}

func MergeSort(arr []int) []int {
	if len(arr) == 1 {
		return arr
	}

	return merge(MergeSort(arr[:len(arr)/2]), MergeSort(arr[len(arr)/2:]))
}
