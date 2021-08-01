package skiplist_list

import (
	"math/rand"
	"time"
)

type Node struct {
	Val    int
	height int
	length []int
	next   []*Node
}

type List struct {
	sentinel *Node
	cnt      int
}

// For debug
func (list *List) print() {
	println("-------------------------------------")
	r := list.sentinel.height
	for r >= 0 {
		cur := list.sentinel
		for cur != nil {
			print(cur.Val, "|", cur.length[r])
			j := cur.length[r]
			for j >= 1 {
				print("\t")
				j--
			}
			cur = cur.next[r]
		}
		println()
		r--
	}
	println("-------------------------------------")
}

func (list *List) Len() int {
	return list.cnt
}

func NewList() *List {
	return &List{
		sentinel: &Node{
			Val:    0,
			height: 0,
			length: []int{0},
			next:   []*Node{nil},
		},
		cnt: 0,
	}
}

func (list *List) findPre(i int) *Node {
	cur := list.sentinel
	r := cur.height
	j := -1
	for r >= 0 {
		for cur.next[r] != nil && j+cur.length[r] < i {
			j += cur.length[r]
			cur = cur.next[r]
		}
		r--
	}
	return cur
}

func (list *List) Get(i int) int {
	if i >= list.cnt {
		panic("index out of range")
	}

	pre := list.findPre(i)
	return pre.next[0].Val
}

func (list *List) Set(i, v int) int {
	if i >= list.cnt {
		panic("index out of range")
	}

	pre := list.findPre(i)
	x := pre.next[0].Val
	pre.next[0].Val = v
	return x
}

func (list *List) pickHeight() int {
	rand.Seed(time.Now().UnixNano() - int64(list.cnt))
	k := rand.Int()
	h := 0
	m := 1
	for k&m != 0 {
		h++
		m <<= 1
	}
	// println("zpickHeight", k, h)
	return h
}

func (list *List) Add(i, v int) {
	if i > list.cnt {
		panic("index out of range")
	}
	height := list.pickHeight()
	node := &Node{
		Val:    v,
		height: height,
		length: make([]int, height+1),
		next:   make([]*Node, height+1),
	}
	list.add(i, node)
}

func (list *List) add(i int, node *Node) {
	cur := list.sentinel
	j := -1
	r := cur.height
	for r >= 0 {
		// 寻找前一个节点
		for cur.next[r] != nil && j+cur.length[r] < i {
			j += cur.length[r]
			cur = cur.next[r]
		}

		cur.length[r]++ // 要往后插入一个节点，length必然+1
		if r <= node.height {
			node.next[r] = cur.next[r]
			cur.next[r] = node
			node.length[r] = cur.length[r] - (i - j) // 最后一个值得length无所谓
			cur.length[r] = i - j                    // 前一个值现在获得了正确的length
		}
		r--
	}

	// 如果高于跳表的节点，则把跳表增高即可
	for list.sentinel.height < node.height {
		list.sentinel.height++
		list.sentinel.length = append(list.sentinel.length, i+1)
		list.sentinel.next = append(list.sentinel.next, node)
	}
	list.cnt++
}

func (list *List) Remove(i int) int {
	cur := list.sentinel
	x := -1
	r := cur.height
	j := -1
	for r >= 0 {
		for cur.next[r] != nil && j+cur.length[r] < i {
			j += cur.length[r]
			cur = cur.next[r]
		}
		cur.length[r]--
		if j+cur.length[r]+1 == i && cur.next[r] != nil {
			toDel := cur.next[r]
			x = toDel.Val
			cur.length[r] += toDel.length[r]
			cur.next[r] = toDel.next[r]
			if cur == list.sentinel && cur.next[r] == nil {
				list.sentinel.height--
				list.sentinel.length = list.sentinel.length[:len(list.sentinel.length)-1]
				list.sentinel.next = list.sentinel.next[:len(list.sentinel.next)-1]
			}
		}
		r--
	}
	list.cnt--
	return x
}
