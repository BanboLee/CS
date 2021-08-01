package skiplist_list

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestList_findPre(t *testing.T) {
	list := NewList()
	cnt := 90000
	for i := 0; i < cnt; i++ {
		list.Add(i, i+1)
	}
	for i := 0; i < cnt; i++ {
		pre := list.findPre(i)
		assert.Equal(t, i, pre.next[0].Val-1)
	}
}

func TestList_Get(t *testing.T) {
	list := NewList()
	cnt := 900000
	for i := 0; i < cnt; i++ {
		list.Add(i, i+1)
	}
	for i := 0; i < cnt; i++ {
		assert.Equal(t, i+1, list.Get(i))
	}
}

func TestList_Set(t *testing.T) {
	list := NewList()
	cnt := 1000000
	for i := 0; i < cnt; i++ {
		list.Add(i, i+1)
	}
	for i := 0; i < cnt; i++ {
		list.Set(i, i+2)
	}
	for i := 0; i < cnt; i++ {
		assert.Equal(t, i+2, list.Get(i))
	}
}

func TestList_Add(t *testing.T) {
	list := NewList()
	cnt := 500000
	for i := 0; i < cnt; i++ {
		list.Add(i, i+1)
	}
	assert.Equal(t, cnt, list.Len())
	r := list.sentinel.height
	for r >= 0 {
		cur := list.sentinel
		j := -1
		for cur.next[r] != nil {
			j += cur.length[r]
			cur = cur.next[r]
			assert.Equal(t, j+1, cur.Val)
		}
		r--
	}
}

func TestList_Remove(t *testing.T) {
	list := NewList()
	cnt := 900000
	for i := 0; i < cnt; i++ {
		list.Add(i, i+1)
	}
	assert.Equal(t, cnt, list.Len())
	assert.Equal(t, 1, list.Remove(0))
	assert.Equal(t, cnt-1, list.Len())
	assert.Equal(t, cnt-1, list.cnt)
	assert.Equal(t, 2, list.Get(0))
	list.Add(0, 1)
	assert.Equal(t, cnt, list.Len())
	assert.Equal(t, cnt, list.Remove(cnt-1))
	assert.Equal(t, cnt-1, list.Len())
	assert.Equal(t, cnt-1, list.Get(cnt-2))
	list.Add(cnt-1, cnt)
	assert.Equal(t, cnt, list.Len())
	assert.Equal(t, cnt/2+1, list.Remove(cnt/2))
	assert.Equal(t, cnt-1, list.Len())
	assert.Equal(t, cnt/2+2, list.Get(cnt/2))
	assert.Equal(t, 1, list.Get(0))
	assert.Equal(t, cnt, list.Get(cnt-2))
}
