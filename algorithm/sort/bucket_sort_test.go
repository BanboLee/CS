package sort

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBucketSort(t *testing.T) {
	arr := []float32{0.1, 0.9, 0.41, 0.42, 0.73, 0.7}
	verify := []float32{0.1, 0.41, 0.42, 0.7, 0.73, 0.9}
	BucketSort(arr)
	assert.EqualValues(t, verify, arr)
}
