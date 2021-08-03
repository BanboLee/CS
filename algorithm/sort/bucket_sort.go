package sort

func BucketSort(arr []float32) {
	buckets := make([][]float32, 10)
	for _, v := range arr {
		buckets[int(v*10)] = append(buckets[int(v*10)], v)
	}

	i := 0
	for _, bucket := range buckets {
		if bucket != nil {
			BubbleSortF(bucket)
			for _, v := range bucket {
				arr[i] = v
				i++
			}
		}
	}
}
