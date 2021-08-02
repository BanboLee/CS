package binary_heap

// get left child`s index.
func left(i int) int {
	return 2*i + 1
}

// get right child`s index.
func right(i int) int {
	return 2*i + 2
}

// get parent`s index
func parent(i int) int {
	return (i - 1) / 2
}

type BinaryMinHeap struct {
	data []int
	n    int
}

func (bh *BinaryMinHeap) Len() int {
	return bh.n
}

func (bh *BinaryMinHeap) swap(i, j int) {
	bh.data[i], bh.data[j] = bh.data[j], bh.data[i]
}

func (bh *BinaryMinHeap) shiftUp(i int) {
	p := parent(i)
	for i > 0 && bh.data[i] < bh.data[p] {
		bh.swap(p, i)
		i = p
		p = parent(i)
	}
}

func (bh *BinaryMinHeap) Add(v int) {
	bh.data = append(bh.data, v)
	bh.n++
	bh.shiftUp(bh.n - 1)
}

func (bh *BinaryMinHeap) compare(x, y int) bool {
	if x >= bh.n {
		return false
	} else if y >= bh.n {
		return true
	}
	return bh.data[x] < bh.data[y]
}

func (bh *BinaryMinHeap) shiftDown(i int) {
	l, r := left(i), right(i)
	for l < bh.n || r < bh.n {
		if bh.compare(l, i) && bh.compare(l, r) {
			bh.swap(l, i)
			i = l
			r = right(i)
			l = left(i)
		} else if bh.compare(r, i) && bh.compare(r, l) {
			bh.swap(r, i)
			i = r
			l = left(i)
			r = right(i)
		} else {
			return
		}
	}
}

// Pop the element at top.
func (bh *BinaryMinHeap) Remove() int {
	if bh.n == 0 {
		return -1
	}
	x := bh.data[0]
	bh.swap(0, bh.n-1)
	bh.data = bh.data[:bh.n-1]
	bh.n--
	bh.shiftDown(0)
	return x
}
