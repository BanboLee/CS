package sort

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInsertionSort(t *testing.T) {
	cnt := 1000
	arr := make([]int, cnt)
	verify := make([]int, cnt)
	for i := range arr {
		arr[i] = cnt - i - 1
		verify[i] = i
	}
	SelectionSort(arr)
	assert.EqualValues(t, verify, arr)
}
