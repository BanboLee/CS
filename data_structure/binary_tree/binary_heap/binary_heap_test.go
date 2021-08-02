package binary_heap

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBinaryMinHeap_shiftUp(t *testing.T) {
	heap := &BinaryMinHeap{
		data: []int{7, 6, 5, 4, 3, 2, 1},
		n:    7,
	}
	heap.shiftUp(heap.n - 1)
	assert.Equal(t, 1, heap.data[0])
}

func TestBinaryMinHeap_Add(t *testing.T) {
	heap := &BinaryMinHeap{}
	cnt := 10000
	for i := 0; i < cnt; i++ {
		heap.Add(i)
	}
	for i, v := range heap.data {
		assert.Equal(t, i, v)
	}
}

func TestBinaryMinHeap_shiftDown(t *testing.T) {
	heap := &BinaryMinHeap{
		data: []int{7, 6, 5, 4, 3, 2, 1},
		n:    7,
	}
	heap.shiftDown(0)
	assert.Equal(t, 7, heap.data[heap.n-1])
}

func TestBinaryMinHeap_Remove(t *testing.T) {
	heap := &BinaryMinHeap{}
	cnt := 10000
	for i := cnt - 1; i >= 0; i-- {
		heap.Add(i)
	}
	for i := 0; i < cnt; i++ {
		assert.Equal(t, i, heap.Remove())
	}
}
