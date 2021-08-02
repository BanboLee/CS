package sort

import (
	"math/rand"
	"time"
)

func QuickSort(arr []int) {
	quickSort(arr, 0, len(arr))
}

func quickSort(arr []int, i, n int) {
	if n <= 1 {
		return
	}

	rand.Seed(time.Now().UnixNano())
	pivot := arr[rand.Intn(n)+i] // 选取任意一个数为标杆
	// make arr[i, p] < pivot, arr[p+1, q-1] = x, arr[q, i+n-1] > x
	p := i - 1 // 小于标杆的集合左边界， 一开始为空
	j := i     // 遍历整个区间的下标
	q := i + n // 大于标杆的集合右边界， 一开始为空
	for j < q {
		if arr[j] < pivot {
			// 小于标杆， 则将其放到左边去
			p++ // 左边集合+1
			arr[j], arr[p] = arr[p], arr[j]
			j++ // j已被置换， 遍历下一个
		} else if arr[j] > pivot {
			// 大于标杆， 则将其放到右边去
			q-- // 右边集合+1
			arr[j], arr[q] = arr[q], arr[j]
			// 当前j虽然被换过来，但不一定就属于左边
		} else {
			j++
		}
	}

	quickSort(arr, i, p-i+1)
	quickSort(arr, p+1, n-q+i)
}
