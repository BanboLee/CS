package sort

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShellSort(t *testing.T) {
	cnt := 1000
	arr := make([]int, cnt)
	verify := make([]int, cnt)
	for i := range arr {
		arr[i] = cnt - i - 1
		verify[i] = i
	}
	ShellSort(arr)
	assert.EqualValues(t, verify, arr)
}
