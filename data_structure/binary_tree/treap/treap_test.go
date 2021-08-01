package binary_tree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInsert(t *testing.T) {
	cnt := 100000
	root := Insert(nil, -1)
	for i := 0; i < cnt; i++ {
		root = Insert(root, i)
	}

	for i := -1; i < cnt; i++ {
		assert.NotNil(t, Search(root, i))
		assert.Equal(t, i, Search(root, i).Val)
	}
}

func TestRemove(t *testing.T) {
	cnt := 100000
	root := Insert(nil, -1)
	for i := 0; i < cnt; i++ {
		root = Insert(root, i)
	}

	for i := -1; i < cnt; i++ {
		assert.NotNil(t, Search(root, i))
		assert.Equal(t, i, Search(root, i).Val)
	}

	for i := cnt / 2; i < cnt; i++ {
		root = Remove(root, i)
	}

	for i := -1; i < cnt/2; i++ {
		assert.NotNil(t, Search(root, i))
		assert.Equal(t, i, Search(root, i).Val)
	}
	for i := cnt / 2; i < cnt; i++ {
		assert.Nil(t, Search(root, i))
	}
}
