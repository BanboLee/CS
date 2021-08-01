package binary_tree

import (
	"banbo/CS/data_structure/binary_tree"
	"testing"

	"github.com/stretchr/testify/assert"
)

func prepareTree(cnt int) *binary_tree.TreeNode {
	var root *binary_tree.TreeNode
	for i := 0; i < cnt; i++ {
		root = Add(root, i+1)
	}
	return root
}

func Test_findEQ(t *testing.T) {
	cnt := 10000
	root := prepareTree(cnt)
	for i := 0; i < cnt; i++ {
		assert.NotNil(t, i+1, findEQ(root, i+1))
		assert.Equal(t, i+1, findEQ(root, i+1).Val)
	}
}

func Test_find(t *testing.T) {
	cnt := 10
	root := prepareTree(cnt)
	for i := -10; i <= 0; i++ {
		assert.Equal(t, 1, find(root, i))
	}

	for i := 1; i <= 10; i++ {
		assert.Equal(t, i, find(root, i))
	}

	for i := 11; i <= 100; i++ {
		assert.Equal(t, -1, find(root, i))
	}
}

func TestAdd(t *testing.T) {
	cnt := 10000
	root := prepareTree(cnt)
	var inOrder []int
	InOrder(root, &inOrder)

	for i, v := range inOrder {
		assert.Equal(t, v, i+1)
	}
}

func TestRemove(t *testing.T) {
	cnt := 10000
	root := prepareTree(cnt)
	root = Remove(root, cnt/2)
	var inOrder []int
	InOrder(root, &inOrder)
	var cases []int
	for i := 1; i <= cnt; i++ {
		if i != cnt/2 {
			cases = append(cases, i)
		}
	}
	assert.EqualValues(t, cases, inOrder)
	assert.Nil(t, findEQ(root, cnt/2))
	root = Add(root, cnt/2)
	assert.NotNil(t, findEQ(root, cnt/2))
	assert.Equal(t, cnt/2, find(root, cnt/2))
}

func TestInOrder(t *testing.T) {
	cnt := 10000
	var cases []int
	for i := 1; i <= cnt; i++ {
		cases = append(cases, i)
	}
	root := prepareTree(cnt)
	var inOrder []int
	InOrder(root, &inOrder)
	assert.EqualValues(t, cases, inOrder)
}

func TestPostOrder(t *testing.T) {
	cnt := 10000
	var cases []int
	for i := cnt; i >= 1; i-- {
		cases = append(cases, i)
	}
	root := prepareTree(cnt)
	var postOrder []int
	PostOrder(root, &postOrder)
	assert.EqualValues(t, cases, postOrder)
}
